package item

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Item struct {
		Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Title       string             `json:"title" validate:"required,max=64"`
		Price       float64            `json:"price" validate:"required"`
		Damage      int                `json:"damage" validate:"required"`
		ImageUrl    string             `json:"image_url" validate:"required,max=255"`
		UsageStatus bool               `json:"usage_status"`
		CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time          `json:"update_at" bson:"update_at"`
	}
)
