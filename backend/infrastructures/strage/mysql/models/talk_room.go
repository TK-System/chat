package models

type TalkRoom struct {
	Id       int    `xorm:"not null pk autoincr INT"`
	Roomname string `xorm:"VARCHAR(50)"`
}
