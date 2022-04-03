/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/03 21:03:03
 */

package app

import (
	"bytes"
	_ "medium-server-go/docs"
	"os/exec"
)

// 执行命令行
func commandExec(name string, argv string) {
	var out bytes.Buffer

	cmd := exec.Command(name, argv)
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
	commandExec("swag", "init")
	commandExec("swag", "fmt")
}
