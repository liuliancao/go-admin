package models

import (
	"go-admin/pkg/util"
	"log"

	"github.com/jinzhu/gorm"
)

//User used for user login and jwt etc.
type User struct {
	gorm.Model

	DID       uint   `json:"d_id" gorm:"comment:'部门id'"`
	Username  string `json:"username" gorm:"unique;not null;comment:'用户名'"`
	Nickname  string `json:"nickname" comment:'页面展示的用户名'`
	TID       string `json:"tid" gorm:"default:3;not null;"` // == User.Type with forcement
	Password  string `json:"password" gorm:"comment:'密码'"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Mail      string `json:"mail"`
	Token     string `json:"token" gorm:"comment:'md5 用户名'"`
	Status    uint   `json:"status"`
	CreatedBy uint   `json:"created_by"`
	UpdatedBy uint   `json:"updated_by"`
}

//UserType used for app user(dev, ops ...)
type UserType struct {
	gorm.Model

	Name        string `json:"name" gorm:"comment:'用户类型名称dev,ops,qa，开发运维等等'"`
	Description string `json:"description"`
	CreatedBy   uint   `json:"created_by"`
	UpdatedBy   uint   `json:"updated_by"`
}

type VUser struct {
	ID         uint   `json:"id"`
	Department string `json:"department"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Usertype   string `json:"usertype"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Mail       string `json:"mail"`
	Status     string `json:"status"`
	CreatedBy  string `json:"created_by"`
	UpdatedBy  string `json:"updated_by"`
}

// CheckAuth for auth check with jwt
func CheckAuth(username, password string) (bool, error) {
	var user User
	//err := db.Select("id").Where(User{username: username, password: password}).First(&user).Error
	password_digest := util.EncodeSha256(password)
	err := db.Debug().First(&user, "username  = ? and password = ?", username, password_digest).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

//GetUser get user by id without password and token
func GetUser(id uint) (*User, error) {
	var user User
	err := db.Where("id = ? and deleted_at is NULL", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	user.Password = ""
	user.Token = ""
	return &user, nil
}

//GetUsers get all users, not going to use for security
func GetUsers() ([]*User, error) {
	var users []*User
	err := db.Where("deleted_at is NULL").Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

//GetVerboseUsers
func GetVerboseUsers() ([]*VUser, error) {
	var vusers []*VUser
	newdb := db.New()
	err := newdb.Debug().Table("user").Select("user.id as id, user.username, user.nickname, user.gender, user.phone, user.mail, user.status, department.name as department, user_type.name as usertype, create_user.username as created_by, update_user.username as updated_by").Joins("left join department on department.id = user.d_id").Joins("left join user_type on user_type.id = user.t_id").Joins("left join user as create_user on create_user.id = user.created_by").Joins("left join user as update_user on update_user.id = user.updated_by").Scan(&vusers).Error
	if err != nil {
		log.Println("get verbose user caught error")
		return nil, err
	}
	return vusers, nil
}
func GetUsersByDid(did uint) ([]*User, error) {
	var users []*User
	err := db.Where("d_id = ? and deleted_at is NULL", did).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}
func EditUser(id uint, did uint, u *User) error {
	var user User
	if _, err := ExistDepartmentByID(did); err != nil {
		return err
	}

	u.Password = util.EncodeSha256(u.Password)

	if err := db.Model(&user).Where("id = ? AND deleted_at is NULL", id).Update(*u).Error; err != nil {
		return err
	}

	return nil
}

func AddUser(data map[string]interface{}) error {
	user := User{
		DID:       data["d_id"].(uint),
		Username:  data["username"].(string),
		Nickname:  data["nickname"].(string),
		Password:  util.EncodeSha256(data["password"].(string)),
		Gender:    data["gender"].(string),
		Phone:     data["phone"].(string),
		Mail:      data["mail"].(string),
		Token:     data["token"].(string),
		Status:    data["status"].(uint),
		CreatedBy: data["created_by"].(uint),
	}
	department := Department{}
	if err := db.Where("id = ?", user.DID).First(&department).Error; err != nil {
		return err
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(id uint) error {
	if err := db.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllUser() error {
	if err := db.Unscoped().Where("deleted_at is not NULL").Delete(&User{}).Error; err != nil {
		return err
	}

	return nil
}

func ExistUserByID(id uint) (bool, error) {
	var user User
	err := db.Select("id").Where("id = ? and deleted_at is NULL", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetIDByMD5(md5 string) (uint, error) {
	var user User
	err := db.Select("id").Where("token = ? and deleted_at is NULL", md5).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	log.Println(user)
	if user.ID > 0 {
		return user.ID, nil
	}

	return 0, nil
}
