
package mysql


type DB struct{
	xorm.Engine
}

func NewDB(user,pass,host string,port interface{},dbname string)(*DB,error) {
	url := fmt.Sprintf("%s:%s@%s:%v/%s",user,pass,host,port,dbname)
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
	DB.Close()
}