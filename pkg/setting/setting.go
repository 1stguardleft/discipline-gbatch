package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type Runner struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var RunnerSetting = &Runner{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
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

	mapTo("runner", RunnerSetting)
	mapTo("database", DatabaseSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
