/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/03 13:44:03
 */

package common

import (
	"framework/proto/result"
	"github.com/gin-gonic/gin"
)

// @Summary 获取 Index 测试
// @Tags    通用
// @Produce json
// @Param   enumstring query    string                      false "string enums"      Enums(A, B, C)
// @Param   enumint    query    int                         false "int enums"         Enums(1, 2, 3)
// @Param   enumnumber query    number                      false "int enums"         Enums(1.1, 1.2, 1.3)
// @Param   string     query    string                      false "string valid"      minlength(5) maxlength(10)
// @Param   int        query    int                         false "int valid"         mininum(1)   maxinum(10)
// @Param   default    query    string                      false "string default"    default(A)
// @Param   enumstring query    string                      false "string enums"      Enums(A, B, C) default(A)
// @Param   example    query    string                      false "string example"    example(string)
// @Param   collection query    []string                    false "string collection" collectionFormat(multi)
// @Param   extensions query    []string                    false "string collection" extensions(x-example=test,x-nullable)
// @Success 200        {object} result.Result{data=sms.Sms} true
// @Router  /public/common/ [get]
func getIndex(ctx *gin.Context) result.Result {
	return result.Ok.WithData("hello world")
}
