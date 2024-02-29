package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"urbskali/ssserver/models"
	"urbskali/ssserver/state"

	"github.com/jackc/pgx/v5"
)

func StartUp() {
	fmt.Println("Starting up...")
	// check if the config file exists
	if _, err := os.Stat("./config.json"); os.IsNotExist(err) {
		log.Fatal("Config file does not exist, try running 'ssserver setup'")
	}
	// load the config
	config, err := models.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// print the config
	fmt.Println(config.String())
	// set the config
	state.Config = config

	conn, err := pgx.Connect(context.Background(), state.Config.DB_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	state.DB = conn
}
