package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type config struct {
	Mysql     Mysql     `json:"mysql"`     // mysql
	Redis     Redis     `json:"redis"`     // mysql
	RateLimit RateLimit `json:"rateLimit"` // mysql
}

var Inst *config

func init() {
	type env struct {
		Dev  config `json:"dev"`
		Prod config `json:"prod"`
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
		Inst = &e.Dev
	} else {
		Inst = &e.Prod
	}
}
