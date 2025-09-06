package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	Bio      string `json:"bio" gorm:"type:text"`
	Email    string `json:"email" gorm:"type:varchar(255);unique;not null"`
	Image    string `json:"image" gorm:"type:varchar(255)"`
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// type LoginRequest struct {
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required,min=8"`
// }

// type RegisterRequest struct {
// 	UserName string `json:"username" validate:"required,username"`
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required,min=8"`
// }

// type AuthResponse struct {
// 	Email       string `json:"email"`
// 	AccessToken string `json:"token"`
// 	UserName    string `json:"username"`
// 	Bio         string `json:"bio"`
// 	Image       string `json:"image"`
// }

// type RefreshTokenRequest struct {
// 	RefreshToken string `json:"refreshtoken" validate:"required"`
// }
