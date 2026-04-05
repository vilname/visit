// Package model для запуска миграций
package main

import (
	"fmt"
	"log"
	"visit/config"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("env: ", err.Error())
	}

	if err := config.InitMigrationDB(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("миграции отработали")
}
