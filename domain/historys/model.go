package historys

import "gorm.io/gorm"

type Historys struct {
	MailId int    `gorm:"not null" json:"mail_id"`
	Title  string `gorm:"size:50;not null" json:"title"`
	Detail string `gorm:"size:1000;not null" json:"detail"`
	// timezone is UTC
	gorm.Model
}

type RegHistoryParam struct {
	MailIds []int  `json:"mail_id"`
	Title   string `json:"title"`
	Detail  string `json:"detail"`
}

type NotifyMailParam struct {
	MailAddress []string `json:"mail_address" validate:"required"`
	Title       string   `json:"title" validate:"required"`
	Detail      string   `json:"detail" validate:"required"`
}
