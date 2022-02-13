package models

import (
	"errors"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model

	Name        string `json:"name" gorm:"unique; not null;comment:'产品名称'"`
	Parent      uint   `json:"parent" gorm:"comment:'上级产品'"`
	Description string `json:"description"`
	Status      uint   `json:"status"`
	CreatedBy   uint   `json:"created_by"`
	UpdatedBy   uint   `json:"updated_by"`
}

type ProductUser struct {
	gorm.Model

	PID uint `json:"p_id" gorm:"comment:'产品负责人-产品id'"`
	UID uint `json:"u_id" gorm:"comment:'产品负责人-用户id'"`

	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
}

func GetProduct(id uint) (*Product, error) {
	var product Product
	err := db.Where("id = ? and deleted_at is NULL", id).First(&product).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &product, nil
}

func GetProducts() ([]*Product, error) {
	var products []*Product
	err := db.Where("deleted_at is NULL").Find(&products).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return products, nil
}

func EditProduct(id uint, data interface{}) error {
	var product Product
	if err := db.Model(&product).Where("id = ? AND deleted_at is ?", id, "NULL").Update(data).Error; err != nil {
		return err
	}

	return nil
}

func AddProduct(data map[string]interface{}) (id uint, err error) {
	product := Product{
		Name:        data["name"].(string),
		Description: data["description"].(string),
		Parent:      data["parent"].(uint),
		Status:      data["status"].(uint),
		CreatedBy:   data["created_by"].(uint),
	}
	if err := db.Create(&product).Error; err != nil {
		return 0, err
	}

	return product.ID, nil
}
func AddProductUser(pid, uid, createdBy uint) error {
	if exists, _ := ExistProductByID(pid); exists == false {
		return errors.New("product not found " + strconv.Itoa(int(pid)))
	}
	if exists, _ := ExistUserByID(uid); exists == false {
		return errors.New("user not found" + strconv.Itoa(int(uid)))
	}

	productUser := ProductUser{
		PID:       pid,
		UID:       uid,
		CreatedBy: createdBy,
	}
	if err := db.Create(&productUser).Error; err != nil {
		return err
	}

	return nil
}

func GetProductUsers(pid uint) ([]*User, error) {
	var users []*User
	err := db.Debug().Table("user").Where("id = any(select uid from product_user where p_id = ?)", pid).Scan(&users).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}
func DeleteProductUser(pid, uid uint) error {
	if err := db.Where("pid = ? AND uid = ? AND deleted_at is NULL").Delete(&ProductUser{}).Error; err != nil {
		return err
	}
	return nil
}
func AddProductWithUID(pid, uid uint) error {

	var user User
	if err := db.Where("id = ? AND deleted_at is NULL", uid).First(&user).Error; err != nil {
		return err
	}
	productUser := ProductUser{
		UID: uid,
		PID: pid,
	}
	if err := db.Create(&productUser).Error; err != nil {
		log.Fatalf("create product %d with user %d failed: %v", pid, uid, err)
	}
	return nil
}

func DeleteProduct(id uint) error {
	if err := db.Where("id = ?", id).Delete(&Product{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllProduct() error {
	if err := db.Unscoped().Where("deleted_at is not NULL").Delete(&Product{}).Error; err != nil {
		return err
	}

	return nil
}

func ExistProductByID(id uint) (bool, error) {
	var product Product
	err := db.Select("id").Where("id = ? and deleted_at is NULL", id).First(&product).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if product.ID > 0 {
		return true, nil
	}

	return false, nil
}
