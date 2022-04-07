/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/03 21:03:03
 */

package app

import (
	"bytes"
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
			<script src="https://cdn.jsdelivr.net/npm/redoc@latest/bundles/redoc.standalone.js"> </script>
		  </body>
		</html>
	`)
}

// 执行命令行
func commandExec(name string, arg ...string) {
	var out bytes.Buffer

	cmd := exec.Command(name, arg...)
	cmd.Stdout = &out

	err := cmd.Start()
	if err != nil {
		panic(err.Error())
	}

	err = cmd.Wait()
	if err != nil {
		panic(err.Error())
	}

	println(out.String())
}

// 执行 swag 更新文档
func init() {
	commandExec("swag", "init", "--parseDependency")
	commandExec("swag", "fmt")
}
