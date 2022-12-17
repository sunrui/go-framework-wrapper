/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/29 18:15:29
 */

package area

// Country 国家
type Country struct {
	Area
	Provinces []Province `json:"provinces"` // 省
}

// Province 省
type Province struct {
	Area
	Cities []City `json:"cities"` // 市
}

// City 市
type City struct {
	Area
	Areas []Area `json:"areas"` // 地区
}

// Area 地区
type Area struct {
	Id   int    `json:"id"`   // 编码
	Name string `json:"name"` // 名称
}
