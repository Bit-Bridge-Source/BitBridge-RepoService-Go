package grpc

import (
	"context"

	common_grpc "github.com/Bit-Bridge-Source/BitBridge-CommonService-Go/public/grpc"
	"github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/internal/model"
	"github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/internal/service"
	"github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/proto/pb"
	public_repo "github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/public"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IRepoGRPCServer interface {
	CreateRepo(ctx context.Context, req *pb.CreateRepoRequest) (*pb.PrivateRepoResponse, error)
	GetPrivateRepo(ctx context.Context, req *pb.IdentifierRequest) (*pb.PrivateRepoResponse, error)
	GetPublicRepo(ctx context.Context, req *pb.IdentifierRequest) (*pb.PublicRepoResponse, error)
	GetPrivateRepos(ctx context.Context, req *pb.IdentifierRequest) (*pb.PrivateReposResponse, error)
	GetPublicRepos(ctx context.Context, req *pb.IdentifierRequest) (*pb.PublicReposResponse, error)
	UpdateRepo(ctx context.Context, req *pb.PrivateRepoRequest) (*pb.PrivateRepoResponse, error)
	DeleteRepo(ctx context.Context, req *pb.IdentifierRequest) (*emptypb.Empty, error)
}

type RepoGRPCServer struct {
	RepoService  service.RepoService
	Interceptors []grpc.UnaryServerInterceptor
	Config       common_grpc.ServerConfig
	pb.UnimplementedRepoServiceServer
}

func NewRepoGRPCServer(repoService service.RepoService, interceptors []grpc.UnaryServerInterceptor) *RepoGRPCServer {
	return &RepoGRPCServer{
		RepoService:  repoService,
		Interceptors: interceptors,
	}
}

func (s *RepoGRPCServer) InitServer(port string, listener common_grpc.Listener) error {
	lis, err := listener.Listen("tcp", port)
	if err != nil {
		return err
	}
	s.Config.Listener = lis
	s.Config.GRPCServer = grpc.NewServer(
		grpc.UnaryInterceptor(common_grpc.ChainUnaryInterceptors(s.Interceptors...)),
	)
	pb.RegisterRepoServiceServer(s.Config.GRPCServer, s)
	return nil
}

func (s *RepoGRPCServer) Run() error {
	return s.Config.GRPCServer.Serve(s.Config.Listener)
}

func (s *RepoGRPCServer) CreateRepo(ctx context.Context, req *pb.CreateRepoRequest) (*pb.PrivateRepoResponse, error) {
	repo, err := s.RepoService.Create(ctx, &public_repo.CreateRepoModel{
		Name:        req.Name,
		Description: req.Description,
		OwnerID:     req.OwnerId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PrivateRepoResponse{
		Id:          repo.ID.Hex(),
		Name:        repo.Name,
		Description: repo.Description,
		OwnerId:     repo.OwnerID,
		CreatedAt:   repo.CreatedAt.String(),
		UpdatedAt:   repo.UpdatedAt.String(),
	}, nil
}

func (s *RepoGRPCServer) GetPrivateRepo(ctx context.Context, req *pb.IdentifierRequest) (*pb.PrivateRepoResponse, error) {
	repo, err := s.RepoService.FindByFindByIdentifier(ctx, req.RepoIdentifier)
	if err != nil {
		return nil, err
	}
	return &pb.PrivateRepoResponse{
		Id:          repo.ID.Hex(),
		Name:        repo.Name,
		Description: repo.Description,
		OwnerId:     repo.OwnerID,
		CreatedAt:   repo.CreatedAt.String(),
		UpdatedAt:   repo.UpdatedAt.String(),
	}, nil
}

func (s *RepoGRPCServer) GetPublicRepo(ctx context.Context, req *pb.IdentifierRequest) (*pb.PublicRepoResponse, error) {
	repo, err := s.RepoService.FindByFindByIdentifier(ctx, req.RepoIdentifier)
	if err != nil {
		return nil, err
	}
	return &pb.PublicRepoResponse{
		Id:          repo.ID.Hex(),
		Name:        repo.Name,
		Description: repo.Description,
		CreatedAt:   repo.CreatedAt.String(),
		UpdatedAt:   repo.UpdatedAt.String(),
	}, nil
}

func (s *RepoGRPCServer) GetPrivateRepos(ctx context.Context, req *pb.IdentifierRequest) (*pb.PrivateReposResponse, error) {
	var privateRepos []*model.PrivateRepoModel
	_, err := primitive.ObjectIDFromHex(req.RepoIdentifier)
	if err == nil {
		privateRepos, err = s.RepoService.FindAllByOwnerID(ctx, req.RepoIdentifier, req.Page, req.PageSize)
	} else {
		privateRepos, err = s.RepoService.FindAllByName(ctx, req.RepoIdentifier, req.Page, req.PageSize)
	}

	if err != nil {
		return nil, err
	}

	var publicRepos []*pb.PrivateRepoResponse
	for _, repo := range privateRepos {
		publicRepos = append(publicRepos, &pb.PrivateRepoResponse{
			Id:          repo.ID.Hex(),
			Name:        repo.Name,
			Description: repo.Description,
			OwnerId:     repo.OwnerID,
			CreatedAt:   repo.CreatedAt.String(),
			UpdatedAt:   repo.UpdatedAt.String(),
		})
	}

	return &pb.PrivateReposResponse{
		Repos: publicRepos,
	}, nil
}

func (s *RepoGRPCServer) GetPublicRepos(ctx context.Context, req *pb.IdentifierRequest) (*pb.PublicReposResponse, error) {
	var privateRepos []*model.PrivateRepoModel
	_, err := primitive.ObjectIDFromHex(req.RepoIdentifier)
	if err == nil {
		privateRepos, err = s.RepoService.FindAllByOwnerID(ctx, req.RepoIdentifier, req.Page, req.PageSize)
	} else {
		privateRepos, err = s.RepoService.FindAllByName(ctx, req.RepoIdentifier, req.Page, req.PageSize)
	}

	if err != nil {
		return nil, err
	}

	var publicRepos []*pb.PublicRepoResponse
	for _, repo := range privateRepos {
		publicRepos = append(publicRepos, &pb.PublicRepoResponse{
			Id:          repo.ID.Hex(),
			Name:        repo.Name,
			Description: repo.Description,
			CreatedAt:   repo.CreatedAt.String(),
			UpdatedAt:   repo.UpdatedAt.String(),
		})
	}

	return &pb.PublicReposResponse{
		Repos: publicRepos,
	}, nil
}
