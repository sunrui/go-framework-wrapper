/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-16 21:44:52
 */

package log

import (
	"framework/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

const (
	maxAge       = 7 * 24 * time.Hour // 最长过期时间
	rotationTime = 24 * time.Hour     // 最大回滚时长
)

// NewLog 创建日志
func NewLog(logConfig config.Log, directory string, filePrefix string) *logrus.Logger {
	// 路径
	filePath := logConfig.Directory + "/" + directory
	// 文件
	fileName := path.Join(filePath, filePrefix+".log")

	// 建立日志目录
	if _, err := os.Stat(filePath); err != nil {
		if err = os.Mkdir(filePath, os.ModePerm); err != nil {
			panic(err.Error())
		}
	}

	// 创建日志实例
	log := logrus.New()
	log.SetLevel(logConfig.Level)
	log.SetOutput(io.MultiWriter(func() *os.File {
		if file, err := os.Create(fileName); err != nil {
			panic(err.Error())
		} else {
			return file
		}
	}()))

	// 可循环的日志配置
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		path.Join(filePath, filePrefix)+".%Y-%m-%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间
		rotatelogs.WithMaxAge(maxAge),
		// 设置日志切割时间间隔
		rotatelogs.WithRotationTime(rotationTime),
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

	log.AddHook(lfshook.NewHook(writerMap, &myFormatter{}))
	log.SetFormatter(&myFormatter{})

	return log
}
