package config

import (
	"gopkg.in/ini.v1"
	"memorandum/model"
	"strings"
)

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
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		println("配置文件路径错误")
	}
	LoadServer(file)
	LoadMysql(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	println(path)
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
