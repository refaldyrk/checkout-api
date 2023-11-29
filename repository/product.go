package repository

import (
	"checkout-api/domain/model"
	"checkout-api/exception"

	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

type IProductRepo interface {
	Create(product *model.Product) error
	GetList(category string, search string, page int, limit int, sort_by string, sort_order string) ([]model.Product, error)
	BulkCreate(products []model.Product) error
	GetByNameAndPrice(name string, price float64) (*model.Product, error)
	GetCategoryList() ([]string, error)
	GetByID(id string) (*model.Product, error)
	AppendProductOrders(id string, orders []model.ProductOrder) error
	GetAssociatedProductOrders(id string) ([]model.ProductOrder, error)
}

func NewProductRepo(db *gorm.DB) IProductRepo {
	return &ProductRepo{
		DB: db,
	}
}

func (slf *ProductRepo) Create(product *model.Product) error {
	err := slf.DB.Create(product).Error
	if err == gorm.ErrRecordNotFound {
		return exception.DbObjNotFound
	} else if err != nil {
		return err
	}
	return nil
}

// GetList retrieves a list of products based on the given category, search keyword, page, limit, and sort order.
//
// Parameters:
// - category: a string representing the category of products to filter by. If empty, no category filter is applied.
// - search: a string representing the keyword to search for in the product names. If empty, no search filter is applied.
// - page: an integer representing the page number of the results to retrieve.
// - limit: an integer representing the maximum number of products per page.
// - sort_order: a string representing the sort order of the results. Can be "asc" or "desc".
//
// Returns:
// - products: a slice of model.Product containing the retrieved products.
// - error: an error if any occurred during the retrieval process.
func (slf *ProductRepo) GetList(category string, search string, page int, limit int, sort_by string, sort_order string) ([]model.Product, error) {
	var products []model.Product

	query := slf.DB.Model(&model.Product{})
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}
	err := query.Order("created_at " + sort_order).
		Limit(limit).
		Offset((page - 1) * limit).
		Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (slf *ProductRepo) BulkCreate(products []model.Product) error {
	err := slf.DB.Create(&products).Error
	if err != nil {
		return err
	}
	return nil
}

func (slf *ProductRepo) GetByNameAndPrice(name string, price float64) (*model.Product, error) {
	var product model.Product
	err := slf.DB.Where("name = ? AND price = ?", name, price).First(&product).Error
	if err == gorm.ErrRecordNotFound {
		return nil, exception.DbObjNotFound
	} else if err != nil {
		return nil, err
	}
	return &product, nil
}

func (slf *ProductRepo) GetCategoryList() ([]string, error) {
	var categories []string
	err := slf.DB.Model(&model.Product{}).Distinct("category").Pluck("category", &categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (slf *ProductRepo) GetByID(id string) (*model.Product, error) {
	var product model.Product
	err := slf.DB.First(&product, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, exception.DbObjNotFound
	} else if err != nil {
		return nil, err
	}
	return &product, nil
}

func (slf *ProductRepo) AppendProductOrders(id string, orders []model.ProductOrder) error {
	err := slf.DB.Model(&model.Product{ID: id}).Association("ProductOrders").Append(orders)
	if err != nil {
		return err
	}
	return nil
}

func (slf *ProductRepo) GetAssociatedProductOrders(id string) ([]model.ProductOrder, error) {
	var productOrders []model.ProductOrder
	err := slf.DB.Model(&model.Product{ID: id}).Association("ProductOrders").Find(&productOrders)
	if err != nil {
		return []model.ProductOrder{}, err
	}
	return productOrders, nil
}

func (slf *ProductRepo) Update(product *model.Product) error {
	err := slf.DB.Save(product).Error
	if err != nil {
		return err
	}
	return nil
}
