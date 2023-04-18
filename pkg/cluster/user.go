package cluster

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	Id       string `json:"id" column:"id"`             // 用户ID
	Username string `json:"username" column:"username"` // 用户名
	Password string `json:"password" column:"password"` // 密码
	Enable   bool   `json:"enable" column:"enable"`     // 是否启用
}

func userDatabase() *gorm.DB {
	return GetDBConnection("k8s")
}

func (User) TableName() string {
	return "user"
}

func GetUserById(id string) (User, error) {
	var user User
	db := userDatabase()
	db.Where("id = ?", id).First(&user)
	return user, db.Error
}

func GetUserList() ([]User, error) {
	var users []User
	db := userDatabase()
	db.Find(&users)
	return users, db.Error
}

func AddUser(user User) error {
	db := userDatabase()
	exist, err := CheckUserExist(user.Username)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("用户已存在")
	}
	db.Create(&user)
	return db.Error
}

func UpdateUser(user User) error {
	db := userDatabase()
	db.Save(&user)
	return db.Error
}

func DeleteUser(id string) error {
	db := userDatabase()
	db.Where("id = ?", id).Delete(&User{})
	return db.Error
}

func GetUserByUsername(username string) (User, error) {
	var user User
	db := userDatabase()
	db.Where("username = ?", username).First(&user)
	return user, db.Error
}

func CheckUser(username string, password string) (bool, error) {
	user, err := GetUserByUsername(username)

	if err != nil {
		return false, err
	}

	if user.Username == "" {
		return false, nil
	}
	if user.Password == password {
		return true, nil
	}
	return false, nil
}

func CheckUserExist(username string) (bool, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return false, err
	}
	return user.Username != "", nil
}

func CheckUserEnable(username string) (bool, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return false, err
	}
	return user.Enable, nil
}

// 初始化一个用户admin
func InitUser() {
	username := "admin"
	password := "admin"

	exist, err := CheckUserExist(username)
	if err != nil {
		panic(err)
	}

	if !exist {
		user := User{
			Username: username,
			Password: password,
			Enable:   true,
		}
		err = AddUser(user)
		if err != nil {
			panic(err)
		}
	}

}
