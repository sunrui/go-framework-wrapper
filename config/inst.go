package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Mysql     Mysql     `json:"mysql"`     // mysql
	Redis     Redis     `json:"redis"`     // redis
	RateLimit RateLimit `json:"rateLimit"` // rateLimit
	Log       Log       `json:"log"`       // log
	Jwt       Jwt       `json:"jwt"`       // jwt
	Request   Request   `json:"request"`   // request
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
