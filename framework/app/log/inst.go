/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-16 23:07:01
 */

package log

import (
	"config"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

// 日志实例
var logger *logrus.Logger

func init() {
	// 建立日志目录
	if _, err := os.Stat(config.Inst().Log.Directory); err != nil {
		if err = os.Mkdir(config.Inst().Log.Directory, os.ModePerm); err != nil {
			panic(err.Error())
		}
	}

	// 文件全路径
	fileName := path.Join(config.Inst().Log.Directory, config.Inst().Log.File+".log")

	// 创建日志实例
	logger = logrus.New()
	logger.SetLevel(config.Inst().Log.Level)
	logger.SetOutput(io.MultiWriter(func() *os.File {
		if file, err := os.Create(fileName); err != nil {
			panic(err.Error())
		} else {
			return file
		}
	}()))

	// 可循环的日志配置
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		path.Join(config.Inst().Log.Directory, config.Inst().Log.File)+".%Y-%m-%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	// hook 机制的设置
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writerMap, &myFormatter{}))
	logger.SetFormatter(&myFormatter{})
}