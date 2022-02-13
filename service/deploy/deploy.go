package deployservice

import (
	"go-admin/models"
)

type HostDeploy struct {
	ID        uint
	AID       uint
	HID       uint
	Status    string
	Version   string
	Stage     string
	Message   string
	AppMacros string
	CreatedBy uint
	UpdatedBy uint
}

func (d *HostDeploy) Add() error {
	hostdeploy := map[string]interface{}{
		"aid":        d.AID,
		"hid":        d.HID,
		"version":    d.Version,
		"message":    d.Message,
		"stage":      d.Stage,
		"app_macros": d.AppMacros,
		"status":     d.Status,
		"created_by": d.CreatedBy,
	}

	if err := models.AddHostDeploy(hostdeploy); err != nil {
		return err
	}

	return nil
}
func (d *HostDeploy) Edit() error {
	return models.EditHostDeploy(d.ID, map[string]interface{}{
		"aid":        d.AID,
		"hid":        d.HID,
		"stage":      d.Stage,
		"version":    d.Version,
		"message":    d.Message,
		"app_macros": d.AppMacros,
		"status":     d.Status,
		"updated_by": d.UpdatedBy,
	})
}

func (d *HostDeploy) Get() (*models.HostDeploy, error) {
	hostdeploy, err := models.GetHostDeploy(d.ID)
	if err != nil {
		return nil, err
	}

	return hostdeploy, nil
}

func (d *HostDeploy) GetAll() ([]*models.HostDeploy, error) {
	return models.GetHostDeploys(map[string]interface{}{
		"a_id":       d.AID,
		"h_id":       d.HID,
		"stage":      d.Stage,
		"version":    d.Version,
		"status":     d.Status,
		"created_by": d.CreatedBy,
	})
}

func (d *HostDeploy) ExistByID() (bool, error) {
	return models.ExistHostDeployByID(d.ID)
}

func (d *HostDeploy) Delete() error {
	return models.DeleteHostDeploy(d.ID)
}
