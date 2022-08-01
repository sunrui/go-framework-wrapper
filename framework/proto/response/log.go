/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/08/01 20:50:01
 */

package response

import (
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

// 记录出错日志
func logResult(ctx *gin.Context, r result.Result) {
	var level = func(r result.Result) string {
		if r.Code == result.Ok.Code {
			return "INFO"
		} else {
			return "ERROR"
		}
	}

	log.Print(" - " + level(r))
}

// 初始化
func init() {
	const logPath = "log"
	var err error

	// 建立日志目录
	if _, err = os.Stat(logPath); err != nil {
		if err = os.Mkdir(logPath, os.ModePerm); err != nil {
			panic(err.Error())
		}
	}

	// 每次启动的时候建立新文件
	var createFile = func() *os.File {
		fileName := time.Now().Format("2006-01-02 15:04:05")
		if file, err := os.Create(logPath + "/error - [" + fileName + "].log"); err != nil {
			panic(err.Error())
		} else {
			return file
		}
	}

	log.SetOutput(io.MultiWriter(createFile()))
	log.SetFlags(log.Ldate | log.Ltime)
}
