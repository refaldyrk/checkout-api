package service

import (
	"checkout-api/domain/model"
	"checkout-api/exception"
	"checkout-api/repository"
	error_utils "checkout-api/utils/error"
	"checkout-api/utils/helper"
)

type ProductService struct {
	productRepo      repository.IProductRepo
	cartRepo         repository.ICartRepo
	productOrderRepo repository.IProductOrderRepo
}

type IProductService interface {
	GetList(category string, search string, page int, limit int, sort_by string, sort_order string) ([]model.Product, error)
	GetCategoryList() ([]string, error)
	AddItemToCart(product_id string, quantity int, user_id string) error
}

func NewProductService(productRepo repository.IProductRepo, cartRepo repository.ICartRepo, productOrderRepo repository.IProductOrderRepo) IProductService {
	return &ProductService{
		productRepo:      productRepo,
		cartRepo:         cartRepo,
		productOrderRepo: productOrderRepo,
	}
}

func (slf *ProductService) GetList(category string, search string, page int, limit int, sort_by string, sort_order string) ([]model.Product, error) {
	products, err := slf.productRepo.GetList(category, search, page, limit, sort_by, sort_order)
	if err != nil {
		return []model.Product{}, err
	}
	return products, nil
}

func (slf *ProductService) GetCategoryList() ([]string, error) {
	categories, err := slf.productRepo.GetCategoryList()
	if err != nil {
		return []string{}, err
	}

	return categories, nil
}

func (slf *ProductService) AddItemToCart(product_id string, quantity int, user_id string) error {
	// check if product exists
	_, err := slf.productRepo.GetByID(product_id)
	if err != nil {
		if err == exception.DbObjNotFound {
			return &error_utils.CustomErr{
				Code:    404,
				Message: "product not found",
			}
		} else {
			return err
		}
	}

	// get latest unchecked-out cart or create if it doesn't exist
	cart, err := slf.cartRepo.GetLast(false, user_id)
	var newCart *model.Cart
	if err != nil {
		if err == exception.DbObjNotFound {
			newCart = &model.Cart{
				ID:     helper.GenerateUUID(),
				UserID: user_id,
			}
			err = slf.cartRepo.Create(newCart)
			if err != nil {
				return err
			}
			cart = newCart
		} else {
			return err
		}
	}

	// check if product is already in cart
	cartAssociatedProductOrder, err := slf.cartRepo.GetAssociatedProductOrders(cart.ID)
	if err != nil {
		return err
	}
	for _, productOrder := range cartAssociatedProductOrder {
		if productOrder.ProductID == product_id {
			return &error_utils.CustomErr{
				Code:    400,
				Message: "product already added to cart",
			}
		}
	}

	// create product order
	err = slf.productOrderRepo.Create(&model.ProductOrder{
		ID:        helper.GenerateUUID(),
		CartID:    cart.ID,
		ProductID: product_id,
		Quantity:  quantity,
	})
	if err != nil {
		return err
	}

	return nil
}
