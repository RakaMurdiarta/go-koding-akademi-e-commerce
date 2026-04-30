package models

import (
	"time"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/common"
	"gorm.io/gorm"
)

type User struct {
	ID         uint
	Email      string `gorm:"size:100;uniqueIndex;not null"`
	Password   *string
	FullName   string
	Role       common.UserRole `gorm:"type:user_role_enum;default:'buyer'"`
	Provider   common.Provider `gorm:"type:provider_enum;default:'local'"`
	ProviderID *string
	AvatarURL  *string
	Session    *string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Addresses []UserAddress
	Shop      *Shop
	Carts     []Cart
}

type UserAddress struct {
	ID           uint
	UserID       uint
	ReceiverName string
	Phone        string
	Address      string
	City         string
	Province     string
	PostalCode   string
	IsDefault    bool
	CreatedAt    time.Time
	User         User
}
