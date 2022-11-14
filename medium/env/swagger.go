/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-14 23:04:00
 */

package env

import (
	"bytes"
	"os/exec"
)

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

// Swag 执行
func Swag() {
	commandExec("swag", "fmt")
	commandExec("swag", "init", "--parseDependency")
}
