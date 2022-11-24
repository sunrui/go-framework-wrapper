/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-10-15 09:11:59
 */

package result

const (
	okCode    = "Ok"
	okMessage = "成功"
)

// Ok 成功
func Ok() Result[any] {
	return newResult[any](okCode, okMessage)
}

// OkWithData 成功并设置 data
func OkWithData[T any](data T) Result[T] {
	ok := newResult[T](okCode, okMessage)
	ok.Data = data
	return ok
}

// M map 对象
type M map[string]any

// OkWithMapData 成功并设置 map data
func OkWithMapData(data M) Result[M] {
	return OkWithData[M](data)
}

// OkWithDataAndPagination 成功并设置 data 和 pagination
func OkWithDataAndPagination[T any](data T, pagination *Pagination) Result[T] {
	ok := OkWithData[T](data)
	ok.Pagination = pagination
	return ok
}
