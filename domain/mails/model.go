package mails

import "gorm.io/gorm"

type Mails struct {
	MailAddress string `gorm:"size:100;not null;unique" json:"mail_address"`
	Name        string `gorm:"size:50;not null" json:"name"`
	// timezone is UTC
	gorm.Model
}

type RegMailParam struct {
	MailAddress string `json:"mail_address" validate:"required"`
	Name        string `json:"name" validate:"required"`
}
