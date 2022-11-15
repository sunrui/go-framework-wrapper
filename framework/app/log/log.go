/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-14 23:23:28
 */

package log

import (
	"config"
	"fmt"
	"framework/app/request"
	"framework/result"
	"framework/token"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

// SetConfig 设置配置项
func SetConfig(_conf config.Log) {
	config.Inst().Log = _conf

	if config.Inst().Log.Level == config.LogNone {
		log.SetOutput(nil)
		return
	}

	// 建立日志目录
	if _, err := os.Stat(config.Inst().Log.Directory); err != nil {
		if err = os.Mkdir(config.Inst().Log.Directory, os.ModePerm); err != nil {
			panic(err.Error())
		}
	}

	// 每次启动的时候建立新文件
	log.SetOutput(io.MultiWriter(func() *os.File {
		var timeLayout string
		if config.IsDev() {
			timeLayout = "2006-01-02"
		} else {
			timeLayout = "2006-01-02 15:04:05"
		}

		fileName := time.Now().Format(timeLayout)
		if file, err := os.Create(config.Inst().Log.Directory + "/" + fileName + ".log"); err != nil {
			panic(err.Error())
		} else {
			return file
		}
	}()))

	log.SetFlags(log.Ldate | log.Ltime)
}

// 获取结果内容
func getResult(ctx *gin.Context, r result.Result[any]) string {
	// 获取 request 对象
	req := request.Get(ctx)

	// 等级
	var logLevel config.LogLevel
	if r.Code == result.OK {
		logLevel = config.LogInfo
	} else {
		logLevel = config.LogError
	}

	var buffer string

	// 时间 - 等级 - IP
	buffer = " - " + string(logLevel) + " - " + req.Ip

	// userId
	if userId := token.GetUserId(); userId != nil {
		buffer += " - userId(" + *userId + ")"
	}

	// 换行
	buffer += "\n"

	// method http://host:port?query protocol
	buffer += req.Method + " " + req.Uri + " " + req.Proto + "\n"

	// header
	for _, header := range req.Headers {
		buffer += header + "\n"
	}

	// 空一行
	buffer += "\n"

	// body
	if req.Body != nil {
		buffer += *req.Body + "\n"
	} else {
		buffer += "<null>\n"
	}

	// 空一行
	buffer += "\n"

	// 结果
	buffer += r.String() + "\n"

	return buffer
}

// Write 写入文件
func Write(ctx *gin.Context, r result.Result[any]) {
	if config.Inst().Log.Level == config.LogNone {
		return
	}

	stream := getResult(ctx, r)
	log.Println(stream)
	fmt.Println(stream)
}

func init() {
	SetConfig(config.Inst().Log)
}
