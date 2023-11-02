package playerHandler

import (
	"github.com/suphachok09790/hello-sekai-shop/config"
	"github.com/suphachok09790/hello-sekai-shop/modules/player/playerUsecase"
)

type (
	PlayerQueueHandlerService interface{}

	playerQueueHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerQueueHandler(cfg *config.Config, playerUsecase playerUsecase.PlayerUsecaseService) PlayerQueueHandlerService {
	return &playerQueueHandler{playerUsecase: playerUsecase}
}
