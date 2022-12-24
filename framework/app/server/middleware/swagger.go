/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-15 21:07:47
 */

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Swagger 文档中间件
func Swagger(ctx *gin.Context) {
	path := ctx.Request.URL.Path

	// 非 /doc 开头不是文档
	if !strings.HasPrefix(path, "/doc/") {
		return
	}

	// 过滤掉非法的 /doc/? 路径
	suffix := filepath.Base(path)
	if suffix != "doc" && suffix != "doc.json" && suffix != "redoc.js" {
		ctx.Redirect(http.StatusFound, "/doc")
		return
	}

	_, _ = ctx.Writer.Write(redoc(suffix))
}

// 文档
func redoc(suffix string) []byte {
	if suffix == "doc.json" {
		data, _ := os.ReadFile("docs/swagger.json")
		return data
	}

	if suffix == "redoc.js" {
		_, file, _, _ := runtime.Caller(0)
		path := filepath.Dir(file)

		data, _ := os.ReadFile(path + "/swagger_redoc.js")
		return data
	}

	json := "swagger/doc.json" // "https://petstore.swagger.io/v2/swagger.json"

	return []byte(fmt.Sprintf(`<!DOCTYPE html>
<html>
  <head>
	<meta charset="utf-8"/>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
	  body {
		margin: 0;
		padding: 0;
	  }
	</style>
  </head>
  <body>
	<redoc spec-url='%s'></redoc>
	<script src="/doc/redoc.js"> </script>
  </body>
</html>`, json))
}
