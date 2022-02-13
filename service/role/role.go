package roleservice

import "go-admin/models"

type Role struct {
	ID          uint
	Name        string
	Description string
	Status      uint
	CreatedBy   uint
	UpdatedBy   uint
}

func (r *Role) Add() error {
	role := map[string]interface{}{
		"name":        r.Name,
		"description": r.Description,
		"status":      r.Status,
		"created_by":  r.CreatedBy,
	}

	if err := models.AddRole(role); err != nil {
		return err
	}

	return nil
}

func (r *Role) Edit() error {
	return models.EditRole(r.ID, map[string]interface{}{
		"name":        r.Name,
		"description": r.Description,
		"status":      r.Status,
		"updated_by":  r.UpdatedBy,
	})
}

func (r *Role) Get() (*models.Role, error) {
	role, err := models.GetRole(r.ID)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *Role) GetAll() ([]*models.Role, error) {
	var roles []*models.Role
	roles, err := models.GetRoles()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *Role) ExistByID() (bool, error) {
	return models.ExistRoleByID(r.ID)
}

func (r *Role) Delete() error {
	return models.DeleteRole(r.ID)
}
