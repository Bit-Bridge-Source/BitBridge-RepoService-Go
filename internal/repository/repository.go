package repository

import (
	"context"

	"github.com/Bit-Bridge-Source/BitBridge-CommonService-Go/public/adapter"
	"github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RepoRepository interface {
	FindById(ctx context.Context, id string) (*model.PrivateRepoModel, error)
	FindByName(ctx context.Context, name string) (*model.PrivateRepoModel, error)
	Create(ctx context.Context, repo *model.PrivateRepoModel) (*model.PrivateRepoModel, error)
	UpdateOne(ctx context.Context, repo *model.PrivateRepoModel) (*model.PrivateRepoModel, error)
	DeleteOne(ctx context.Context, repo *model.PrivateRepoModel) error
}

type MongoRepoRepository struct {
	Collection adapter.MongoAdapter
}

func NewRepoRepository(collection adapter.MongoAdapter) *MongoRepoRepository {
	return &MongoRepoRepository{
		Collection: collection,
	}
}

func (m *MongoRepoRepository) FindById(ctx context.Context, id string) (*model.PrivateRepoModel, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	repo := &model.PrivateRepoModel{}
	err = m.Collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(repo)

	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (m *MongoRepoRepository) FindByName(ctx context.Context, name string) (*model.PrivateRepoModel, error) {
	repo := &model.PrivateRepoModel{}
	err := m.Collection.FindOne(ctx, bson.M{"name": name}).Decode(repo)

	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (m *MongoRepoRepository) Create(ctx context.Context, repo *model.PrivateRepoModel) (*model.PrivateRepoModel, error) {
	_, err := m.Collection.InsertOne(ctx, repo)

	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (m *MongoRepoRepository) UpdateOne(ctx context.Context, repo *model.PrivateRepoModel) (*model.PrivateRepoModel, error) {
	_, err := m.Collection.UpdateOne(ctx, bson.M{"_id": repo.ID}, bson.M{"$set": repo})

	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (m *MongoRepoRepository) DeleteOne(ctx context.Context, repo *model.PrivateRepoModel) error {
	_, err := m.Collection.DeleteOne(ctx, bson.M{"_id": repo.ID})

	if err != nil {
		return err
	}

	return nil
}
