package uitls

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	f, err := ini.Load("../config/config.ini")
	if err != nil {
		log.Printf("配置文件读取错误:%s",err)
	}



}

// func LoadServer(file *ini.File) {
// 	AppMode = file.Section("server").Key("AppMode").MustString("debug")
// 	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
// }

// func LoadData(file *ini.File) {
// 	Db = file.Section("database").Key("Db").MustString("debug")
// 	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
// 	DbPort = file.Section("database").Key("DbPort").MustString("3306")
// 	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
// 	DbPassWord = file.Section("database").Key("DbPassWord").MustString("admin123")
// 	DbName = file.Section("database").Key("DbName").MustString("ginblog")
// }