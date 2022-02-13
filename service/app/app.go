package appservice

import "go-admin/models"

type App struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PID         uint   `json:"pid"`
	EID         uint   `json:"eid"`
	Parent      uint   `json:"parent"`
	Description string `json:"description"`
	Status      uint   `json:"status"`
	CreatedBy   uint   `json:"created_by"`
	UpdatedBy   uint   `json:"updated_by"`
}

type AppEnv struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   uint   `json:"created_by"`
	UpdatedBy   uint   `json:"updated_by"`
}

type AppUser struct {
	AID       uint `json:"a_id"`
	UID       uint `json:"u_id"`
	CreatedBy uint `json:"created_by"`
}

type AppCluster struct {
	AID       uint `json:"aid"`
	CID       uint `json:"cid"`
	CreatedBy uint `json:"created_by"`
}

func (a *App) Add() error {
	app := map[string]interface{}{
		"name":        a.Name,
		"pid":         a.PID,
		"eid":         a.EID,
		"parent":      a.Parent,
		"description": a.Description,
		"status":      a.Status,
		"created_by":  a.CreatedBy,
	}

	if err := models.AddApp(app); err != nil {
		return err
	}

	return nil
}

func (a *App) Edit() error {
	return models.EditApp(a.ID, map[string]interface{}{
		"name":        a.Name,
		"description": a.Description,
		"parent":      a.Parent,
		"pid":         a.PID,
		"eid":         a.EID,
		"status":      a.Status,
		"updated_by":  a.UpdatedBy,
	})
}

func (a *App) Get() (*models.App, error) {
	app, err := models.GetApp(a.ID)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) GetAll() ([]*models.App, error) {
	return models.GetApps(map[string]interface{}{
		"pid":    a.PID,
		"eid":    a.EID,
		"parent": a.Parent,
		"status": a.Status,
	})
}

func (a *App) ExistByID() (bool, error) {
	return models.ExistAppByID(a.ID)
}

func (a *App) Delete() error {
	return models.DeleteApp(a.ID)
}
func (ae *AppEnv) Add() error {
	appenv := map[string]interface{}{
		"name":        ae.Name,
		"description": ae.Description,
		"created_by":  ae.CreatedBy,
	}

	if err := models.AddAppEnv(appenv); err != nil {
		return err
	}

	return nil
}

func (ae *AppEnv) Edit() error {
	return models.EditAppEnv(ae.ID, map[string]interface{}{
		"name":        ae.Name,
		"description": ae.Description,
		"updated_by":  ae.UpdatedBy,
	})
}

func (ae *AppEnv) Get() (*models.AppEnv, error) {
	appenv, err := models.GetAppEnv(ae.ID)
	if err != nil {
		return nil, err
	}

	return appenv, nil
}

func (ae *AppEnv) GetAll() ([]*models.AppEnv, error) {
	var appenvs []*models.AppEnv
	appenvs, err := models.GetAppEnvs()
	if err != nil {
		return nil, err
	}
	return appenvs, nil
}

func (ae *AppEnv) ExistByID() (bool, error) {
	return models.ExistAppEnvByID(ae.ID)
}

func (ae *AppEnv) Delete() error {
	return models.DeleteAppEnv(ae.ID)
}

func (ac *AppCluster) Add() error {
	return models.AddAppCluster(ac.AID, ac.CID, ac.CreatedBy)
}
func (ac *AppCluster) Delete() error {
	return models.DeleteAppCluster(ac.AID, ac.CID)
}

func (au *AppUser) Add() error {
	return models.AddAppUser(au.AID, au.UID, au.CreatedBy)
}

func (au *AppUser) Get(aid uint) ([]*models.User, error) {
	return models.GetAppUsers(aid)
}
