package paymentHandler

import (
	"github.com/suphachok09790/hello-sekai-shop/config"
	"github.com/suphachok09790/hello-sekai-shop/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandlerService interface{}

	paymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config,paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentHttpHandlerService {
	return &paymentHttpHandler{
		cfg: cfg,
		paymentUsecase: paymentUsecase,
	}
}
