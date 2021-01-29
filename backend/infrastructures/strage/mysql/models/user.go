package models

import (
	"chat/backend/interfaces/gateway/database"

	"golang.org/x/xerrors"
	"xorm.io/xorm"
)



type User struct {
	Id       int    `xorm:"not null pk autoincr INT"`
	Username string `xorm:"not null VARCHAR(20)"`
	Mail     string `xorm:"not null VARCHAR(50)"`
	Pass     string `xorm:"not null VARCHAR(20)"`
	Sex      int    `xorm:"INT"`
	Lang     int    `xorm:"INT"`
}


type UserTransfer struct{
	*xorm.Engine
} 


type UserRaw struct{
	user *User
}

func (m *UserRaw)ID()int{
	return m.user.Id
}
func (m *UserRaw)Name ()string{
return m.user.Username 
}
func (m *UserRaw)Mail     ()string{
return m.user.Mail     
}
func (m *UserRaw)Pass     ()string{
return m.user.Pass     
}
func (m *UserRaw)Sex      ()int{
return m.user.Sex      
}   
func (m *UserRaw)Lang     ()int{
return m.user.Lang     
}   

func (t *UserTransfer)GetUser(id int)(database.User,error){
	u := &User{}
	has,err :=t.Engine.Where("id = ?",id).Get(u)
	if err !=nil{
		return nil,err
	}

	if !has{
		return nil,xerrors.New("no user")
	}

	return &UserRaw{user:u},nil
}

func (t *UserTransfer)GetUsers(ids []int)([]database.User,error){
	u := make([]User,0,len(ids))
	err :=t.Engine.Table("user").
	In("id",[]interface{}{ids[0]}).
	Find(&u)
	if err !=nil{
		return nil,err
	}

	retUser := make([]database.User,len(u))
	for i,v := range u{
		retUser[i]=&UserRaw{
			user: &v,
		}
	}

	return retUser,nil
}

func (t *UserTransfer)CreateUser(ids []int)error{
	u := make([]*User,0,len(ids))
	err :=t.Engine.
	In("id",[]interface{}{ids}).
	Find(&u)
	if err !=nil{
		return err
	}

	return nil
}
