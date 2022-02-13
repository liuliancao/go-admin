package resourceservice

import "go-admin/models"

type ResourceType struct {
	ID           uint
	Name         string `json:"name"`
	Nickname     string `json:"nickname"`
	Tag          string `json:"tag"`
	Description  string `json:"desciption"`
	HTMLFormJSON string `json:"html_form_json"`

	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
}

type Resource struct {
	ID uint

	Name        string `json:"name"`
	TID         uint   `json:"t_id"`
	Description string `json:"description"`
	DataJSON    string `json:"data_json"`

	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
}

func (rt *ResourceType) Add() (id uint, err error) {
	resourceType := map[string]interface{}{
		"name":           rt.Name,
		"nickname":       rt.Nickname,
		"tag":            rt.Tag,
		"description":    rt.Description,
		"html_form_json": rt.HTMLFormJSON,
		"created_by":     rt.CreatedBy,
	}

	id, err = models.AddResourceType(resourceType)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (rt *ResourceType) Edit() error {
	return models.EditResourceType(rt.ID, map[string]interface{}{
		"name":           rt.Name,
		"nickname":       rt.Nickname,
		"tag":            rt.Tag,
		"description":    rt.Description,
		"html_form_json": rt.HTMLFormJSON,
		"updated_by":     rt.UpdatedBy,
	})
}

func (rt *ResourceType) Get() (*models.ResourceType, error) {
	resourceType, err := models.GetResourceType(rt.ID)
	if err != nil {
		return nil, err
	}

	return resourceType, nil
}

func (rt *ResourceType) GetAll() ([]*models.ResourceType, error) {
	var resourceTypes []*models.ResourceType
	resourceTypes, err := models.GetResourceTypes()
	if err != nil {
		return nil, err
	}
	return resourceTypes, nil
}

func (rt *ResourceType) ExistByID() (bool, error) {
	return models.ExistResourceTypeByID(rt.ID)
}
func (rt *ResourceType) Delete() error {
	return models.DeleteResourceType(rt.ID)
}

func (r *Resource) Add() (id uint, err error) {
	resource := map[string]interface{}{
		"name":        r.Name,
		"t_id":        r.TID,
		"description": r.Description,
		"data_json":   r.DataJSON,
		"created_by":  r.CreatedBy,
	}

	id, err = models.AddResource(resource)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Resource) Edit() error {
	return models.EditResource(r.ID, map[string]interface{}{
		"name":        r.Name,
		"t_id":        r.TID,
		"description": r.Description,
		"data_json":   r.DataJSON,
		"updated_by":  r.UpdatedBy,
	})
}

func (r *Resource) Get() (*models.Resource, error) {
	resource, err := models.GetResource(r.ID)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (r *Resource) GetAll() ([]*models.Resource, error) {
	var resources []*models.Resource
	resources, err := models.GetResources()
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func (r *Resource) ExistByID() (bool, error) {
	return models.ExistResourceByID(r.ID)
}
func (r *Resource) Delete() error {
	return models.DeleteResource(r.ID)
}
