package tag_service

import "go-admin/models"

type Tag struct {
	ID          uint
	Name        string
	Description string
	CreatedBy   string
	UpdatedBy   string
}

func (t *Tag) Add() error {
	tag := map[string]interface{}{
		"name":        t.Name,
		"description": t.Description,
		"created_by":  t.CreatedBy,
	}

	if err := models.AddTag(tag); err != nil {
		return err
	}

	return nil
}

func (t *Tag) Edit() error {
	return models.EditTag(t.ID, map[string]interface{}{
		"name":        t.Name,
		"description": t.Description,
		"updated_by":  t.UpdatedBy,
	})
}

func (t *Tag) Get() (*models.Tag, error) {
	tag, err := models.GetTag(t.ID)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (t *Tag) GetAll() ([]*models.Tag, error) {
	var tags []*models.Tag
	tags, err := models.GetTags()
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}
