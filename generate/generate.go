/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/16 16:07:16
 */

package main

import (
	"framework/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

		// 拷贝项目
		err = utils.CopyDirectory(src, dst)
		if err != nil {
			return err
		}

		// 将项目下面的 template 更为 model 对象相关
		src = filepath.Dir(dst) + "/template"
		dst = filepath.Dir(dst) + "/" + strings.ToLower(model)

		// 移除旧的 dst 生成好的模板
		_ = os.RemoveAll(dst)

		// 重新改名 dst 生成的模板
		_ = os.Rename(src, dst)

		// 处理每一个文件
		if list, err := ioutil.ReadDir(dst); err == nil {
			// 递归每一个文件
			for _, item := range list {
				fileName := filepath.Join(dst, item.Name())

				fileBytes, err := ioutil.ReadFile(fileName)
				if err != nil {
					return err
				}

				// 处理内容部分
				fileContent := string(fileBytes)
				fileContent = strings.ReplaceAll(fileContent, "Template", model)
				fileContent = strings.ReplaceAll(fileContent, strings.ToLower("Template"), strings.ToLower(model))
				fileContent = strings.ReplaceAll(fileContent, "模板", modelName)
				fileContent = strings.ReplaceAll(fileContent, "generate/service/core/test", "service/core/test")

				err = ioutil.WriteFile(fileName, []byte(fileContent), os.ModeDevice)
				if err != nil {
					return err
				}

				println(fileName)
			}
		} else {
			return err
		}
	}

	return nil
}
