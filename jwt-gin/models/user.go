package models

import (
	"html"
	"strings"
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"ottoTest/utils/token"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:256;not null;unique" json:"username"`
	Password string `gorm:"size:256;not null;" json:"password"`
	Email string `gorm:"size:256;not null;unique" json:"email"`
}

func DeleteUserByID(uid uint) (error) {

	if err := DB.Where("id = ?", uid).Delete(&User{}).Error; err != nil {
		return errors.New("User not found!")
	}
	
	return nil

}

func ChangeUserPassword(uid uint, password string) (error) {

	var user User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newpass:= string(hashedPassword)

	if err := DB.Model(&user).Where("id = ?", uid).Update("password", newpass).Error; err != nil {
		return errors.New("User not found!")
	}
	
	return nil

}

func GetUserByID(uid uint) (User,error) {

	var u User

	if err := DB.First(&u,uid).Error; err != nil {
		return u,errors.New("User not found!")
	}

	u.PrepareGive()
	
	return u,nil

}

func (u *User) PrepareGive(){
	u.Password = ""
}

func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string,error) {
	
	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token,err := token.GenerateToken(u.ID)

	if err != nil {
		return "",err
	}

	return token,nil
	
}

func (u *User) SaveUser() (*User, error) {

	var err error

	// create user
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	wallet := Wallet{UserID: u.ID, Balance: 0}

	// initialize wallet with 0 balance
	err = DB.Create(&wallet).Error
	if err != nil {
		return &User{}, err
	}
	
	return u, nil
}

func (u *User) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username 
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}