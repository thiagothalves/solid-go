package dip

import "fmt"

type IMessageSocmed interface {
	Create() *MessageSocmed
}

type MessageSocmed struct {
	Body            string
	MessageTemplate IMessageTemplate
}

type IMessageEmail interface {
	Create() *MessageEmail
}

type MessageEmail struct {
	Subject         string
	Body            string
	MessageTemplate IMessageTemplate
}

func (m *MessageSocmed) Create() *MessageSocmed {
	template := m.MessageTemplate.Create()
	m = &MessageSocmed{
		Body: "Hey guys! " + template,
	}

	fmt.Println("Create Message")
	fmt.Printf("%+v\n", m)
	return m
}

func (m *MessageEmail) Create() *MessageEmail {
	template := m.MessageTemplate.CreateHtml()

	m = &MessageEmail{
		Subject: "Info",
		Body:    "Hey guys! " + template,
	}

	fmt.Println("Create Message")
	fmt.Printf("%+v\n", m)
	return m
}
