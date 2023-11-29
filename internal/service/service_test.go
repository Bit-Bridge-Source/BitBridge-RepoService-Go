package service_test

import (
	"context"
	"testing"

	"github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/internal/model"
	"github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/internal/service"
	public_repo "github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/public"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Create(ctx context.Context, repo *model.PrivateRepoModel) (*model.PrivateRepoModel, error) {
	args := r.Called(ctx, repo)
	return args.Get(0).(*model.PrivateRepoModel), args.Error(1)
}

func (r *RepositoryMock) FindById(ctx context.Context, id string) (*model.PrivateRepoModel, error) {
	args := r.Called(ctx, id)
	return args.Get(0).(*model.PrivateRepoModel), args.Error(1)
}

func (r *RepositoryMock) FindByName(ctx context.Context, name string) (*model.PrivateRepoModel, error) {
	args := r.Called(ctx, name)
	return args.Get(0).(*model.PrivateRepoModel), args.Error(1)
}

func (r *RepositoryMock) UpdateOne(ctx context.Context, repo *model.PrivateRepoModel) (*model.PrivateRepoModel, error) {
	args := r.Called(ctx, repo)
	return args.Get(0).(*model.PrivateRepoModel), args.Error(1)
}

func (r *RepositoryMock) DeleteOne(ctx context.Context, repo *model.PrivateRepoModel) error {
	args := r.Called(ctx, repo)
	return args.Error(0)
}

func TestCreate_Success(t *testing.T) {
	ctx := context.TODO()
	repositoryMock := new(RepositoryMock)

	service := service.NewRepoService(repositoryMock)
	repoToBeCreated := &public_repo.CreateRepoModel{
		Name:        "Test",
		Description: "Test",
		OwnerID:     primitive.NewObjectID().Hex(),
	}

	repositoryMock.On("Create", ctx, mock.Anything).Return(&model.PrivateRepoModel{}, nil)

	_, err := service.Create(ctx, repoToBeCreated)

	assert.Nil(t, err)
}

func TestCreate_Error(t *testing.T) {
	ctx := context.TODO()
	repositoryMock := new(RepositoryMock)

	service := service.NewRepoService(repositoryMock)
	repoToBeCreated := &public_repo.CreateRepoModel{
		Name:        "Test",
		Description: "Test",
		OwnerID:     primitive.NewObjectID().Hex(),
	}

	repositoryMock.On("Create", ctx, mock.Anything).Return(&model.PrivateRepoModel{}, assert.AnError)

	_, err := service.Create(ctx, repoToBeCreated)

	assert.NotNil(t, err)

	repositoryMock.AssertExpectations(t)
}

func TestFindById_Success(t *testing.T) {
	ctx := context.TODO()
	repositoryMock := new(RepositoryMock)

	service := service.NewRepoService(repositoryMock)
	id := primitive.NewObjectID().Hex()

	repositoryMock.On("FindById", ctx, mock.Anything).Return(&model.PrivateRepoModel{}, nil)

	_, err := service.FindById(ctx, id)

	assert.Nil(t, err)

	repositoryMock.AssertExpectations(t)
}

func TestFindById_Error(t *testing.T) {
	ctx := context.TODO()
	repositoryMock := new(RepositoryMock)

	service := service.NewRepoService(repositoryMock)
	id := primitive.NewObjectID().Hex()

	repositoryMock.On("FindById", ctx, mock.Anything).Return(&model.PrivateRepoModel{}, assert.AnError)

	_, err := service.FindById(ctx, id)

	assert.NotNil(t, err)

	repositoryMock.AssertExpectations(t)
}

func TestFindByName_Success(t *testing.T) {
	ctx := context.TODO()
	repositoryMock := new(RepositoryMock)

	service := service.NewRepoService(repositoryMock)
	name := "Test"

	repositoryMock.On("FindByName", ctx, mock.Anything).Return(&model.PrivateRepoModel{}, nil)

	_, err := service.FindByName(ctx, name)

	assert.Nil(t, err)

	repositoryMock.AssertExpectations(t)
}

func TestFindByName_Error(t *testing.T) {
	ctx := context.TODO()
	repositoryMock := new(RepositoryMock)

	service := service.NewRepoService(repositoryMock)
	name := "Test"

	repositoryMock.On("FindByName", ctx, mock.Anything).Return(&model.PrivateRepoModel{}, assert.AnError)

	_, err := service.FindByName(ctx, name)

	assert.NotNil(t, err)

	repositoryMock.AssertExpectations(t)
}

func TestFindByFindByIdentifier_Id(t *testing.T) {
	ctx := context.TODO()
	repositoryMock := new(RepositoryMock)

	service := service.NewRepoService(repositoryMock)
	id := primitive.NewObjectID().Hex()

	repositoryMock.On("FindById", ctx, mock.Anything).Return(&model.PrivateRepoModel{}, nil)

	_, err := service.FindByFindByIdentifier(ctx, id)

	assert.Nil(t, err)

	repositoryMock.AssertExpectations(t)
}

func TestFindByFindByIdentifier_Name(t *testing.T) {
	ctx := context.TODO()
	repositoryMock := new(RepositoryMock)

	service := service.NewRepoService(repositoryMock)
	name := "Test"

	repositoryMock.On("FindByName", ctx, mock.Anything).Return(&model.PrivateRepoModel{}, nil)

	_, err := service.FindByFindByIdentifier(ctx, name)

	assert.Nil(t, err)

	repositoryMock.AssertExpectations(t)
}

func TestFindByFindByIdentifier_Error(t *testing.T) {
	ctx := context.TODO()
	repositoryMock := new(RepositoryMock)

	service := service.NewRepoService(repositoryMock)
	identifier := "Test"

	repositoryMock.On("FindByName", ctx, mock.Anything).Return(&model.PrivateRepoModel{}, assert.AnError)

	_, err := service.FindByFindByIdentifier(ctx, identifier)

	assert.NotNil(t, err)

	repositoryMock.AssertExpectations(t)
}

func TestUpdate_Success(t *testing.T) {
	ctx := context.TODO()
	repositoryMock := new(RepositoryMock)

	service := service.NewRepoService(repositoryMock)
	repoToBeUpdated := &model.PrivateRepoModel{}

	repositoryMock.On("UpdateOne", ctx, mock.Anything).Return(&model.PrivateRepoModel{}, nil)

	_, err := service.Update(ctx, repoToBeUpdated)

	assert.Nil(t, err)

	repositoryMock.AssertExpectations(t)
}

func TestUpdate_Error(t *testing.T) {
	ctx := context.TODO()
	repositoryMock := new(RepositoryMock)

	service := service.NewRepoService(repositoryMock)
	repoToBeUpdated := &model.PrivateRepoModel{}

	repositoryMock.On("UpdateOne", ctx, mock.Anything).Return(&model.PrivateRepoModel{}, assert.AnError)

	_, err := service.Update(ctx, repoToBeUpdated)

	assert.NotNil(t, err)

	repositoryMock.AssertExpectations(t)
}

func TestDelete_Success(t *testing.T) {
	ctx := context.TODO()
	repositoryMock := new(RepositoryMock)

	service := service.NewRepoService(repositoryMock)
	repoToBeDeleted := &model.PrivateRepoModel{}

	repositoryMock.On("DeleteOne", ctx, mock.Anything).Return(nil)

	err := service.Delete(ctx, repoToBeDeleted)

	assert.Nil(t, err)

	repositoryMock.AssertExpectations(t)
}

func TestDelete_Error(t *testing.T) {
	ctx := context.TODO()
	repositoryMock := new(RepositoryMock)

	service := service.NewRepoService(repositoryMock)
	repoToBeDeleted := &model.PrivateRepoModel{}

	repositoryMock.On("DeleteOne", ctx, mock.Anything).Return(assert.AnError)

	err := service.Delete(ctx, repoToBeDeleted)

	assert.NotNil(t, err)

	repositoryMock.AssertExpectations(t)
}
