/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-06 17:53:55
 */

package redis

import (
	"log"
	"testing"
	"time"
)

func TestRedis_Set(t *testing.T) {
	Inst.Set("hello", "world", 10*time.Second)
}

func TestRedis_GetString(t *testing.T) {
	Inst.Set("hello", "world", 10*time.Second)

	value := Inst.GetString("hello")
	if value == nil {
		log.Fatalf("get key failed")
	} else {
		log.Println(*value)
	}
}

func TestRedis_GetJson(t *testing.T) {
	Inst.Set("hello", "{\"hello\":\"world\"}", 10*time.Second)

	var dst any
	r := Inst.GetJson("hello", &dst)
	if !r {
		log.Fatalf("get key failed")
	} else {
		log.Println(dst)
	}
}

func TestRedis_Exists(t *testing.T) {
	Inst.Set("hello", "world", 10*time.Second)

	r := Inst.Exists("hello")
	if !r {
		log.Fatalf("get key falied")
	}

	r = Inst.Exists("hello-not-exist")
	if r {
		log.Fatalf("key not exist")
	}
}

func TestRedis_Del(t *testing.T) {
	Inst.Set("hello", "world", 10*time.Second)

	r := Inst.Exists("hello")
	if !r {
		log.Fatalf("get key falied")
	}
	Inst.Del("hello")

	r = Inst.Exists("hello")
	if r {
		log.Fatalf("key not exist")
	}
}

func TestRedis_SetHash(t *testing.T) {
	Inst.SetHash("hello", "key", "{\"hello\":\"world\"}")
}

func TestRedis_GetHash(t *testing.T) {
	Inst.SetHash("hello", "key", "{\"hello\":\"world\"}")
	value := Inst.GetHash("hello", "key")
	if value == nil {
		log.Fatalf("get key failed")
	} else {
		log.Println(*value)
	}
}

func TestRedis_GetHashJson(t *testing.T) {
	Inst.SetHash("hello", "key", "{\"hello\":\"world\"}")
	var dst any
	r := Inst.GetHashJson("hello", "key", &dst)
	if !r {
		log.Fatalf("get key failed")
	} else {
		log.Println(dst)
	}
}
