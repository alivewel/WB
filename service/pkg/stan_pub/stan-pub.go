package stan_pub

import (
	"log"
	"sync"
	"time"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

var (
	Async     = false
	ClusterID = "test-cluster"
	ClientID  = "stan-pub"
	Subj      = "foo"
	UserCreds string
)

func PublishMessage(msg []byte, subj string, async bool, userCreds string, clusterID string, clientID string) {
	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Streaming Example Publisher")}
	// Use UserCredentials
	if userCreds != "" {
		opts = append(opts, nats.UserCredentials(userCreds))
	}
	// Connect to NATS
	nc, err := nats.Connect(stan.DefaultNatsURL, opts...)
	if err != nil {

		log.Fatal(err)
	}
	defer nc.Close()
	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, stan.DefaultNatsURL)
	}
	defer sc.Close()

	ch := make(chan bool)
	var glock sync.Mutex
	var guid string
	acb := func(lguid string, err error) {
		glock.Lock()
		log.Printf("Received ACK for guid %s\n", lguid)
		defer glock.Unlock()
		if err != nil {
			log.Fatalf("Error in server ack for guid %s: %v\n", lguid, err)
		}
		if lguid != guid {
			log.Fatalf("Expected a matching guid in ack callback, got %s vs %s\n", lguid, guid)
		}
		ch <- true
	}

	if !async {
		err = sc.Publish(subj, msg)
		if err != nil {
			log.Fatalf("Error during publish: %v\n", err)
		}
		// log.Printf("Published [%s] : '%s'\n", subj, msg)
	} else {
		glock.Lock()
		guid, err = sc.PublishAsync(subj, msg, acb)
		if err != nil {
			log.Fatalf("Error during async publish: %v\n", err)
		}
		glock.Unlock()
		if guid == "" {
			log.Fatal("Expected non-empty guid to be returned.")
		}
		log.Printf("Published [%s] : '%s' [guid: %s]\n", subj, msg, guid)

		select {
		case <-ch:
			break
		case <-time.After(5 * time.Second):
			log.Fatal("timeout")
		}
	}
}
