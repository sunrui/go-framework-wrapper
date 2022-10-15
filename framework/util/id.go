/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-05-10 11:40:43
 */

package util

import (
	"github.com/google/uuid"
	"github.com/matoous/go-nanoid/v2"
	"strings"
)

// CreateUuid 创建唯一 uuid
func CreateUuid() string {
	id := uuid.NewString()
	id = strings.ToUpper(id)
	id = strings.ReplaceAll(id, "-", "")

	return id
}

// CreateNanoid 创建唯一 nanoid
func CreateNanoid() string {
	const dictionary = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if id, err := gonanoid.Generate(dictionary, 12); err != nil {
		panic(err.Error())
	} else {
		return id
	}
}
