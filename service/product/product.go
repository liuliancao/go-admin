package productservice

import "go-admin/models"

type Product struct {
	ID          uint
	Name        string
	Description string
	Parent      uint
	Status      uint
	CreatedBy   uint
	UpdatedBy   uint
}

type ProductUser struct {
	ID        uint `json:"id"`
	PID       uint `json:"pid"`
	UID       uint `json:"uid"`
	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
}

func (p *Product) Add() (id uint, err error) {
	product := map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"parent":      p.Parent,
		"status":      p.Status,
		"created_by":  p.CreatedBy,
	}

	id, err = models.AddProduct(product)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *Product) Edit() error {
	return models.EditProduct(p.ID, map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"parent":      p.Parent,
		"status":      p.Status,
		"updated_by":  p.UpdatedBy,
	})
}

func (p *Product) Get() (*models.Product, error) {
	product, err := models.GetProduct(p.ID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) GetAll() ([]*models.Product, error) {
	var products []*models.Product
	products, err := models.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *Product) ExistByID() (bool, error) {
	return models.ExistProductByID(p.ID)
}

func (p *Product) Delete() error {
	return models.DeleteProduct(p.ID)
}
func (pu *ProductUser) Add() error {
	return models.AddProductUser(pu.PID, pu.UID, pu.CreatedBy)
}
func (pu *ProductUser) Get(pid uint) ([]*models.User, error) {
	return models.GetProductUsers(pid)
}
