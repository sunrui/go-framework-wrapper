/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-01 01:34:17
 */

package swag

import (
	"io"
	"log"
	"os"
	"os/exec"
	"testing"
)

// 执行命令行
func commandExec(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	var stdout io.ReadCloser
	var readByte []byte
	var err error

	if stdout, err = cmd.StdoutPipe(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = stdout.Close()
	}()

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if readByte, err = io.ReadAll(stdout); err != nil {
		log.Fatal(err)
	} else if len(readByte) > 0 {
		log.Println(string(readByte))
	}
}

// TestSwag 执行
func TestSwag(t *testing.T) {
	_ = os.Chdir("../../passport/api/public")
	dir, _ := os.Getwd()
	t.Log("swag dir: " + dir)

	commandExec("swag", "fmt")
	commandExec("swag", "init", "--parseDependency")
}
