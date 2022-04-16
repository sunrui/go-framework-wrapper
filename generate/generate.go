/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/16 14:56:16
 */

package generate

import (
	"framework/utils"
	"os"
	"path/filepath"
)

var model string
var modelName string

// 处理文件
func processFile(pwd, project string) error {
	var err error

	// 移除现有的模板
	_ = os.RemoveAll(pwd + "/../" + project + "/")

	// 拷贝项目
	err = utils.CopyDirectory(pwd+"/"+project+"/", pwd+"/../"+project+"/")
	if err != nil {
		return err
	}

	// 将项目下面的 template 更为 model 对象相关
	dir := filepath.Dir(project)
	_ = os.Rename(dir+"/template", dir+"/"+model)

	return err
}

// 生成模板
func Run(model, modelName string) error {
	var err error
	// 获取当前项目根目录
	pwd, _ := os.Getwd()

	var projects = []string{
		"api-admin",
		"api-user",
		"service",
	}

	for _, project := range projects {
		err = processFile(pwd, project)
		if err != nil {
			return err
		}
	}

	return nil
}
