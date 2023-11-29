package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/internal/model"
	"github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoAdapterMock struct {
	mock.Mock
}

func (m *MongoAdapterMock) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document, opts)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MongoAdapterMock) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update, opts)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MongoAdapterMock) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

func (m *MongoAdapterMock) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.SingleResult)
}

type SingleResultWrapper struct {
	decoder SingleResultDecoder
}

func (s *SingleResultWrapper) Decode(v interface{}) error {
	return s.decoder.Decode(v)
}

type SingleResultDecoder interface {
	Decode(v interface{}) error
}

type SingleResultMock struct {
	mock.Mock
}

func (s *SingleResultMock) Decode(val interface{}) error {
	args := s.Called(val)
	return args.Error(0)
}

func TestCreateRepo_Success(t *testing.T) {
	ctx := context.TODO()
	adapterMock := new(MongoAdapterMock)

	repository := repository.NewRepoRepository(adapterMock)
	repoToBeCreated := &model.PrivateRepoModel{
		ID:          primitive.NewObjectID(),
		Name:        "test",
		OwnerID:     primitive.NewObjectID().Hex(),
		Description: "test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	adapterMock.On("InsertOne", ctx, repoToBeCreated, mock.Anything).Return(&mongo.InsertOneResult{}, nil)

	_, err := repository.Create(ctx, repoToBeCreated)

	assert.Nil(t, err)

	adapterMock.AssertExpectations(t)
}

func TestCreateRepo_Error_Duplicate(t *testing.T) {
	ctx := context.TODO()
	adapterMock := new(MongoAdapterMock)

	repository := repository.NewRepoRepository(adapterMock)
	repoToBeCreated := &model.PrivateRepoModel{
		ID:          primitive.NewObjectID(),
		Name:        "test",
		OwnerID:     primitive.NewObjectID().Hex(),
		Description: "test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	adapterMock.On("InsertOne", ctx, repoToBeCreated, mock.Anything).Return(&mongo.InsertOneResult{}, mongo.WriteException{})

	_, err := repository.Create(ctx, repoToBeCreated)

	assert.NotNil(t, err)

	adapterMock.AssertExpectations(t)
}

func TestFindById_Sucess(t *testing.T) {
	ctx := context.TODO()
	adapterMock := new(MongoAdapterMock)

	repository := repository.NewRepoRepository(adapterMock)
	repoExpected := &model.PrivateRepoModel{
		ID:          primitive.NewObjectID(),
		Name:        "test",
		OwnerID:     primitive.NewObjectID().Hex(),
		Description: "test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	sr := mongo.NewSingleResultFromDocument(repoExpected, nil, bson.DefaultRegistry)

	adapterMock.On("FindOne", ctx, bson.M{"_id": repoExpected.ID}, mock.Anything).Return(sr)

	repo, err := repository.FindById(ctx, repoExpected.ID.Hex())

	assert.Nil(t, err)
	assert.Equal(t, repoExpected.ID, repo.ID)

	adapterMock.AssertExpectations(t)
}

func TestFindById_Error_InvalidID(t *testing.T) {
	ctx := context.TODO()
	adapterMock := new(MongoAdapterMock)

	repository := repository.NewRepoRepository(adapterMock)

	_, err := repository.FindById(ctx, "invalid")

	assert.NotNil(t, err)

	adapterMock.AssertExpectations(t)
}

func TestFindById_Error_NotFound(t *testing.T) {
	ctx := context.TODO()

	id := primitive.NewObjectID()
	adapterMock := new(MongoAdapterMock)

	sr := mongo.NewSingleResultFromDocument(&model.PrivateRepoModel{}, errors.New("not found"), bson.DefaultRegistry)

	adapterMock.On("FindOne", ctx, bson.M{"_id": id}, mock.Anything).Return(sr)

	repository := repository.NewRepoRepository(adapterMock)

	_, err := repository.FindById(ctx, id.Hex())

	assert.NotNil(t, err)

	adapterMock.AssertExpectations(t)
}

func TestFindByName_Sucess(t *testing.T) {
	ctx := context.TODO()
	adapterMock := new(MongoAdapterMock)

	repository := repository.NewRepoRepository(adapterMock)
	repoExpected := &model.PrivateRepoModel{
		ID:          primitive.NewObjectID(),
		Name:        "test",
		OwnerID:     primitive.NewObjectID().Hex(),
		Description: "test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	sr := mongo.NewSingleResultFromDocument(repoExpected, nil, bson.DefaultRegistry)

	adapterMock.On("FindOne", ctx, bson.M{"name": repoExpected.Name}, mock.Anything).Return(sr)

	repo, err := repository.FindByName(ctx, repoExpected.Name)

	assert.Nil(t, err)
	assert.Equal(t, repoExpected.Name, repo.Name)

	adapterMock.AssertExpectations(t)
}

func TestFindByName_Error_NotFound(t *testing.T) {
	ctx := context.TODO()

	name := "test"
	adapterMock := new(MongoAdapterMock)

	sr := mongo.NewSingleResultFromDocument(&model.PrivateRepoModel{}, errors.New("not found"), bson.DefaultRegistry)

	adapterMock.On("FindOne", ctx, bson.M{"name": name}, mock.Anything).Return(sr)

	repository := repository.NewRepoRepository(adapterMock)

	_, err := repository.FindByName(ctx, name)

	assert.NotNil(t, err)

	adapterMock.AssertExpectations(t)
}

func TestUpdateOne_Sucess(t *testing.T) {
	ctx := context.TODO()
	adapterMock := new(MongoAdapterMock)

	repository := repository.NewRepoRepository(adapterMock)
	repoExpected := &model.PrivateRepoModel{
		ID:          primitive.NewObjectID(),
		Name:        "test",
		OwnerID:     primitive.NewObjectID().Hex(),
		Description: "test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	adapterMock.On("UpdateOne", ctx, bson.M{"_id": repoExpected.ID}, bson.M{"$set": repoExpected}, mock.Anything).Return(&mongo.UpdateResult{}, nil)

	repo, err := repository.UpdateOne(ctx, repoExpected)

	assert.Nil(t, err)
	assert.Equal(t, repoExpected.ID, repo.ID)

	adapterMock.AssertExpectations(t)
}

func TestUpdateOne_Error_NotFound(t *testing.T) {
	ctx := context.TODO()
	adapterMock := new(MongoAdapterMock)

	repository := repository.NewRepoRepository(adapterMock)
	repoExpected := &model.PrivateRepoModel{
		ID:          primitive.NewObjectID(),
		Name:        "test",
		OwnerID:     primitive.NewObjectID().Hex(),
		Description: "test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	adapterMock.On("UpdateOne", ctx, bson.M{"_id": repoExpected.ID}, bson.M{"$set": repoExpected}, mock.Anything).Return(&mongo.UpdateResult{}, mongo.ErrNoDocuments)

	_, err := repository.UpdateOne(ctx, repoExpected)

	assert.NotNil(t, err)

	adapterMock.AssertExpectations(t)
}

func TestDeleteOne_Sucess(t *testing.T) {
	ctx := context.TODO()
	adapterMock := new(MongoAdapterMock)

	repository := repository.NewRepoRepository(adapterMock)
	repoExpected := &model.PrivateRepoModel{
		ID:          primitive.NewObjectID(),
		Name:        "test",
		OwnerID:     primitive.NewObjectID().Hex(),
		Description: "test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	adapterMock.On("DeleteOne", ctx, bson.M{"_id": repoExpected.ID}, mock.Anything).Return(&mongo.DeleteResult{}, nil)

	err := repository.DeleteOne(ctx, repoExpected)

	assert.Nil(t, err)

	adapterMock.AssertExpectations(t)
}

func TestDeleteOne_Error_NotFound(t *testing.T) {
	ctx := context.TODO()
	adapterMock := new(MongoAdapterMock)

	repository := repository.NewRepoRepository(adapterMock)
	repoExpected := &model.PrivateRepoModel{
		ID:          primitive.NewObjectID(),
		Name:        "test",
		OwnerID:     primitive.NewObjectID().Hex(),
		Description: "test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	adapterMock.On("DeleteOne", ctx, bson.M{"_id": repoExpected.ID}, mock.Anything).Return(&mongo.DeleteResult{}, mongo.ErrNoDocuments)

	err := repository.DeleteOne(ctx, repoExpected)

	assert.NotNil(t, err)

	adapterMock.AssertExpectations(t)
}
