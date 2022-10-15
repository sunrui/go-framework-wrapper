/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-06-14 10:11:05
 */

package middleware

import (
	"bytes"
	"framework/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/exec"
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

// 文档中间件
func redoc(suffix string) []byte {
	if suffix == "doc.json" {
		data, _ := os.ReadFile("docs/swagger.json")
		return data
	}

	if suffix == "redoc.js" {
		_, file, _, _ := runtime.Caller(0)
		path := filepath.Dir(file)

		data, _ := os.ReadFile(path + "/doc_redoc.js")
		return data
	}

	return []byte(`
			<!DOCTYPE html>
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
				<!--<redoc spec-url='https://petstore.swagger.io/v2/swagger.json'></redoc>-->
				<redoc spec-url='swagger/doc.json'></redoc>
				<script src="/doc/redoc.js"> </script>
			  </body>
			</html>
		`)
}

// 执行命令行
func commandExec(name string, arg ...string) {
	var out bytes.Buffer
	var err error

	cmd := exec.Command(name, arg...)
	cmd.Stdout = &out

	if err = cmd.Start(); err != nil {
		panic(err.Error())
	}

	if err = cmd.Wait(); err != nil {
		panic(err.Error())
	}

	println(out.String())
}

// 执行 swag 更新文档
func init() {
	if config.Cur().Swagger.Enable {
		commandExec("swag", "fmt")
		commandExec("swag", "init", "--parseDependency")
	}
}
