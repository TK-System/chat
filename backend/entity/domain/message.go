package domain

import "time"

type MessageID int

type Message interface {
	ID() MessageID
	Text() string
	TimeStamp() time.Time
}

func NewMessage(id MessageID,text string,timestamp time.Time)Message{
	return &message{
		id:id,
		text:text,
		timestamp:timestamp,		
	}
}

type message struct{
	id MessageID
	text string
	timestamp time.Time
}
func(m *message)ID() MessageID{return m.id}
func(m *message)Text() string{return m.text}
func(m *message)TimeStamp() time.Time{return m.timestamp}
