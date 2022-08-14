/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/16 15:08:16
 */

package main

// @title    Medium 模板接口文档
// @version  1.0
// @host     127.0.0.1:8080
// @BasePath
func main() {
	if err := runGenerate("Test", "测试", "sunrui"); err != nil {
		println(err.Error())
	}
}
