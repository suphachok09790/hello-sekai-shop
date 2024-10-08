package database

import (
	"context"
	"log"
	"time"

	"github.com/suphachok09790/hello-sekai-shop/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DbConn(pctx context.Context, cfg *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.Url))
	if err != nil {
		log.Fatalf("Error: Contect to database error:  %s", err.Error())
	}

	//Ping database เพื่อเช็คว่าเชื่อมได้ไหม
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Errpr: Pinging to database error:  %s", err.Error())
	}
	return client
}
