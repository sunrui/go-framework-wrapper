/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-07 00:42:07
 */

package mysql

// FindOne 查找某一个
func FindOne[T any](query interface{}, args ...interface{}) *T {
	var dst T

	if r := Inst.DB.Where(query, args).Find(&dst); r.Error != nil {
		panic(r.Error.Error())
	} else if r.RowsAffected == 1 {
		return &dst
	}

	return nil
}
