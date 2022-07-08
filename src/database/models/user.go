package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Identifier string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}

func GetUserByIdentifier(db *gorm.DB, identifier string) (*User, error) {
	var user User

	err := db.Debug().Model(User{}).Where("identifier = ?", identifier).Take(&user)
	
	if err != nil {
		return &User{}, errors.New("error getting user")
	}

	return &user, nil
}