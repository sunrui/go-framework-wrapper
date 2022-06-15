/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-06-15 15:22:23
 */

package _sample

import (
	"framework/proto/response"
	"github.com/gin-gonic/gin"
)

func getSample(ctx *gin.Context) {
	response.New(ctx).Data("hello world")
}

func putSample(ctx *gin.Context) {
	response.New(ctx).Data("hello world")
}
