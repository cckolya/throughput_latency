package main

import (
	"devhands/config"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("config:", err)
	}

	db, err := NewPostgres(config.Postgres)
	if err != nil {
		fmt.Println("postgres:", err)
	}

	cache, err := NewRedis(config.Redis)
	if err != nil {
		fmt.Println("redis:", err)
	}

	handler := NewHandler(db, cache)

	err = handler.Handle()
	if err != nil {
		fmt.Println("handler:", err)
	}
}
