package models

type Friend struct {
	Me     int `xorm:"not null pk INT"`
	You    int `xorm:"not null pk index INT"`
	Roomid int `xorm:"not null index INT"`
	Love   int `xorm:"not null default 1 TINYINT(1)"`
}
