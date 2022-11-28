/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-28 15:31:21
 */

package redis

import (
	"framework/config"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	rediz := New(config.Redis{
		Host:     "localhost",
		Port:     6379,
		Password: "",
		Database: 0,
	})

	rediz.Set("hello", []byte("world"), time.Duration(60))

	value, ttl, ok := rediz.Get("hello")
	t.Log(string(value), ttl, ok)

	value, ttl, ok = rediz.Get("hello-not-exist")
	t.Log(string(value), ttl, ok)

	if ok := rediz.Exists("hello"); !ok {
		t.Fatalf("hello exist")
	}

	if ttl, ok := rediz.GetTtl("hello"); !ok {
		t.Fatalf("hello not exist")
	} else {
		t.Log(ttl)
	}

	if _, ok := rediz.GetTtl("hello-not-exist"); ok {
		t.Fatalf("hello not exist")
	}

	if ok := rediz.Del("hello"); !ok {
		t.Fatalf("delete failed")
	} else {
		if ok = rediz.Del("hello"); ok {
			t.Fatalf("cannot delete")
		}
	}

	rediz.SetHash("hash", "hello", "world")
	value, ok = rediz.GetHash("hash", "hello")
	t.Log(string(value), ok)
}
