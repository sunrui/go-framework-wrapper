/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-14 23:23:28
 */

package log

import (
	"io"
	"log"
	"os"
	"time"
)

// Init 初始化日志
func Init(level LevelType, filePath string) {
	if level == "NONE" {
		log.SetOutput(nil)
		return
	}

	// 建立日志目录
	if _, err := os.Stat(filePath); err != nil {
		if err = os.Mkdir(filePath, os.ModePerm); err != nil {
			panic(err.Error())
		}
	}

	// 每次启动的时候建立新文件
	log.SetOutput(io.MultiWriter(func() *os.File {
		fileName := time.Now().Format("2006-01-02 15:04:05")
		if file, err := os.Create(filePath + "/" + fileName + ".log"); err != nil {
			panic(err.Error())
		} else {
			return file
		}
	}()))

	log.SetFlags(log.Ldate | log.Ltime)
}
