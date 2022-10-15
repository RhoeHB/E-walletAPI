package models

import (
	// "html"
	// "strings"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	// "golang.org/x/crypto/bcrypt"
)

type Wallet struct {
	gorm.Model
	UserID uint
	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Balance float64 `gorm:"size:256;not null" json:"balance"`
}


func ReturnBalance(uid uint) (float64,error) {

	var w Wallet

	if err := DB.Where("user_id = ?", uid).First(&w).Error; err != nil {
		return w.Balance,errors.New("User's wallet not found!")
	}

	return w.Balance,nil

}

func UpdateBalance(uid uint, debit, credit float64) (Wallet,error){

	var w Wallet

	if err := DB.Where("user_id = ?", uid).First(&w).Error; err != nil {
		return w,errors.New("User's wallet not found!")
	}
	
	balance := w.Balance + debit - credit

	if err := DB.Model(&w).Where("user_id = ?", uid).Update("Balance", balance).Error; err != nil {
		return w,errors.New("Fail to update balance!")
	}
	
	return w,nil

}