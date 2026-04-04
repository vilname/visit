// Package main сервис для обработки визитов в медицинских организациях
package main

import (
	"fmt"
	"os"
	"visit/config"
	"visit/config/storage"

	"github.com/joho/godotenv"
)

// Package main
func main() {
	fmt.Println("Start App v0.01")

	err := godotenv.Load()
	if err != nil {
		return
	}

	dbUrl := os.Getenv("DATABASE_URL")

	err = storage.InitDB(dbUrl)

	if err != nil {
		fmt.Println("main: " + err.Error())
		return
	}

	router := config.InitRoute()

	defer func() {
		db := storage.GetDB()
		db.Close()
	}()

	defer func() {
		db := storage.GetDB()
		db.Close()
	}()

	fmt.Println("init")

	err = router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return
	}
}
