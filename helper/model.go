package helper

import coreHelper "github.com/0LuigiCode0/msa-core/helper"

//Config модель конфига
type Config struct {
	Host      string                    `json:"host"`
	Port      int32                     `json:"port"`
	DBS       map[string]*DbConfig      `json:"dbs"`
	Handlers  map[string]*HandlerConfig `json:"handlers"`
	Admins    []*Admin                  `json:"admins"`
	Observers []*HandlerConfig          `json:"observers"`
	Services  []*HandlerConfig          `json:"services"`
}

type DbConfig struct {
	Type     DBType `json:"type"`
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type HandlerConfig struct {
	Type     HandlerType           `json:"type"`
	Host     string                `json:"host"`
	Port     int32                 `json:"port"`
	Key      string                `json:"key"`
	Group    coreHelper.GroupsType `json:"group"`
	User     string                `json:"user"`
	Password string                `json:"password"`
	IsTSL    bool                  `json:"is_tsl"`
}

type Admin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type DBType string

const (
	Postgres DBType = "postgres"
	Mongodb  DBType = "mongodb"
)

type HandlerType string

const (
	TCP  HandlerType = "tcp"
	MQTT HandlerType = "mqtt"
	WS   HandlerType = "ws"
	GRPC HandlerType = "grpc"
)

//Основные константы
const (
	ConfigDir  = "./source/configs/"
	ConfigFile = "configServer.json"
	Secret     = "qfdQjmVLiW"
)

//Названеия коллекции

type Collection string

const (
	CollUsers Collection = "users"
)

//Ключи контекста

type CtxKey int

const (
	CtxKeyValue CtxKey = iota
)

// Названия функций

const (
	SelectByID    = "select/id"
	SelectByLogin = "select/login"
)
