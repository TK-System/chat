package env

import (
	"github.com/BurntSushi/toml"
	"golang.org/x/xerrors"
)

type Settings struct{
	ApiPort string 
	HtmlPageDirPath string
}


func CreateSettings(path string)(*Settings,error){
	//f, err := os.Open("example.toml")
    //if err != nil {
    //    return nil,xerrors.Errorf("cannot read setting file")
	//}
	config := &Settings{}
	_,err := toml.DecodeFile(path,config)
	if err != nil {
        return nil,xerrors.Errorf("cannot read setting file")
	}
	return config,nil
}
