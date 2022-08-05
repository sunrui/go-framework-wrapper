/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

// swagger 配置对象
type swagger struct {
	Enable bool `json:"enable"` // 是否启用
}

// Swagger 配置
func Swagger() *swagger {
	return &conf.Swagger
}
