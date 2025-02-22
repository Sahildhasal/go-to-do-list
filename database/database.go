package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn

func InitDb() {
	err := godotenv.Load()

	if err != nil {
		log.Printf("Error while loading DotEnv %v", err)
	}

	connectionString := os.Getenv("DATABASE_URL")

	DB, err = pgx.Connect(context.Background(), connectionString)

	if err != nil {
		log.Printf("Error while Connecting to the Database %v", err)
	}

	if err := DB.Ping(context.Background()); err != nil {
		log.Printf("Error while Pinging %v", err)
	}

	query := `CREATE TABLE IF NOT EXISTS todo(id SERIAL PRIMARY KEY, name TEXT NOT NULL, description TEXT NOT NULL, date DATE NOT NULL);`

	result, err := DB.Exec(context.Background(), query)

	if err != nil {
		log.Printf("Error while Creating Table %v", err)
	}

	log.Printf("Table %v", result)
}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/jackc/pgx/v5"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	connStr := os.Getenv("DATABASE_URL")
// 	conn, err := pgx.Connect(context.Background(), connStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close(context.Background())
// 	_, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS playing_with_neon(id SERIAL PRIMARY KEY, name TEXT NOT NULL, value REAL);")
// 	if err != nil {
// 		panic(err)
// 	}
// 	_, err = conn.Exec(context.Background(), "INSERT INTO playing_with_neon(name, value) SELECT LEFT(md5(i::TEXT), 10), random() FROM generate_series(1, 10) s(i);")
// 	if err != nil {
// 		panic(err)
// 	}
// 	rows, err := conn.Query(context.Background(), "SELECT * FROM playing_with_neon")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var id int32
// 		var name string
// 		var value float32
// 		if err := rows.Scan(&id, &name, &value); err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("%d | %s | %f\n", id, name, value)
// 	}
// }
