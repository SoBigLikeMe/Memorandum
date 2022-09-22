package config

import (
	"gopkg.in/ini.v1"
	"memorandum/model"
	"strings"
)

//初始化变量
var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbUser     string
	DbPassWord string
	DbName     string
	DbPort     string
)

func Init() {
	//加载配置文件
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		println("配置文件路径错误")
	}

	LoadServer(file) //加载Server配置
	LoadMysql(file)  //加载mysql配置

	//mysql连接指令
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	//println(path)

	//进行数据库连接
	model.Database(path)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
}
