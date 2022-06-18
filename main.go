package main

import (
	"context"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/charCharacter/history/relations/models"
	"github.com/charCharacter/history/relations/packages/utils"
	"github.com/charCharacter/history/relations/repository/tarantool"

	"github.com/charCharacter/history/relations/configs"
	pb "github.com/charCharacter/history/relations/pb"
)

type RelationsServer struct {
	pb.UnimplementedRelationsServiceServer
	JWT  utils.JwtWrapper
	repo *tarantool.Repo
}

func NewRelationsServer(repo *tarantool.Repo, jwt utils.JwtWrapper) (*RelationsServer, error) {

	return &RelationsServer{
		repo: repo,
		JWT:  jwt,
	}, nil
}

type LoginError string

func (err LoginError) Error() string {
	return string(err)
}
func (s *RelationsServer) ArticleCreate(ctx context.Context, in *pb.ArticleCreateRequest) (*pb.ArticleCreateResponse, error) {

	resp, err := s.repo.ArticleCreate(&models.Article{
		Title:      in.ArticleTitle,
		ArticleUID: in.ArticleUid,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ArticleCreateResponse{
		ArticleTitle: resp.Title,
		ArticleUid:   resp.ArticleUID,
	}, nil
}

func (s *RelationsServer) ArticleFind(ctx context.Context, in *pb.ArticleFindRequest) (*pb.ArticleFindResponse, error) {
	resp, resp2, err := s.repo.ArticleGet(in.ArticleUid)
	if err != nil {
		return nil, err
	}
	log.Infoln(resp2)
	return &pb.ArticleFindResponse{
		ArticleTitle: resp.Title,
		ArticleUid:   resp.ArticleUID,
	}, nil
}

func (s *RelationsServer) RelationsCreateKeywords(ctx context.Context, in *pb.RelationsCreateKeywordsRequest) (*pb.RelationsCreateKeywordsResponse, error) {
	_, err := s.repo.RelationsCreate(&models.ArticleKeywords{
		ArticleUID: in.ArticleUid,
		Keywords:   in.Keywords,
	})
	if err != nil {
		return nil, err
	}
	return &pb.RelationsCreateKeywordsResponse{
		ArticleUid: in.ArticleUid,
	}, nil
}

func (s *RelationsServer) RelationsList(ctx context.Context, in *emptypb.Empty) (*pb.RelationsListResponse, error) {
	res, err := s.repo.RelationsList()
	if err != nil {
		return nil, err
	}
	log.Infoln(res)
	return nil, nil
}
func (s *RelationsServer) ArticleFindAll(ctx context.Context, in *emptypb.Empty) (*pb.ArticleFindAllResponse, error) {
	res, err := s.repo.ArticleList()
	if err != nil {
		return nil, err
	}
	log.Infoln(res)
	return nil, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("can't load app config with error [%s]", err)
	}
	level, err := log.ParseLevel(config.Logger.Level)
	if err != nil {
		level = log.WarnLevel
	}
	log.SetLevel(level)
	lis, err := net.Listen("tcp", config.GRPC.Address)
	if err != nil {
		log.Fatalf("Failed to listen: [%s]", err)
	}
	db, err := tarantool.New(config.Tarantool.Address, config.Tarantool.User, config.Tarantool.Pass)
	if err != nil {
		log.Fatalf("Failed to listen: [%s]", err)
	}
	jwt := utils.JwtWrapper{
		SecretKey:       config.JWT.SecretKey,
		Issuer:          config.JWT.Issuer,
		ExpirationHours: config.JWT.ExpirationHours,
	}
	serv, err := NewRelationsServer(&db, jwt)
	if err != nil {
		log.Fatalf("Failed to start server: [%s]", err)
	}
	s := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			//Relations.StreamServerInterceptor(myRelationsFunction),
			recovery.StreamServerInterceptor(),
			validator.StreamServerInterceptor(),
		),
		grpc.ChainUnaryInterceptor(
			//Relations.UnaryServerInterceptor(myRelationsFunction),
			recovery.UnaryServerInterceptor(),
			validator.UnaryServerInterceptor(),
		),
	)
	pb.RegisterRelationsServiceServer(s, serv)
	log.Infoln("Serving gRPC on ", config.GRPC.Address)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		config.GRPC.Address,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to dial server: [%s]", err)
	}

	gwmux := runtime.NewServeMux()
	err = pb.RegisterRelationsServiceHandler(context.Background(), gwmux, conn)

	if err != nil {
		log.Fatalf("Failed to register gateway: [%s]", err)
	}

	gwServer := &http.Server{
		Addr:        config.Rest.Address,
		Handler:     gwmux,
		ReadTimeout: config.Rest.Timeout,
	}
	log.Infoln("Serving gRPC-Gateway on ", config.Rest.Address)
	log.Fatalln(gwServer.ListenAndServe())

}
