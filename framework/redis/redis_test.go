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

	rediz.Set("hello", "world", time.Duration(60))

	value, ttl, ok := rediz.GetString("hello")
	t.Log(*value, ttl, ok)

	value, ttl, ok = rediz.GetString("hello-not-exist")
	t.Log(value, ttl, ok)

	if ok := rediz.Exists("hello"); !ok {
		t.Fatalf("hello exist")
	}

	if ttl, ok := rediz.getTtl("hello"); !ok {
		t.Fatalf("hello not exist")
	} else {
		t.Log(ttl)
	}

	if _, ok := rediz.getTtl("hello-not-exist"); ok {
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
	t.Log(*value, ok)
}
