package repo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PublicRepoModel struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type CreateRepoModel struct {
	OwnerID     string `json:"ownerId"`                 // Owner's ID
	Name        string `json:"name" binding:"required"` // Repo name
	Description string `json:"description"`             // Repo description
}
