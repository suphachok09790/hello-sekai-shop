package playerHandler

import (
	"github.com/suphachok09790/hello-sekai-shop/config"
	"github.com/suphachok09790/hello-sekai-shop/modules/player/playerUsecase"
)

type (
	PlayerHttpHandlerService interface{}

	playerHttpHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerHttpHandler(cfg *config.Config, playerUsecase playerUsecase.PlayerUsecaseService) PlayerHttpHandlerService {
	return &playerHttpHandler{playerUsecase: playerUsecase}
}
