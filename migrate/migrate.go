package main

import (
	"context"
	"fmt"
	"go-rest-api/db"
	"log"
)

func main() {
	client := db.NewClient()
	defer fmt.Println("Successfully migrated")
	defer client.Close()

	// 自動マイグレーションを実行する。
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
