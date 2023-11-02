package inventoryUsecase

import "github.com/suphachok09790/hello-sekai-shop/modules/inventory/inventoryRepository"

type (
	InventoryUsecaseService interface{}

	inventoryUsecase struct {
		inventoryRepository inventoryRepository.InventoryRepositoryService
	}
)

func NewInventoryUsecase(inventoryRepository inventoryRepository.InventoryRepositoryService) InventoryUsecaseService{
	return &inventoryUsecase{inventoryRepository:inventoryRepository}
}