/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/03 13:44:03
 */

package common

import (
	"framework/app/result"
	"github.com/gin-gonic/gin"
)

// @Summary 获取 Index 测试2
// @Tags    通用2
// @Produce json
// @Param   enumstring query    string   false "string enums"      Enums(A, B, C)
// @Param   enumint    query    int      false "int enums"         Enums(1, 2, 3)
// @Param   enumnumber query    number   false "int enums"         Enums(1.1, 1.2, 1.3)
// @Param   string     query    string   false "string valid"      minlength(5) maxlength(10)
// @Param   int        query    int      false "int valid"         minimum(1)   maximum(10)
// @Param   default    query    string   false "string default"    default(A)
// @Param   example    query    string   false "string example"    example(string)
// @Param   collection query    []string false "string collection" collectionFormat(multi)
// @Param   extensions query    []string false "string collection" extensions(x-example=test,x-nullable)
// @Success 200        {object} result.Result{data=result.Result}
// @Router  /public/common/ [get]
func getError(_ *gin.Context) *result.Result {
	panic("error - " + build.Format("2006-01-02 15:04:05"))
}
