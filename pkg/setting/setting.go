package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}
type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string
}

type Log struct {
	LogPath  string
	Encoding string // json or console

	Development bool
	ServiceName string
	Level       string
}

var DatabaseSetting = &Database{}
var AppSetting = &App{}
var LogSetting = &Log{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("database", DatabaseSetting)
	mapTo("app", AppSetting)
	mapTo("log", LogSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
