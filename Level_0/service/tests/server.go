package server

import (
	"database/sql"
	"log"
	"service/pkg/database"
	"service/pkg/memorycache"
	"service/pkg/stan_sub"

	nats "github.com/nats-io/nats.go"

	server "service/web"
	"time"
)

func setupAndRunServer() (*sql.DB, *memorycache.Cache, *nats.Conn) {
	// Подключение к NATS
	nc, err := stan_sub.ConnectNATS(stan_sub.URL, stan_sub.UserCreds)
	if err != nil {
		log.Fatal(err)
	}
	// defer nc.Close()

	// Подключение к NATS Streaming
	sc, err := stan_sub.ConnectNATSStreaming(nc, stan_sub.ClusterID, stan_sub.ClientID)
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, stan_sub.URL)
	}

	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	// defer db.Close()

	err = database.CreateTable(db)
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы в базе данных: %v", err)
	}

	// Создаем контейнер для кэша
	cache := memorycache.New(5*time.Minute, 10*time.Minute)

	// Получение данных из БД
	err = database.RetrieveData(db, cache)
	if err != nil {
		log.Fatalf("Ошибка при извлечении данных из базы данных: %v", err)
	}
	
	go server.RunServer(cache)

	// Обработка подписки
	go stan_sub.SubscribeAndListen(sc, "foo", db, cache)
	return db, cache, nc
}
