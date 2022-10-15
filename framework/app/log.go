/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-10-15 11:18:48
 */

package app

import (
	"framework/config"
	"framework/request"
	"framework/result"
	"framework/token"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

// Level 日志等级
type Level string

const (
	LevelTrace Level = "TRACE" // trace
	LevelError Level = "ERROR" // error
)

// log 目录
const logPath = "log"

// 初始化日志
func initLog() {
	// 建立日志目录
	if _, err := os.Stat(logPath); err != nil {
		if err = os.Mkdir(logPath, os.ModePerm); err != nil {
			panic(err.Error())
		}
	}

	// 每次启动的时候建立新文件
	var createFile = func() *os.File {
		var timeLayout = func() string {
			if config.IsDev() {
				return "2006-01-02"
			} else {
				return "2006-01-02 15:04:05"
			}
		}()

		fileName := time.Now().Format(timeLayout)
		if file, err := os.Create(logPath + "/" + fileName + ".log"); err != nil {
			panic(err.Error())
		} else {
			return file
		}
	}

	log.SetOutput(io.MultiWriter(createFile()))
	log.SetFlags(log.Ldate | log.Ltime)
}

// SetLog 设置
func SetLog(enable bool, level Level) {
	if enable {
		initLog()
	}

	config.Get().Log.Enable = enable
	config.Get().Log.Level = string(level)
}

// writeResult 写入结果
func writeResult(ctx *gin.Context, r result.Result) {
	// 获取 request 对象
	req := request.GetRequest(ctx)

	// 等级
	var level = func() Level {
		if r.Code == result.Ok.Code {
			return LevelTrace
		} else {
			return LevelError
		}
	}()

	// 判断是否需要输出
	if level == LevelTrace && string(level) != config.Get().Log.Level {
		return
	}

	var buffer string

	// 时间 - 等级 - IP
	buffer += " - " + string(level) + " - " + req.Ip

	// userId
	if t, err := token.Get(ctx); err == nil {
		buffer += " - " + t.UserId
	}

	// 换行
	buffer += "\n\n"

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
	buffer += r.String(true) + "\n"

	// 打印输出
	log.Println(buffer)
}

// WriteLog 记录日志
func WriteLog(ctx *gin.Context, result result.Result) {
	go func() {
		// 异步记录日志
		if config.Get().Log.Enable {
			writeResult(ctx, result)
		}
	}()
}

// 初始化
func init() {
	if config.Get().Log.Enable {
		initLog()
	}
}
