/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/16 15:21:16
 */

package main

import (
	"framework/utils"
	"os"
	"path/filepath"
)

// 生成模板
func runGenerate(model, modelName string) error {
	var err error
	// 获取当前项目根目录
	pwd, _ := os.Getwd()

	var projects = []string{
		"api-admin",
		"api-user",
		"service/core",
	}

	for _, project := range projects {
		src := pwd + "/" + project + "/template"
		dst := pwd + "/../" + project + "/template"

		// 移除现有的模板
		_ = os.RemoveAll(dst)

		continue

		// 拷贝项目
		err = utils.CopyDirectory(src, dst)
		if err != nil {
			return err
		}

		// 将项目下面的 template 更为 model 对象相关
		dir := filepath.Dir(project)
		_ = os.Rename(dir+"/template", dir+"/"+model)

		return err
	}

	return nil
}
