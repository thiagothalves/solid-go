package lsp

import (
	"fmt"
)

func Run() {
	fmt.Println("Run lsp (liskov subtitution principle)")

	messageTemplateWebinar := &MessageTemplateWebinar{}
	message := &Message{
		MessageTemplate: messageTemplateWebinar,
	}

	messagePayload := message.Create()
	fmt.Println()

	user := &User{}
	sender := &Sender{
		Sender:   user.GetSender(),
		Receiver: user.GetReceiver(),
		Message:  messagePayload,
	}

	sender.SendWhatsapp()
	fmt.Println()
	sender.SendTelegram()
}
