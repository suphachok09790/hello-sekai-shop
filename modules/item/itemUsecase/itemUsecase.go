package itemUsecase

import "github.com/suphachok09790/hello-sekai-shop/modules/item/itemRepository"

type (
	ItemUsecaseService interface{}

	itemUsecase struct {
		itemRepository itemRepository.ItemRepositoryService
	}
)

func NewItemUsecase(itemRepository itemRepository.ItemRepositoryService) ItemUsecaseService {
	return &itemUsecase{itemRepository}
}
