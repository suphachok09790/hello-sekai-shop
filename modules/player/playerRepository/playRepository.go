package playerRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	PlayerRepositoryService interface{}

	playRepository struct {
		db *mongo.Client
	}
)

func NewPlayerRepository(db *mongo.Client) PlayerRepositoryService {
	return &playRepository{db: db}
}

func (r *playRepository) playerDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("player_db")
}
