package main

import (
	"context"
	"log"
	"os"

	"github.com/suphachok09790/hello-sekai-shop/config"
	"github.com/suphachok09790/hello-sekai-shop/pkg/database/migration"
)

func main() {
	ctx := context.Background()
	_ = ctx

	// Initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	switch cfg.App.Name {
	case "player":
	case "auth":
		migration.AuthMigrate(ctx, &cfg)
	case "item":
	case "inventory":
	case "payment":
	}
}
