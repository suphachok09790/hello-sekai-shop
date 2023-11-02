package paymentHandler

import (
	"github.com/suphachok09790/hello-sekai-shop/config"
	"github.com/suphachok09790/hello-sekai-shop/modules/payment/paymentUsecase"
)

type (
	PaymentQueueHandlerService interface{}

	paymentQueueHandler struct {
		cfg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentQueueHandlerService {
	return &paymentQueueHandler{
		cfg:            cfg,
		paymentUsecase: paymentUsecase,
	}
}
