/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 20:04:42
 */

package glog

import (
	"framework/app/glog/log"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGLog(t *testing.T) {
	var testFileLog *log.Log
	var err error

	if testFileLog, err = log.New(log.Config{
		Directory: "logs",
		Level:     logrus.DebugLevel,
	}, "test", "test"); err != nil {
		t.Fatal(err.Error())
	}

	gLog := GLog{
		Layout: DefaultLayout{},
		Appenders: []Appender{
			&ConsoleAppender{},
			&FileAppender{
				Debug: testFileLog,
				Info:  testFileLog,
				Warn:  testFileLog,
				Error: testFileLog,
			},
		},
	}

	gLog.Debug("hello world")
	gLog.Warn("hello world")
	gLog.Info("hello world")
	gLog.Error("hello world")
}
