package repository

import "Store/internal/oas"

// get all products
func (r *Repository) GetProducts() ([]oas.Product, error) {
	var products []oas.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// get product by id
func (r *Repository) GetProductById(id int) (*oas.Product, error) {
	var product oas.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// create product
func (r *Repository) CreateProduct(product *oas.Product) error {
	err := r.db.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

// delete product
func (r *Repository) DeleteProduct(id int) error {
	err := r.db.Delete(&oas.Product{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// update product
func (r *Repository) UpdateProduct(product *oas.Product) error {
	err := r.db.Save(&product).Error
	if err != nil {
		return err
	}
	return nil
}
