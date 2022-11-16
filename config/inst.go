package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Mysql     mysql     `json:"mysql"`     // Mysql
	Redis     redis     `json:"redis"`     // Redis
	RateLimit rateLimit `json:"rateLimit"` // RateLimit
	Log       log       `json:"log"`       // Log
	Request   request   `json:"request"`   // Request
	Token     token     `json:"token"`     // Token
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
