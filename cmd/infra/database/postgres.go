package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var pgxPool *pgxpool.Pool

func InitPostgresConn() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
	ctx := context.Background()
	var (
		host     = os.Getenv("POSTGRES_HOST")
		port     = os.Getenv("POSTGRES_PORT")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
	)

	fmt.Println("host", host)
	urlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable", host, port, user, password)
	p, err := pgxpool.New(ctx, urlConn)
	if err != nil {
		log.Fatal(err)
	}

	if err = p.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	pgxPool = p
}

func GetPostgresPool() *pgxpool.Pool {
	return pgxPool
}
