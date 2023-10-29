package dip

import (
	"fmt"
	"time"
)

type IMessageTemplate interface {
	Create() string
	CreateHtml() string
}

type MessageTemplate struct {
	Name      string
	StartedAt time.Time
	EndedAt   time.Time
}

type MessageTemplateWebinar struct {
	MessageTemplate
	Title  string
	Author string
}

type MessageTemplateCompetition struct {
	MessageTemplate
	Name  string
	Level string
}

func (m *MessageTemplateWebinar) New() *MessageTemplateWebinar {
	m = &MessageTemplateWebinar{
		MessageTemplate: MessageTemplate{
			Name:      "WEBINAR",
			StartedAt: time.Now(),
			EndedAt:   time.Now().Add(1 * time.Hour),
		},
		Title:  "Software engineer in Teresópolis / RJ",
		Author: "Thiago",
	}
	return m
}

func (m *MessageTemplateWebinar) Create() string {
	v := m.New()
	res := fmt.Sprintf("Webinar %s, ministered by %s will start %s until %s", v.Title, v.Author, v.StartedAt, v.EndedAt)
	return res
}

func (m *MessageTemplateWebinar) CreateHtml() string {
	v := m.New()
	res := fmt.Sprintf("[Html] Webinar %s, ministered by %s will start %s until %s", v.Title, v.Author, v.StartedAt, v.EndedAt)
	return res
}

func (m *MessageTemplateCompetition) New() *MessageTemplateCompetition {
	m = &MessageTemplateCompetition{
		MessageTemplate: MessageTemplate{
			Name:      "COMPETITION",
			StartedAt: time.Now(),
			EndedAt:   time.Now().Add(1 * time.Hour),
		},
		Name:  "Competitive Programming",
		Level: "Hard",
	}
	return m
}

func (m *MessageTemplateCompetition) Create() string {
	res := fmt.Sprintf("Competition %s, level %s will start %s until %s", m.Name, m.Level, m.StartedAt, m.EndedAt)
	return res
}

func (m *MessageTemplateCompetition) CreateHtml() string {
	res := fmt.Sprintf("[Html] Competition %s, level %s will start %s until %s", m.Name, m.Level, m.StartedAt, m.EndedAt)
	return res
}
