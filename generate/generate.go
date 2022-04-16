/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/16 13:57:16
 */

package main

import (
	"framework/utils"
	"os"
)

// 生成模板
func generate(model, modelName string) error {
	// 获取当前项目根目录
	pwd, _ := os.Getwd()

	// 拷贝 api-admin
	err := utils.CopyDirectory(pwd+"/api-admin", pwd+"/../api-admin")
	if err != nil {
		return err
	}

	return nil
}

// 初始化
func init() {
	err := generate("Test", "测试")
	if err != nil {
		println(err.Error())
	}
}
