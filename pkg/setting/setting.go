package setting

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"
)

type App struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	CharSet     string
	TablePrefix string
}

var DatabaseSetting = &Database{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("database", DatabaseSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

func MakeDBString() (string, string) {
	// "root:root123@tcp(127.0.0.1:3306)/gobatch?charset=utf8&parseTime=true"
	dbstring := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&%s",
		DatabaseSetting.User,
		DatabaseSetting.Password,
		DatabaseSetting.Host,
		DatabaseSetting.Name,
		DatabaseSetting.CharSet,
		"parseTime=true")
	return DatabaseSetting.Type, dbstring
}
