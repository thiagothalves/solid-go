package lsp

import "fmt"

type IMessage interface {
	Create() string
}

type Message struct {
	Body            string
	MessageTemplate IMessageTemplate
}

func (m *Message) Create() string {
	template := m.MessageTemplate.Create()
	m = &Message{
		Body: "Hey guys! " + template,
	}

	res := m.Body

	fmt.Println("Create Message")
	fmt.Printf("%+v\n", m)
	return res
}
