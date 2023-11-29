package model

import (
	"time"

	repo "github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/public"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PrivateRepoModel struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	OwnerID     string             `json:"ownerId" bson:"owner_id"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

// To PublicRepoModel
func (privateRepoModel *PrivateRepoModel) ToPublicRepoModel() *repo.PublicRepoModel {
	return &repo.PublicRepoModel{
		ID:        privateRepoModel.ID,
		Name:      privateRepoModel.Name,
		CreatedAt: privateRepoModel.CreatedAt,
	}
}
