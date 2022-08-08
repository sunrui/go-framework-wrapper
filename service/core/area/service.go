/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 14:26:31
 */

package area

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

// 国家对象
var country Country

// GetCountry 获取国家
func GetCountry() Country {
	return country
}

// GetProvinces 获取省
func GetProvinces() []Province {
	var provinces []Province

	for _, province := range country.Provinces {
		province.Cities = nil
		provinces = append(provinces, province)
	}

	return provinces
}

// GetCity 获取市
func GetCity(provinceId int) []City {
	// 根据省 id 获取省节点
	var province *Province
	for _, one := range country.Provinces {
		if one.Id == provinceId {
			province = &one
			break
		}
	}

	// 不存在的省默认返回 nil
	if province == nil {
		return nil
	}

	var cities []City
	for _, city := range province.Cities {
		city.Areas = nil
		cities = append(cities, city)
	}

	return cities
}

// GetArea 获取地区
func GetArea(cityId int) []Area {
	// 根据市 id 获取市节点
	var city *City
	for _, province := range country.Provinces {
		for _, one := range province.Cities {
			if one.Id == cityId {
				city = &one
				break
			}
		}
	}

	// 不存在的城市默认返回 nil
	if city == nil {
		return nil
	} else {
		return city.Areas
	}
}

// 加载当前配置
func init() {
	_, file, _, _ := runtime.Caller(0)
	path := filepath.Dir(file)

	if stream, err := os.ReadFile(path + "/area.json"); err != nil {
		panic(err.Error())
	} else if err = json.Unmarshal(stream, &country); err != nil {
		panic(err.Error())
	}
}
