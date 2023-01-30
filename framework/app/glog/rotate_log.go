/*
 * Copyright (c) 2023 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2023-01-30 10:49:02
 */

package glog

import (
	"bytes"
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
	Enable    bool         `json:"enable"`    //  启用
	Directory string       `json:"directory"` // 路径
	Level     logrus.Level `json:"level"`     // 等级
}

// RotateLog 日志
type RotateLog struct {
	*logrus.Logger // log
}

// 自定义格式化
type rotateLogFormatter struct {
}

// Format 格式化
func (m *rotateLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteString(entry.Message)

	return b.Bytes(), nil
}

// New 创建
func New(config Config, directory string, filePrefix string) (*RotateLog, error) {
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

	// 日志配置
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

	// hook 设置
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	log.AddHook(lfshook.NewHook(writerMap, &rotateLogFormatter{}))
	log.SetFormatter(&rotateLogFormatter{})

	return &RotateLog{
		log,
	}, nil
}
