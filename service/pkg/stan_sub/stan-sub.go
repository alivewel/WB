package stan_sub

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"database/sql"
	"service/pkg/database"
	"service/pkg/memorycache"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"

	_ "github.com/lib/pq"
)

func printMsg(m *stan.Msg, i int) {
	dataStr := string(m.Data)
	log.Printf("[#%d] Received: %s\n", i, dataStr)
}

var (
	URL         = stan.DefaultNatsURL
	ClusterID   = "test-cluster"
	ClientID    = "stan-sub"
	showTime    = false
	startSeq    uint64
	deliverAll  = true
	newOnly     = true
	deliverLast = false
	startDelta  string
	durable     string
	qgroup      string
	unsubscribe bool
	UserCreds   string

	numMsg = 0
)

func ConnectNATS(url, userCreds string) (*nats.Conn, error) {
	opts := []nats.Option{nats.Name("NATS Streaming Example Subscriber")}
	if userCreds != "" {
		opts = append(opts, nats.UserCredentials(userCreds))
	}

	nc, err := nats.Connect(url, opts...)
	return nc, err
}

func ConnectNATSStreaming(nc *nats.Conn, clusterID, clientID string) (stan.Conn, error) {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	return sc, err
}

func SubscribeAndListen(sc stan.Conn, subj string, db *sql.DB, cache *memorycache.Cache) {
	// Handling Subscription Options
	startOpt := stan.StartAt(pb.StartPosition_NewOnly)
	if startSeq != 0 {
		startOpt = stan.StartAtSequence(startSeq)
	} else if deliverLast {
		startOpt = stan.StartWithLastReceived()
	} else if deliverAll && !newOnly {
		startOpt = stan.DeliverAllAvailable()
	} else if startDelta != "" {
		ago, err := time.ParseDuration(startDelta)
		if err != nil {
			sc.Close()
			log.Fatal(err)
		}
		startOpt = stan.StartAtTimeDelta(ago)
	}

	// Defining a Message Processing Function
	mcb := func(msg *stan.Msg) {
		printMsg(msg, numMsg)
		numMsg++

		message := string(msg.Data)

		if !isValidJSONObject(message) {
			log.Printf("Message is not valid JSON.")
		} else {
			err := database.InsertJSON(db, message)
			if err != nil {
				log.Printf("Error when inserting data into database: %v", err)
			} else {
				id := cache.Count() + 1
				cache.Set(strconv.Itoa(id), message, 5*time.Minute)
			}
		}
	}

	// Create a subscription
	sub, err := sc.QueueSubscribe(subj, qgroup, mcb, startOpt, stan.DurableName(durable))
	if err != nil {
		sc.Close()
		log.Fatal(err)
	}

	log.Printf("Listening on [%s], clientID=[%s], qgroup=[%s] durable=[%s]\n", subj, ClientID, qgroup, durable)

	if showTime {
		log.SetFlags(log.LstdFlags)
	}

	// Waiting for SIGINT signal
	waitForInterrupt(sub, sc)
}

func waitForInterrupt(sub stan.Subscription, sc stan.Conn) {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			if durable == "" || unsubscribe {
				sub.Unsubscribe()
			}
			sc.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}

func isValidJSONObject(s string) bool {
	if s == "{}" {
		return false
	}

	var js interface{}
	err := json.Unmarshal([]byte(s), &js)
	return err == nil && isObject(js)
}

func isObject(js interface{}) bool {
	_, isObject := js.(map[string]interface{})
	return isObject
}
