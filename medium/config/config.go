/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-07 00:15:14
 */

package config

import (
	"encoding/json"
	"os"
)

const file = "config.json"

type env[T any] struct {
	Dev  T `json:"dev"`
	Prod T `json:"prod"`
}

func Load[T any]() T {
	var e env[T]

	if stream, err := os.ReadFile(file); err != nil {
		panic(err.Error())
	} else if err = json.Unmarshal(stream, &e); err != nil {
		panic(err.Error())
	}

	if IsDev() {
		return e.Dev
	} else {
		return e.Prod
	}
}
