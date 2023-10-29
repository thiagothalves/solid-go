package isp

import (
	"fmt"
)

func Run() {
	fmt.Println("Run isp (interface segregation principle)")

	messageTemplateWebinar := &MessageTemplateWebinar{}
	messageSocmed := &MessageSocmed{
		MessageTemplate: messageTemplateWebinar,
	}
	messageEmail := &MessageEmail{
		MessageTemplate: messageTemplateWebinar,
	}

	messageSocmedPayload := messageSocmed.Create()
	messageEmailPayload := messageEmail.Create()
	fmt.Println()

	user := &User{}
	senderSocmed := &SenderSocmed{
		Sender:   user.GetSender(),
		Receiver: user.GetReceiver(),
		Message:  messageSocmedPayload.Body,
	}
	senderEmail := &SenderEmail{
		Sender:   user.GetSender(),
		Receiver: user.GetReceiver(),
		Subject:  messageEmailPayload.Subject,
		Body:     messageEmailPayload.Body,
	}

	senderSocmed.SendWhatsapp()
	fmt.Println()
	senderSocmed.SendTelegram()
	fmt.Println()
	senderEmail.SendEmail()
}
