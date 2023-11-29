package service

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/internal/model"
	"github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/internal/repository"
	public_repo "github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/public"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RepoService interface {
	Create(ctx context.Context, repo *public_repo.CreateRepoModel) (*model.PrivateRepoModel, error)
	FindById(ctx context.Context, id string) (*model.PrivateRepoModel, error)
	FindByName(ctx context.Context, name string) (*model.PrivateRepoModel, error)
	FindByFindByIdentifier(ctx context.Context, identifier string) (*model.PrivateRepoModel, error)
	Update(ctx context.Context, repo *model.PrivateRepoModel) (*model.PrivateRepoModel, error)
	Delete(ctx context.Context, repo *model.PrivateRepoModel) error
}

type RepoServiceImpl struct {
	Repository repository.RepoRepository
}

func NewRepoService(repository repository.RepoRepository) *RepoServiceImpl {
	return &RepoServiceImpl{
		Repository: repository,
	}
}

func (s *RepoServiceImpl) Create(ctx context.Context, repo *public_repo.CreateRepoModel) (*model.PrivateRepoModel, error) {
	repo.Name = normalizeRepoName(repo.Name)
	privateRepo := &model.PrivateRepoModel{
		ID:          primitive.NewObjectID(),
		Name:        normalizeRepoName(repo.Name),
		Description: repo.Description,
		OwnerID:     repo.OwnerID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.Repository.Create(ctx, privateRepo)
}

func (s *RepoServiceImpl) FindById(ctx context.Context, id string) (*model.PrivateRepoModel, error) {
	return s.Repository.FindById(ctx, id)
}

func (s *RepoServiceImpl) FindByName(ctx context.Context, name string) (*model.PrivateRepoModel, error) {
	return s.Repository.FindByName(ctx, name)
}

func (s *RepoServiceImpl) FindByFindByIdentifier(ctx context.Context, identifier string) (*model.PrivateRepoModel, error) {
	_, err := primitive.ObjectIDFromHex(identifier)
	if err == nil {
		return s.Repository.FindById(ctx, identifier)
	}

	return s.Repository.FindByName(ctx, identifier)
}

func (s *RepoServiceImpl) Update(ctx context.Context, repo *model.PrivateRepoModel) (*model.PrivateRepoModel, error) {
	return s.Repository.UpdateOne(ctx, repo)
}

func (s *RepoServiceImpl) Delete(ctx context.Context, repo *model.PrivateRepoModel) error {
	return s.Repository.DeleteOne(ctx, repo)
}

func normalizeRepoName(name string) string {
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ToLower(name)
	re := regexp.MustCompile(`-{2,}`)
	name = re.ReplaceAllString(name, "-")
	return name
}
