package utils

import (
	"fmt"
	"log"
	"os"

	"urbskali/sserver/state"
)

func StartUp() {
	fmt.Println("Starting up...")
	// check if the config file exists
	if _, err := os.Stat("./config.json"); os.IsNotExist(err) {
		log.Fatal("Config file does not exist, try running 'ssserver setup'")
	}
	// load the config
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// print the config
	fmt.Println(config.String())
	// set the config
	state.Config = config
}
