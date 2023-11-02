package authRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	AuthRepositoryService interface{}

	authRepository struct {
		db *mongo.Client
	}
)

// chain pattern
func NewAuthRepository(db *mongo.Client) AuthRepositoryService {
	return &authRepository{db}
}

// ฟังกืชันต่อdatabase
func (r *authRepository) authDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("auth_db")
}
