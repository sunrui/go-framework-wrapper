/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/16 14:11:16
 */

package app

import (
	"bytes"
	"framework/config"
	"io/ioutil"
	"os/exec"
)

// redoc 文档中间件
func redoc(suffix string) []byte {
	if suffix == "doc.json" {
		data, _ := ioutil.ReadFile("docs/swagger.json")
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
			<redoc spec-url='swagger/doc.json'></redoc>
  			<!--
				<redoc spec-url='http://petstore.swagger.io/v2/swagger.json'></redoc>
			-->
			<script src="https://rebilly.github.io/ReDoc/releases/latest/redoc.min.js"> </script>
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
	if !config.IsDebug() {
		commandExec("swag", "init", "--parseDependency")
		commandExec("swag", "fmt")
	}
}
