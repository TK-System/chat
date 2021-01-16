package models

type UserGroup struct {
	Userid  int `xorm:"not null pk INT"`
	Groupid int `xorm:"not null pk index INT"`
}
