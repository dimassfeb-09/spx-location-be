package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	config := GetConfig()

	var (
		host     = config.PostgresConfig.Host     // Perbarui host ke host Supabase
		port     = config.PostgresConfig.Port     // Port biasanya sama
		user     = config.PostgresConfig.User     // Ganti dengan user Supabase Anda
		password = config.PostgresConfig.Password // Ganti dengan password Supabase Anda
		dbname   = config.PostgresConfig.DBName   // Ganti dengan nama database Supabase Anda
	)

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?binary_parameters=yes", user, password, host, port, dbname)

	// Create a new connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database")
	return db
}
