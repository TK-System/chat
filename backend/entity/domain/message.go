package domain

import "time"

type Message interface {
	ID() int
	Text() string
	TimeStamp() time.Time
}
