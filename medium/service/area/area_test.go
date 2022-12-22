/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-22 16:39:28
 */

package area

import (
	"fmt"
	"framework/app/util"
	"testing"
)

var country Country

// TestMain 初始化前准备
func TestMain(m *testing.M) {
	country = NewCountry()
	fmt.Println(util.ToJson(country, false))

	m.Run()
}

// TestCountry_GetProvinces 获取所有的省
func TestCountry_GetProvinces(t *testing.T) {
	provinces := country.GetProvinces()
	fmt.Println(util.ToJson(provinces, true))
}

func TestCountry_GetCities(t *testing.T) {
	provinces := country.GetProvinces()
	fmt.Println(util.ToJson(provinces, false))

	var province Province
	for _, provinceOne := range provinces {
		if provinceOne.Name == "山东省" {
			province = provinceOne
			break
		}
	}
	fmt.Println(util.ToJson(province, true))

	var city City
	for _, cityOne := range country.GetCities(province.Id) {
		if cityOne.Name == "枣庄市" {
			city = cityOne
			break
		}
	}
	fmt.Println(util.ToJson(city, true))

	counties := country.GetCounties(city.Id)
	fmt.Println(util.ToJson(counties, true))
}

func TestCountry_GetCounties(t *testing.T) {
	counties := country.GetCounties(370400)
	fmt.Println(util.ToJson(counties, true))
}

func TestCountry_Get(t *testing.T) {
	province, city, county := country.Get(370402)
	fmt.Println(province, city, county)
}
