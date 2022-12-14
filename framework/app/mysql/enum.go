/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-14 19:48:31
 */

package mysql

// SortType 排序类型
type SortType string

const (
	AscType  SortType = "ASC"  // 升序
	DescType SortType = "DESC" // 降序
)
