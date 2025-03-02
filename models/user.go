package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username          string     `json:"username" gorm:"unique;not null"`
	Email             string     `json:"email" gorm:"unique;not null"`
	Password          string     `json:"-" gorm:"not null"`
	Name              string     `json:"name" gorm:"not null"`
	Phone             *string    `json:"phone" gorm:"unique;default:null"`
	IsPremium         bool       `json:"is_premium" gorm:"default:false"`
	IsActive          bool       `json:"is_active" gorm:"default:true"`
	Status            string     `json:"status" gorm:"default:'active'"`
	EmailVerified     bool       `json:"email_verified" gorm:"default:false"`
	PhoneVerified     bool       `json:"phone_verified" gorm:"default:false"`
	VerificationToken string     `json:"verification_token,omitempty" gorm:"type:varchar(100)"`
	Role              string     `json:"role" gorm:"default:'user'"`
	ResetToken        string     `json:"-" gorm:"type:varchar(100)"`
	ResetTokenExpires *time.Time `json:"-" gorm:"default:null"`
}

// HashPassword hashes a plain password
func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// ComparePassword checks the provided password
func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
