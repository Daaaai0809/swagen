package main

import (
	"fmt"
	"os"

	"github.com/Daaaai0809/swagen/cmd"
	"github.com/Daaaai0809/swagen/config"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	c, err := config.NewConfig()
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err.Error())
		os.Exit(1)
	}

	config.SetConfig(c)

	cmd.Execute()
}
