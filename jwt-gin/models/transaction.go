package models

import (
	// "html"
	// "strings"
	"errors"
	// "fmt"
	// "time"

	"github.com/jinzhu/gorm"
	// "golang.org/x/crypto/bcrypt"
)

type Transaction struct {
	gorm.Model
	WalletID uint
	Wallet Wallet `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Debit float64 `gorm:"size:256" json:"debit"`
	Credit float64 `gorm:"size:256" json:"credit"`
	Remarks string `gorm:"size:256" json:"remarks"`
}

// type TransactionDetails struct {
// 	CreatedAt  time.Time
// 	Debit float64
// 	Credit float64
// 	Remarks string
//   }
  

func (t *Transaction) CreateTransaction() (*Transaction, error) {

	var err error
	
	err = DB.Create(&t).Error
	if err != nil {
		return &Transaction{}, err
	}
	return t, nil
}

func GetTransactions(wallet uint) ([]Transaction, error) {
	

	var t []Transaction
	if err := DB.Where("wallet_id = ?", wallet).Find(&t).Error; err != nil {
		return t,errors.New("No transaction has been made!")
	}

	return t, nil
}