/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/08/01 20:50:01
 */

package response

import (
	"framework/config"
	"framework/proto/result"
	"framework/proto/token"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

// 记录出错日志
func logResult(ctx *gin.Context, r result.Result) {
	// 协议
	var http = func(ctx *gin.Context) string {
		if ctx.Request.TLS != nil {
			return "https://"
		} else {
			return "http://"
		}
	}

	// 等级
	var level = func(r result.Result) string {
		if r.Code == result.Ok.Code {
			return "TRACE"
		} else {
			return "ERROR"
		}
	}

	var buffer string

	// 时间 - 等级 - IP
	buffer += " - " + level(r) + " - " + ctx.ClientIP()

	// userId
	if t, err := token.GetToken(ctx); err == nil {
		buffer += " - " + t.UserId
	}

	// 换行
	buffer += "\n\n"

	// method http://host:port?query protocol
	buffer += ctx.Request.Method + " " + http(ctx) + ctx.Request.Host + ctx.Request.RequestURI + " " + ctx.Request.Proto + "\n"

	// header
	for key, value := range ctx.Request.Header {
		buffer += key + ": " + value[0] + "\n"
	}

	// 空一行
	buffer += "\n"

	// body
	body, exists := ctx.Get("body")
	if exists {
		buffer += body.(string) + "\n"
	} else {
		buffer += "<null>\n"
	}

	// 空一行
	buffer += "\n"

	// 结果
	buffer += r.String(true) + "\n"

	// 打印输出
	log.Println(buffer)
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
		var fileName string
		if config.IsDebug() {
			fileName = time.Now().Format("2006-01-02")
		} else {
			fileName = time.Now().Format("2006-01-02 15:04:05")
		}

		if file, err := os.Create(logPath + "/access - " + fileName + ".log"); err != nil {
			panic(err.Error())
		} else {
			return file
		}
	}

	log.SetOutput(io.MultiWriter(createFile()))
	log.SetFlags(log.Ldate | log.Ltime)
}
