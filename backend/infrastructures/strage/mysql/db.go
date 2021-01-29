package mysql

import (
	"chat/backend/infrastructures/strage/mysql/models"
	"fmt"

	"xorm.io/core"
	"xorm.io/xorm"

	_ "github.com/go-sql-driver/mysql"
)




type DB struct{
	*xorm.Engine
}
//e("mysql", "root:root@tcp([127.0.0.1]:3306)/hoge?charset=utf8mb4&parseTime=true")
func NewDB(user,pass,host string,port interface{},dbname string)(*DB,error) {
	url := fmt.Sprintf("%s:%s@tcp([%s]:%v)/%s",user,pass,host,port,dbname)
	engine, err := xorm.NewEngine("mysql", url) 
	if err != nil {
		return nil,err
	}
	engine.ShowSQL(true)
	engine.SetMapper(core.GonicMapper{})
	return  &DB{
		Engine: engine,
	}	,nil

} 
 
func (db *DB)Close(){
	db.Engine.Close()
}
func (db *DB)NewUserTransfer()*models.UserTransfer{
	return &models.UserTransfer{
		Engine:db.Engine,
	}
}

