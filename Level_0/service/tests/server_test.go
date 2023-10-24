package server

import (
	"database/sql"
	"log"
	"os"
	"service/pkg/database"
	"service/pkg/memorycache"
	"service/pkg/stan_pub"
	"time"

	nats "github.com/nats-io/nats.go"

	"testing"
)

var cache *memorycache.Cache
var db *sql.DB
var nc *nats.Conn

func TestMain(m *testing.M) {
	// Starting a server in a goroutine
	var serverDone chan struct{}
	serverDone = make(chan struct{})
	go func() {
		db, cache, nc = setupAndRunServer()
		close(serverDone)
	}()

	// Running tests
	time.Sleep(1 * time.Second)
	code := m.Run()

	// Waiting for the server to complete
	<-serverDone

	os.Exit(code)
}

func TestMessageString(t *testing.T) {
	var msg = []byte("Hello, Wildberries!")

	numbRecordBeforePubCache := cache.Count()
	numbRecordBeforePubDB, err := database.GetRecordCount(db)
	if err != nil {
		log.Fatalf("Error sending database query: ! %v", err)
	}
	log.Printf("numbRecordBeforePubCache: %v", numbRecordBeforePubCache)
	log.Printf("numbRecordBeforePubDB: %v", numbRecordBeforePubDB)
	stan_pub.PublishMessage(msg, stan_pub.Subj, stan_pub.Async, stan_pub.UserCreds, stan_pub.ClusterID, stan_pub.ClientID)

	time.Sleep(1 * time.Second)
	numbRecordAfterPubCache := cache.Count()
	numbRecordAfterPubDB, err := database.GetRecordCount(db)
	if err != nil {
		log.Fatalf("Error sending database query: %v", err)
	}

	log.Printf("numbRecordAfterPubCache: %v", numbRecordAfterPubCache)
	log.Printf("numbRecordAfterPubDB: %v", numbRecordAfterPubDB)

	if numbRecordBeforePubDB != numbRecordAfterPubDB || numbRecordBeforePubCache != numbRecordAfterPubCache {
		t.Errorf("An incorrect record was added to the database or cache")
	}

}

func TestMessageInt(t *testing.T) {
	var msg = []byte("12345")

	numbRecordBeforePubCache := cache.Count()
	numbRecordBeforePubDB, err := database.GetRecordCount(db)
	if err != nil {
		log.Fatalf("Error sending database query: ! %v", err)
	}
	log.Printf("numbRecordBeforePubCache: %v", numbRecordBeforePubCache)
	log.Printf("numbRecordBeforePubDB: %v", numbRecordBeforePubDB)
	stan_pub.PublishMessage(msg, stan_pub.Subj, stan_pub.Async, stan_pub.UserCreds, stan_pub.ClusterID, stan_pub.ClientID)

	time.Sleep(1 * time.Second)
	numbRecordAfterPubCache := cache.Count()
	numbRecordAfterPubDB, err := database.GetRecordCount(db)
	if err != nil {
		log.Fatalf("Error sending database query: %v", err)
	}

	log.Printf("numbRecordAfterPubCache: %v", numbRecordAfterPubCache)
	log.Printf("numbRecordAfterPubDB: %v", numbRecordAfterPubDB)

	if numbRecordBeforePubDB != numbRecordAfterPubDB || numbRecordBeforePubCache != numbRecordAfterPubCache {
		t.Errorf("An incorrect record was added to the database or cache")
	}
}

func TestMessageBool(t *testing.T) {
	var msg = []byte("true")

	numbRecordBeforePubCache := cache.Count()
	numbRecordBeforePubDB, err := database.GetRecordCount(db)
	if err != nil {
		log.Fatalf("Error sending database query: ! %v", err)
	}
	log.Printf("numbRecordBeforePubCache: %v", numbRecordBeforePubCache)
	log.Printf("numbRecordBeforePubDB: %v", numbRecordBeforePubDB)
	stan_pub.PublishMessage(msg, stan_pub.Subj, stan_pub.Async, stan_pub.UserCreds, stan_pub.ClusterID, stan_pub.ClientID)

	time.Sleep(1 * time.Second)
	numbRecordAfterPubCache := cache.Count()
	numbRecordAfterPubDB, err := database.GetRecordCount(db)
	if err != nil {
		log.Fatalf("Error sending database query: %v", err)
	}

	log.Printf("numbRecordAfterPubCache: %v", numbRecordAfterPubCache)
	log.Printf("numbRecordAfterPubDB: %v", numbRecordAfterPubDB)

	if numbRecordBeforePubDB != numbRecordAfterPubDB || numbRecordBeforePubCache != numbRecordAfterPubCache {
		t.Errorf("An incorrect record was added to the database or cache")
	}
}

func TestMessageArray(t *testing.T) {
	var msg = []byte("[1, 2, 3]")

	numbRecordBeforePubCache := cache.Count()
	numbRecordBeforePubDB, err := database.GetRecordCount(db)
	if err != nil {
		log.Fatalf("Error sending database query: ! %v", err)
	}
	log.Printf("numbRecordBeforePubCache: %v", numbRecordBeforePubCache)
	log.Printf("numbRecordBeforePubDB: %v", numbRecordBeforePubDB)
	stan_pub.PublishMessage(msg, stan_pub.Subj, stan_pub.Async, stan_pub.UserCreds, stan_pub.ClusterID, stan_pub.ClientID)

	time.Sleep(1 * time.Second)
	numbRecordAfterPubCache := cache.Count()
	numbRecordAfterPubDB, err := database.GetRecordCount(db)
	if err != nil {
		log.Fatalf("Error sending database query: %v", err)
	}

	log.Printf("numbRecordAfterPubCache: %v", numbRecordAfterPubCache)
	log.Printf("numbRecordAfterPubDB: %v", numbRecordAfterPubDB)

	if numbRecordBeforePubDB != numbRecordAfterPubDB || numbRecordBeforePubCache != numbRecordAfterPubCache {
		t.Errorf("An incorrect record was added to the database or cache")
	}
}

func TestMessageBrackets(t *testing.T) {
	var msg = []byte("{}")

	numbRecordBeforePubCache := cache.Count()
	numbRecordBeforePubDB, err := database.GetRecordCount(db)
	if err != nil {
		log.Fatalf("Error sending database query: ! %v", err)
	}
	log.Printf("numbRecordBeforePubCache: %v", numbRecordBeforePubCache)
	log.Printf("numbRecordBeforePubDB: %v", numbRecordBeforePubDB)
	stan_pub.PublishMessage(msg, stan_pub.Subj, stan_pub.Async, stan_pub.UserCreds, stan_pub.ClusterID, stan_pub.ClientID)

	time.Sleep(1 * time.Second)
	numbRecordAfterPubCache := cache.Count()
	numbRecordAfterPubDB, err := database.GetRecordCount(db)
	if err != nil {
		log.Fatalf("Error sending database query: %v", err)
	}

	log.Printf("numbRecordAfterPubCache: %v", numbRecordAfterPubCache)
	log.Printf("numbRecordAfterPubDB: %v", numbRecordAfterPubDB)

	if numbRecordBeforePubDB != numbRecordAfterPubDB || numbRecordBeforePubCache != numbRecordAfterPubCache {
		t.Errorf("An incorrect record was added to the database or cache")
	}
}

func TestCorrectMessage(t *testing.T) {
	jsonFile := "../data/data.json"
	msg, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
		os.Exit(1)
	}

	numbRecordBeforePubCache := cache.Count()
	numbRecordBeforePubDB, err := database.GetRecordCount(db)
	if err != nil {
		log.Fatalf("Error sending database query: ! %v", err)
	}
	log.Printf("numbRecordBeforePubCache: %v", numbRecordBeforePubCache)
	log.Printf("numbRecordBeforePubDB: %v", numbRecordBeforePubDB)
	stan_pub.PublishMessage(msg, stan_pub.Subj, stan_pub.Async, stan_pub.UserCreds, stan_pub.ClusterID, stan_pub.ClientID)

	time.Sleep(1 * time.Second)
	numbRecordAfterPubCache := cache.Count()
	numbRecordAfterPubDB, err := database.GetRecordCount(db)
	if err != nil {
		log.Fatalf("Error sending database query: %v", err)
	}

	log.Printf("numbRecordAfterPubCache: %v", numbRecordAfterPubCache)
	log.Printf("numbRecordAfterPubDB: %v", numbRecordAfterPubDB)

	if numbRecordBeforePubDB+1 != numbRecordAfterPubDB || numbRecordBeforePubCache+1 != numbRecordAfterPubCache {
		t.Errorf("An incorrect record was added to the database or cache")
	}
}
