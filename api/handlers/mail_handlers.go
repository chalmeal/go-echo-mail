package handlers

import (
	"errors"
	"go-echo-mail/api/support"
	"go-echo-mail/common"
	"go-echo-mail/config"
	"go-echo-mail/domain/historys"
	"go-echo-mail/domain/mails"
	"net/http"

	"github.com/labstack/echo/v4"
)

// get all history
func GetAllHistorys(c echo.Context) error {
	r := historys.GetAllHistorys()
	return common.Res(c, http.StatusOK, r)
}

// register mail
func RegisterMail(c echo.Context) error {
	p := new(mails.RegMailParam)
	c.Bind(p)
	if err := c.Validate(p); err != nil {
		return common.Res(c, http.StatusBadRequest, err)
	}

	if err := mails.RegMail(p); err != nil {
		return common.Res(c, http.StatusBadRequest, err)
	}

	r := "Register mail success"

	return common.Res(c, http.StatusOK, r)
}

// notify mail
// and create a history of notify sent
func NotifyMail(c echo.Context) error {
	p := new(historys.NotifyMailParam)
	c.Bind(p)
	if err := c.Validate(p); err != nil {
		return common.Res(c, http.StatusBadRequest, err)
	}

	mail := mails.GetMailByAddress(p.MailAddress)

	var ids []int
	var address []string
	for _, a := range p.MailAddress {
		for _, m := range *mail {
			if m.MailAddress == a {
				ids = append(ids, int(m.ID))
				address = append(address, m.MailAddress)
			}
		}
	}
	if len(address) != len(p.MailAddress) {
		err := errors.New("contains a non-existent email address")
		return common.Res(c, http.StatusBadRequest, err.Error())
	}

	sp := &support.SendMailParam{
		From:    config.MAIL_FROM,
		To:      address,
		Subject: p.Title,
		Body:    p.Detail,
	}
	if err := support.SendMail(sp); err != nil {
		return common.Res(c, http.StatusInternalServerError, err)
	}

	h := &historys.RegHistoryParam{
		MailIds: ids,
		Detail:  p.Detail,
	}
	if err := historys.RegHistory(h); err != nil {
		return common.Res(c, http.StatusInternalServerError, err)
	}

	return common.Res(c, http.StatusOK, "Send mail success")
}
