package departmentservice

import (
	"go-admin/models"
)

type Department struct {
	ID          uint
	Name        string
	Description string
	Status      uint
	Parent      uint
	CreatedBy   uint
	UpdatedBy   uint
}

func (d *Department) Add() error {
	department := map[string]interface{}{
		"name":        d.Name,
		"description": d.Description,
		"parent":      d.Parent,
		"status":      d.Status,
		"created_by":  d.CreatedBy,
	}

	if err := models.AddDepartment(department); err != nil {
		return err
	}

	return nil
}
func (d *Department) Edit() error {
	return models.EditDepartment(d.ID, map[string]interface{}{
		"name":        d.Name,
		"description": d.Description,
		"parent":      d.Parent,
		"status":      d.Status,
		"updated_by":  d.UpdatedBy,
	})
}

func (d *Department) Get() (*models.Department, error) {
	department, err := models.GetDepartment(d.ID)
	if err != nil {
		return nil, err
	}

	return department, nil
}

func (d *Department) GetAll() ([]*models.Department, error) {
	var departments []*models.Department
	departments, err := models.GetDepartments()
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (d *Department) ExistByID() (bool, error) {
	return models.ExistDepartmentByID(d.ID)
}

func (d *Department) Delete() error {
	return models.DeleteDepartment(d.ID)
}
