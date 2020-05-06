package model

import "time"

const DomainModelNameUser DomainModelName = "User"

type User struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Account     string     `json:"account" validate:"required"`
	FirstName   string     `json:"first_name" validate:"required"`
	LastName    string     `json:"last_name" validate:"required"`
	MailAddress string     `json:"mail_address" validate:"required,email"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

type UserList struct {
	Users   []User `json:"users"`
	HasNext bool   `json:"has_next"`
	Cursor  uint   `json:"cursor"`
}
