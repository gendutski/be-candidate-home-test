package repository

import "github.com/gendutski/be-candidate-home-test/core/entity"

type PromotionRepo interface {
	// get promotion by products
	// will return map[int64] where int64 is product id
	GetPromotionByProducts(products []*entity.Product) (map[int64][]*entity.Promotion, error)
}
