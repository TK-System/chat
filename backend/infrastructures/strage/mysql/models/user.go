package models



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


func (t *UserTransfer)GetUser(id int)(*User,error){
	u := &User{}
	has,err :=t.Engine.Where("id = ?",id).Get(u)
	if err !=nil{
		return nil,err
	}

	if !has{
		return nil,xerrors.New("no user")
	}

	return u,nil
}

func (t *UserTransfer)GetUsers(ids []int)([]*User,error){
	u := make([]*User{},0,len(id))
	err :=t.Engine.
	In("id",ids...).
	Find(&u)
	if err !=nil{
		return nil,err
	}

	return u,nil
}

func (t *UserTransfer)CreateUser(ids []int)(int,error){
	u := make([]*User{},0,len(id))
	err :=t.Engine.
	In("id",ids...).
	Find(&u)
	if err !=nil{
		return nil,err
	}

	return u,nil
}
