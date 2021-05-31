package notification

import (
	"fmt"
	"gb-go-architecture/lesson-7/shop_new/server/models"
	"net/smtp"
)

type SmtpClient struct {
	addr string
	from string
	auth smtp.Auth
}

const emailTemplate = `To: "%s" <%s>
From: "MyShop" <%s>
Subject: New order: %d

Hello!

You have created the new order %d.
Items: %v

Best Regards, MyShop
`

func (s *SmtpClient) SendOrderCreated(order *models.Order) error {
	text := fmt.Sprintf(
		emailTemplate, order.CustomerName, order.CustomerEmail, s.from, order.ID, order.ID, order.ItemIDs,
	)

	fmt.Printf("email to %s about order %d\n", order.CustomerEmail, order.ID)
	return smtp.SendMail(
		s.addr, s.auth, s.from, []string{order.CustomerEmail}, []byte(text),
	)
}

func NewSMTPClient(from, password, host, port string) (*SmtpClient, error) {
	return &SmtpClient{
		addr: host + ":" + port,
		from: from,
		auth: smtp.PlainAuth("", from, password, host),
	}, nil
}
