package main

import (
	"context"
	"fmt"
	"os"
	"visit/config"
	"visit/config/storage"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Start App v0.01")

	err := godotenv.Load()

	storage.InitDB()
	router := config.InitRoute()

	defer func() {
		db := storage.GetDB()
		db.Close()
	}()

	defer func(ctx context.Context) {
		db := storage.GetDB()
		db.Close()
	}(context.Background())

	fmt.Println("init")

	err = router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return
	}
}
