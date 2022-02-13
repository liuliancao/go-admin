package macroservice

import (
	"go-admin/models"
)

type Macro struct {
	ID          uint
	Namespace   string
	Key         string
	Name        string
	MType       string
	Value       string
	Encrypt     uint
	Description string
	CreatedBy   uint
	UpdatedBy   uint
}

func (m *Macro) Add() error {
	macro := map[string]interface{}{
		"namespace":   m.Namespace,
		"key":         m.Key,
		"name":        m.Name,
		"m_type":      m.MType,
		"encrypt":     m.Encrypt,
		"value":       m.Value,
		"description": m.Description,
		"created_by":  m.CreatedBy,
	}

	if err := models.AddMacro(macro); err != nil {
		return err
	}

	return nil
}
func (m *Macro) Edit() error {
	return models.EditMacro(m.ID, map[string]interface{}{
		"namespace":   m.Namespace,
		"key":         m.Key,
		"name":        m.Name,
		"m_type":      m.MType,
		"encrypt":     m.Encrypt,
		"value":       m.Value,
		"description": m.Description,
		"updated_by":  m.UpdatedBy,
	})
}

func (m *Macro) Get() (*models.Macro, error) {
	macro, err := models.GetMacro(m.ID)
	if err != nil {
		return nil, err
	}

	return macro, nil
}

func (m *Macro) GetAll(namespace, mType string) ([]*models.Macro, error) {
	var macros []*models.Macro
	macros, err := models.GetMacros(map[string]interface{}{
		"namespace": namespace,
		"m_type":    mType,
	})
	if err != nil {
		return nil, err
	}
	return macros, nil
}

func (m *Macro) ExistByID() (bool, error) {
	return models.ExistMacroByID(m.ID)
}
func (m *Macro) ExistByNamespaceKey() (uint, error) {
	return models.ExistMacroByNamespaceKey(m.Namespace, m.Key)
}

func (m *Macro) Delete() error {
	return models.DeleteMacro(m.ID)
}
