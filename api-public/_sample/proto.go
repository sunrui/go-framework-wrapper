/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-06-15 15:23:26
 */

package _sample

import "framework/proto/request"

type getSampleReq struct {
	request.PageRequest
	Name string `json:"name" validate:"required"` // 名称
}
