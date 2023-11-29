package service

import (
	"checkout-api/domain/model"
	"checkout-api/exception"
	"checkout-api/repository"
	error_utils "checkout-api/utils/error"
	"time"
)

type CartService struct {
	cartRepo         repository.ICartRepo
	productRepo      repository.IProductRepo
	productOrderRepo repository.IProductOrderRepo
}

type ICartService interface {
	GetCartItems(user_id string) ([]model.ProductOrder, error)
	RemoveItemFromCart(order_id string, user_id string) error
	CheckoutCart(user_id string, money_input float64) (float64, error)
}

func NewCartService(cartRepo repository.ICartRepo, productRepo repository.IProductRepo, productOrderRepo repository.IProductOrderRepo) *CartService {
	return &CartService{
		cartRepo:         cartRepo,
		productRepo:      productRepo,
		productOrderRepo: productOrderRepo,
	}
}

func (slf *CartService) GetCartItems(user_id string) ([]model.ProductOrder, error) {
	// get latest unchecked-out cart
	cart, err := slf.cartRepo.GetLast(false, user_id)
	if err != nil {
		if err == exception.DbObjNotFound {
			return []model.ProductOrder{}, &error_utils.CustomErr{
				Code:    400,
				Message: "no items ordered",
			}
		} else {
			return []model.ProductOrder{}, err
		}
	}

	// get all associated product orders
	orders, err := slf.cartRepo.GetAssociatedProductOrders(cart.ID)
	if err != nil {
		return []model.ProductOrder{}, err
	}

	return orders, nil
}

func (slf *CartService) RemoveItemFromCart(order_id string, user_id string) error {
	// get latest unchecked-out cart
	cart, err := slf.cartRepo.GetLast(false, user_id)
	if err != nil {
		if err == exception.DbObjNotFound {
			return &error_utils.CustomErr{
				Code:    400,
				Message: "no items ordered",
			}
		} else {
			return err
		}
	}

	// check if product is in cart
	orders, err := slf.cartRepo.GetAssociatedProductOrders(cart.ID)
	if err != nil {
		return err
	}
	var orderTarget model.ProductOrder
	for _, order := range orders {
		if order.ID == order_id {
			orderTarget = order
			break
		}
	}
	if orderTarget.ID == "" {
		return &error_utils.CustomErr{
			Code:    404,
			Message: "product is not in cart",
		}
	}

	// remove order from cart association
	err = slf.cartRepo.RemoveOrderAssociations(cart.ID, []model.ProductOrder{orderTarget})
	if err != nil {
		return err
	}

	// delete order
	err = slf.productOrderRepo.Delete(orderTarget.ID)
	if err != nil {
		return err
	}

	// delete cart if orders is empty
	orders, err = slf.cartRepo.GetAssociatedProductOrders(cart.ID)
	if err != nil {
		return err
	}
	if len(orders) == 0 {
		err = slf.cartRepo.Delete(cart.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (slf *CartService) CheckoutCart(user_id string, money_input float64) (float64, error) {
	// get latest unchecked-out cart
	cart, err := slf.cartRepo.GetLast(false, user_id)
	if err != nil {
		if err == exception.DbObjNotFound {
			return 0, &error_utils.CustomErr{
				Code:    400,
				Message: "no items ordered",
			}
		} else {
			return 0, err
		}
	}

	// check if product is in cart
	orders, err := slf.cartRepo.GetAssociatedProductOrders(cart.ID)
	if err != nil {
		return 0, err
	}
	if len(orders) == 0 {
		return 0, &error_utils.CustomErr{
			Code:    400,
			Message: "cart is empty",
		}
	}

	// sum total price
	totalPrice := 0.0
	for _, order := range orders {
		totalPrice += float64(order.Quantity) * order.Product.Price
	}

	// check if money is enough
	if totalPrice > money_input {
		return 0, &error_utils.CustomErr{
			Code:    400,
			Message: "money is not enough",
		}
	}

	// total money return
	totalMoneyReturn := money_input - totalPrice

	// update cart
	cart.IsCheckout = true
	cart.CheckedOutAt = time.Now()
	err = slf.cartRepo.Save(cart)
	if err != nil {
		return 0, err
	}

	return totalMoneyReturn, nil
}
