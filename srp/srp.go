package srp

import (
	"fmt"
)

func Run() {
	fmt.Println("Run srp (single responsibility principle)")

	message := &Message{}
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
