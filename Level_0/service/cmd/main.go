package main

import (
	"log"

	"service/pkg/database"
	"service/pkg/memorycache"
	"service/pkg/stan_sub"
	server "service/web"
	"time"
)

func main() {
	// Connecting to NATS
	nc, err := stan_sub.ConnectNATS(stan_sub.URL, stan_sub.UserCreds)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Connecting to NATS Streaming
	sc, err := stan_sub.ConnectNATSStreaming(nc, stan_sub.ClusterID, stan_sub.ClientID)
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, stan_sub.URL)
	}

	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	err = database.CreateTable(db)
	if err != nil {
		log.Fatalf("Error creating table in database: %v", err)
	}

	// Create a container for the cache
	cache := memorycache.New(5*time.Minute, 10*time.Minute)

	// Retrieving data from the database
	err = database.RetrieveData(db, cache)
	if err != nil {
		log.Fatalf("Error while retrieving data from database: %v", err)
	}

	go server.RunServer(cache)

	// Subscription Processing
	stan_sub.SubscribeAndListen(sc, "foo", db, cache)
}
