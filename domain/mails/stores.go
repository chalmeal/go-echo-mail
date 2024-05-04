package mails

import "go-echo-mail/common"

var (
	db = common.DbConnect()
)

func getAllMails() *[]Mails {
	m := []Mails{}
	db.Find(&m)
	return &m
}

func GetMailByAddress(a []string) *[]Mails {
	m := []Mails{}
	db.Where("mail_address in (?)", a).Find(&m)
	return &m
}

func RegMail(p *RegMailParam) error {
	result := db.Create(&Mails{
		MailAddress: p.MailAddress,
		Name:        p.Name,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
