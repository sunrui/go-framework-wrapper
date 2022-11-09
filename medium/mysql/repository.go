/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-07 00:42:07
 */

package mysql

import (
	"errors"
	"fmt"
)

// FindById 根据 id 查找
func FindById[T any](id string) *T {
	var dst T

	if r := Inst.DB.Model(dst).Where("id = ?", id).Find(&dst); r.Error != nil {
		panic(r.Error.Error())
	} else if r.RowsAffected == 1 {
		return &dst
	} else {
		return nil
	}
}

// FindOne 根据查件查找一个
func FindOne[T any](query interface{}, args ...interface{}) *T {
	var dst []T

	if r := Inst.DB.Model(dst).Where(query, args).Find(&dst); r.Error != nil {
		panic(r.Error.Error())
	} else if r.RowsAffected > 1 {
		panic(errors.New(fmt.Sprintf("find %d record", r.RowsAffected)))
	} else if r.RowsAffected == 1 {
		return &dst[0]
	} else {
		return nil
	}
}

// FindMany 根据查件查找一个或多个
func FindMany[T any](query interface{}, args ...interface{}) []T {
	var dst []T

	if r := Inst.DB.Model(dst).Where(query, args).Find(&dst); r.Error != nil {
		panic(r.Error.Error())
	} else if r.RowsAffected >= 1 {
		return dst
	} else {
		return nil
	}
}

// DeleteById 根据 id 删除
func DeleteById[T any](id string) bool {
	var dst T

	if r := Inst.DB.Model(dst).Where("id = ?", id).Delete(&dst); r.Error != nil {
		panic(r.Error.Error())
	} else if r.RowsAffected >= 1 {
		return true
	} else {
		return false
	}
}
