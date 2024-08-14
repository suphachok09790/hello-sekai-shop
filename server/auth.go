package server

import (
	"log"

	"github.com/suphachok09790/hello-sekai-shop/modules/auth/authHandler"
	authPb "github.com/suphachok09790/hello-sekai-shop/modules/auth/authPb"
	"github.com/suphachok09790/hello-sekai-shop/modules/auth/authRepository"
	"github.com/suphachok09790/hello-sekai-shop/modules/auth/authUsecase"
	"github.com/suphachok09790/hello-sekai-shop/pkg/grpccon"
)

func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Auth gRPC server listing on %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	// Health Check
	auth.GET("", s.healthCheckService)
}
