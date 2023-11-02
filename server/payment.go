package server

import (
	"github.com/suphachok09790/hello-sekai-shop/modules/payment/paymentHandler"
	"github.com/suphachok09790/hello-sekai-shop/modules/payment/paymentRepository"
	"github.com/suphachok09790/hello-sekai-shop/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	repo := paymentRepository.NewPaymentRepository(s.db)
	usecase := paymentUsecase.NewPaymentUsecase(repo)
	httpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, usecase)
	queueHandler := paymentHandler.NewPaymentQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = queueHandler

	payment := s.app.Group("/payment_v1")

	// Health Check
	payment.GET("", s.healthCheckService)
}
