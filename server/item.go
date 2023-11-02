package server

import (
	"github.com/suphachok09790/hello-sekai-shop/modules/item/itemHandler"
	"github.com/suphachok09790/hello-sekai-shop/modules/item/itemRepository"
	"github.com/suphachok09790/hello-sekai-shop/modules/item/itemUsecase"
)

func (s *server) itemService() {
	repo := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repo)
	httpHandler := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	// Health Check
	item.GET("", s.healthCheckService)
}
