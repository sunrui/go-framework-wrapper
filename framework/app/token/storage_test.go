/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-28 16:33:50
 */

package token

import (
	"framework/app/redis"
	"testing"
)

func TestJwtStorage(t *testing.T) {
	tokenStorage := NewJwtStorage([]byte("123456"))

	if tokenString, err := tokenStorage.Set(Payload{UserId: "userId"}, 60); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log("tokenString: " + tokenString)
		token, ttl, err := tokenStorage.Get(tokenString)
		t.Log(token, ttl, err)
	}
}

func TestRedisStorage(t *testing.T) {
	rediz, _ := redis.New(redis.Config{
		Host:     "localhost",
		Port:     6379,
		Password: "honeysenselt",
		Database: 0,
	})

	tokenStorage := NewRedisStorage(rediz, "sid")

	if tokenString, err := tokenStorage.Set(Payload{UserId: "userId"}, 60); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log("tokenString: " + tokenString)
		token, ttl, err := tokenStorage.Get(tokenString)
		t.Log(token, ttl, err)
	}
}
