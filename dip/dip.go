package dip

import (
	"fmt"
)

type Factory struct {
	MessageSocmed IMessageSocmed
	MessageEmail  IMessageEmail
	User          IUser
}

func Run() {
	fmt.Println("Run dip (dependency inversion principle)")

	messageTemplateWebinar := &MessageTemplateWebinar{}
	MessageSocmed := &MessageSocmed{
		MessageTemplate: messageTemplateWebinar,
	}
	MessageEmail := &MessageEmail{
		MessageTemplate: messageTemplateWebinar,
	}
	User := &User{}

	factory := Factory{MessageSocmed, MessageEmail, User}
	Execute(&factory)
}

func Execute(f *Factory) {
	messageSocmedPayload := f.MessageSocmed.Create()
	messageEmailPayload := f.MessageEmail.Create()
	fmt.Println()

	senderSocmed := &SenderSocmed{
		Sender:   f.User.GetSender(),
		Receiver: f.User.GetReceiver(),
		Message:  messageSocmedPayload.Body,
	}
	senderEmail := &SenderEmail{
		Sender:   f.User.GetSender(),
		Receiver: f.User.GetReceiver(),
		Subject:  messageEmailPayload.Subject,
		Body:     messageEmailPayload.Body,
	}

	senderSocmed.SendWhatsapp()
	fmt.Println()
	senderSocmed.SendTelegram()
	fmt.Println()
	senderEmail.SendEmail()
}
