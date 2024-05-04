package domain

import (
	"go-echo-mail/common"
	"go-echo-mail/domain/historys"
	"go-echo-mail/domain/mails"
)

var (
	db = common.DbConnect()
)

func SetUp() {
	db.AutoMigrate(&historys.Historys{})
	db.AutoMigrate(&mails.Mails{})
}
