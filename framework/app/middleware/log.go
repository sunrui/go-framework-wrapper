/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-16 21:44:52
 */

package middleware

import (
	"config"
	"fmt"
	"framework/app/request"
	"framework/result"
	"framework/token"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"time"
)

// 获取结果内容
func getBuffer(ctx *gin.Context, r result.Result[any]) string {
	// 获取 request 对象
	req := request.Get(ctx)

	// 等级
	var logLevel logrus.Level
	if r.Code == result.OK {
		logLevel = logrus.DebugLevel
	} else {
		logLevel = logrus.ErrorLevel
	}

	var buffer string

	// 时间 - 等级 - IP
	buffer = " - " + logLevel.String() + " - " + req.Ip

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

var logger *logrus.Logger

// WriteLog 写入文件
func WriteLog(ctx *gin.Context, r result.Result[any]) {
	stream := getBuffer(ctx, r)
	log.Println(stream)
	fmt.Println(stream)

	// 写入日志信息
	//log.WithFields(logrus.Fields{
	//	"url":    ctx.Request.RequestURI,
	//	"method": ctx.Request.Method,
	//	"IP":     ctx.ClientIP(),
	//})
}

func init() {
	fileName := path.Join(config.Inst().Log.Directory, config.Inst().Log.File)

	// 每次启动的时候建立新文件
	log.SetOutput(io.MultiWriter(func() *os.File {
		if file, err := os.Create(fileName); err != nil {
			panic(err.Error())
		} else {
			return file
		}
	}()))

	// 实例化
	logger = logrus.New()
	// 日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置 rotatelogs,实现文件分割
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y-%m-%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour), //以hour为单位的整数
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(1*time.Minute),
	)
	// hook机制的设置
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	//给logrus添加hook
	logger.AddHook(lfshook.NewHook(writerMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

}
