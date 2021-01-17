package models

import (
	"time"
)

type Message struct {
	Id       int       `xorm:"not null pk autoincr INT"`
	Userid   int       `xorm:"not null index INT"`
	Message  string    `xorm:"TEXT"`
	Createat time.Time `xorm:"TIMESTAMP"`
	Roomid   int       `xorm:"not null index INT"`
}
