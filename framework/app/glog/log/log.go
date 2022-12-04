/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-16 21:44:52
 */

package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

const (
	maxAge       = 7 * 24 * time.Hour // 最长过期时间
	rotationTime = 24 * time.Hour     // 切割时间间隔
)

// Config 配置
type Config struct {
	Directory string       `json:"directory"` // 路径
	Level     logrus.Level `json:"level"`     // 等级
}

// Log 日志
type Log struct {
	*logrus.Logger // log
}

// New 创建
func New(config Config, directory string, filePrefix string) (*Log, error) {
	// 路径
	filePath := config.Directory + "/" + directory
	// 文件
	fileName := path.Join(filePath, filePrefix+".log")

	var err error

	// 建立目录
	if _, err = os.Stat(filePath); err != nil {
		if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
			return nil, err
		}
	}

	// 创建文件
	var file *os.File
	if file, err = os.Create(fileName); err != nil {
		return nil, err
	}

	// 创建实例
	log := logrus.New()
	log.SetLevel(config.Level)
	log.SetOutput(file)

	// 可循环的配置
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		path.Join(filePath, filePrefix)+".%Y-%m-%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间
		rotatelogs.WithMaxAge(maxAge),
		// 设置切割时间间隔
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

	return &Log{
		log,
	}, nil
}
