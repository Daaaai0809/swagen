package main

import (
	"fmt"
	"os"

	"github.com/Daaaai0809/swagen"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	c, err := swagen.NewConfig()
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err.Error())
		os.Exit(1)
	}

	swagen.SetConfig(c)

	Execute()
}
