/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/07/28 22:23:28
 */

package utils

import (
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"log"
)

// LogResult 日志记录
func LogResult(ctx *gin.Context, result result.Result) {
	log.Println(result)

	// TODO  默认不打印到文件里
}
