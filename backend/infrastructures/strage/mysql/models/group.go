package models

type Group struct {
	Id        int    `xorm:"not null pk autoincr INT"`
	Groupname string `xorm:"not null VARCHAR(50)"`
	Roomid    int    `xorm:"not null index INT"`
}
