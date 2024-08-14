package server

import (
	"log"

	"github.com/suphachok09790/hello-sekai-shop/modules/item/itemHandler"
	itemPb "github.com/suphachok09790/hello-sekai-shop/modules/item/itemPb"
	"github.com/suphachok09790/hello-sekai-shop/modules/item/itemRepository"
	"github.com/suphachok09790/hello-sekai-shop/modules/item/itemUsecase"
	"github.com/suphachok09790/hello-sekai-shop/pkg/grpccon"
)

func (s *server) itemService() {
	repo := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repo)
	httpHandler := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)

		itemPb.RegisterItemGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Item gRPC server listing on %s", s.cfg.Grpc.ItemUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	// Health Check
	item.GET("", s.healthCheckService)
}
