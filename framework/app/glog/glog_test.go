/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 20:04:42
 */

package glog

import (
	"framework/app/build"
	"framework/app/glog/log"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGLog(t *testing.T) {
	var err error
	// 初始化 test 日志
	var testFileLog *log.Log
	if testFileLog, err = log.New(log.Config{
		Directory: "logs",
		Level:     logrus.DebugLevel,
	}, "test", "test"); err != nil {
		t.Fatal(err.Error())
	}

	gLog := GLog{
		Layout: DefaultLayout{},
		Appenders: []Appender{
			&FileAppender{
				Debug: testFileLog,
				Info:  testFileLog,
				Warn:  testFileLog,
				Error: testFileLog,
			},
		},
	}

	if build.IsDev() {
		gLog.Appenders = append(gLog.Appenders, &ConsoleAppender{})
	}

	gLog.Debug("hello world")
	gLog.Warn("hello world")
	gLog.Info("hello world")
	gLog.Error("hello world")
}
