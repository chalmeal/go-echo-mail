package historys

import "go-echo-mail/common"

var (
	db = common.DbConnect()
)

func GetAllHistorys() *[]Historys {
	c := []Historys{}
	db.Find(&c)
	return &c
}

func RegHistory(p *RegHistoryParam) error {
	for _, id := range p.MailIds {
		result := db.Create(&Historys{
			MailId: id,
			Title:  p.Title,
			Detail: p.Detail,
		})
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
