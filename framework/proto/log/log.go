/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/08/02 00:41:02
 */

package log

import (
	"framework/config"
	"framework/proto/request"
	"framework/proto/result"
	"framework/proto/token"
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

// Set 设置
func Set(enable bool, level Level) {
	config.Log().Enable = enable
	config.Log().Level = string(level)
}

// WriteResult 写入结果
func WriteResult(ctx *gin.Context, r result.Result) {
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
	if level == LevelTrace && string(level) != config.Log().Level {
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

// 初始化
func init() {
	const logPath = "log"

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
		if file, err := os.Create(logPath + "/access - " + fileName + ".log"); err != nil {
			panic(err.Error())
		} else {
			return file
		}
	}

	log.SetOutput(io.MultiWriter(createFile()))
	log.SetFlags(log.Ldate | log.Ltime)
}
