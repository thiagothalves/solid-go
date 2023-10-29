package ocp

import (
	"fmt"
)

func Run() {
	fmt.Println("Run ocp (open closed principle)")

	messageTemplateCompetition := &MessageTemplateCompetition{}
	message := &Message{
		MessageTemplateCompetition: messageTemplateCompetition,
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
