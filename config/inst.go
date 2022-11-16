package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Mysql     mysql     `json:"mysql"`     // mysql
	Redis     redis     `json:"redis"`     // redis
	RateLimit rateLimit `json:"rateLimit"` // rateLimit
	Log       log       `json:"log"`       // log
	Request   request   `json:"request"`   // request
	Token     token     `json:"token"`     // token
}

var inst *Config

func Inst() *Config {
	if inst != nil {
		return inst
	}

	type env struct {
		Dev  Config `json:"dev"`
		Prod Config `json:"prod"`
	}

	_, file, _, _ := runtime.Caller(0)
	path := filepath.Dir(file)

	var e env
	if stream, err := os.ReadFile(path + "/config.json"); err != nil {
		panic(err.Error())
	} else if err = json.Unmarshal(stream, &e); err != nil {
		panic(err.Error())
	}

	if IsDev() {
		inst = &e.Dev
	} else {
		inst = &e.Prod
	}

	return inst
}
