package repository_test

import (
	"context"
	"errors"
	"reflect"
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

func (m *MongoAdapterMock) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
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

type MongoCursorMock struct {
	mock.Mock
	current int
	data    []interface{} // Holds the slice of data to be returned
}

func (m *MongoCursorMock) Next(ctx context.Context) bool {
	return m.current < len(m.data)
}

func (m *MongoCursorMock) Decode(val interface{}) error {
	args := m.Called(val)
	if m.current < len(m.data) {
		// Copy the data to the passed pointer
		reflect.ValueOf(val).Elem().Set(reflect.ValueOf(m.data[m.current]).Elem())
		m.current++
		return args.Error(0)
	}
	return errors.New("no more items")
}

func (m *MongoCursorMock) Close(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Helper function to setup the mock data
func (m *MongoCursorMock) SetData(data []interface{}) {
	m.data = data
}

func TestFindAll_Success(t *testing.T) {
	ctx := context.TODO()
	adapterMock := new(MongoAdapterMock)
	repository := repository.NewRepoRepository(adapterMock)

	// Create a slice of expected results
	expectedRepos := []*model.PrivateRepoModel{
		{
			ID:          primitive.NewObjectID(),
			Name:        "test1",
			OwnerID:     primitive.NewObjectID().Hex(),
			Description: "test description 1",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "test2",
			OwnerID:     primitive.NewObjectID().Hex(),
			Description: "test description 2",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	// Mock the Find call to return a cursor with expected results
	cursor := &MongoCursorMock{} // You would need to implement MongoCursorMock
	cursor.On("Next", mock.Anything).Return(true).Times(len(expectedRepos))
	cursor.On("Decode", mock.Anything).Return(func(val interface{}) error {
		*val.(*model.PrivateRepoModel) = *expectedRepos[0]
		expectedRepos = expectedRepos[1:]
		return nil
	}).Times(len(expectedRepos))
	cursor.On("Close", mock.Anything).Return(nil)

	adapterMock.On("Find", ctx, mock.Anything, mock.Anything).Return(cursor, nil)

	repos, err := repository.FindAllByOwnerID(ctx, primitive.NewObjectID().Hex(), 1, 10)

	assert.Nil(t, err)
	assert.Equal(t, len(expectedRepos), len(repos))
	for i, repo := range repos {
		assert.Equal(t, expectedRepos[i].ID, repo.ID)
		assert.Equal(t, expectedRepos[i].Name, repo.Name)
		// ... more assertions as needed
	}

	adapterMock.AssertExpectations(t)
	cursor.AssertExpectations(t)
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
