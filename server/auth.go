package server

import (
	"github.com/suphachok09790/hello-sekai-shop/modules/auth/authHandler"
	"github.com/suphachok09790/hello-sekai-shop/modules/auth/authRepository"
	"github.com/suphachok09790/hello-sekai-shop/modules/auth/authUsecase"
)

func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	// Health Check
	auth.GET("", s.healthCheckService)
}
