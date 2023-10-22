package database

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"service/pkg/memorycache"
	"time"
)

func ConnectToDB() (*sql.DB, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("../configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading configuration: %v", err)
	}

	user := viper.GetString("user")
	dbname := viper.GetString("dbname")
	password := viper.GetString("password")
	sslmode := viper.GetString("sslmode")

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", user, dbname, password, sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertJSON(db *sql.DB, jsonString string) error {
	query := "INSERT INTO temp (msg) VALUES ($1)"
	_, err := db.Exec(query, jsonString)
	return err
}

func CreateTable(db *sql.DB) error {
	query := "CREATE TABLE IF NOT EXISTS temp (id serial primary key, msg varchar NOT NULL)"
	_, err := db.Exec(query)
	return err
}

func RetrieveData(db *sql.DB, c *memorycache.Cache) error {
	query := "SELECT id, msg FROM temp"
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var msg string
		if err := rows.Scan(&id, &msg); err != nil {
			return err
		}
		c.Set(id, msg, 5*time.Minute)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func GetRecordCount(db *sql.DB) (int, error) {
	query := "SELECT COUNT(*) FROM temp"
	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
