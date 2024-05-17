package repository

import "github.com/gendutski/be-candidate-home-test/core/entity"

type ProductRepo interface {
	GetProductBySerials(serials []string) ([]*entity.Product, error)
	GetProductByIDs(ids []int64) ([]*entity.Product, error)
	SubmitCheckout(payload *entity.Checkout) error
}
