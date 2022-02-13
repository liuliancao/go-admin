package userservice

import (
	"go-admin/models"
)

type Auth struct {
	Username string
	Password string
}
type User struct {
	ID        uint   `json:"id"`
	DID       uint   `json:"d_id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Mail      string `json:"mail"`
	Token     string `json:"token"`
	Status    uint   `json:"status"`
	CreatedBy uint   `json:"created_by"`
	UpdatedBy uint   `json:"updated_by"`
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}
func (u *User) Add() error {
	user := map[string]interface{}{
		"d_id":       u.DID,
		"username":   u.Username,
		"nickname":   u.Nickname,
		"password":   u.Password,
		"gender":     u.Gender,
		"phone":      u.Phone,
		"mail":       u.Mail,
		"token":      u.Token,
		"status":     u.Status,
		"created_by": u.CreatedBy,
	}

	if err := models.AddUser(user); err != nil {
		return err
	}

	return nil
}
func (u *User) Edit() error {
	user := &models.User{
		DID:       u.DID,
		Username:  u.Username,
		Nickname:  u.Nickname,
		Password:  u.Password,
		Gender:    u.Gender,
		Phone:     u.Phone,
		Mail:      u.Mail,
		Token:     u.Token,
		Status:    u.Status,
		UpdatedBy: u.UpdatedBy,
	}
	return models.EditUser(u.ID, u.DID, user)
}

func (u *User) Get() (*models.User, error) {
	user, err := models.GetUser(u.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (u *User) GetIDByMD5(md5 string) (uint, error) {
	userID, err := models.GetIDByMD5(md5)
	if err != nil {
		return 0, err
	}

	return uint(userID), nil
}

func (u *User) GetAll() ([]*models.VUser, error) {
	var users []*models.VUser
	users, err := models.GetVerboseUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) ExistByID() (bool, error) {
	return models.ExistUserByID(u.ID)
}

func (u *User) Delete() error {
	return models.DeleteUser(u.ID)
}
