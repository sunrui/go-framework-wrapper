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

// 地区
type areaBase struct {
	Id   int    `json:"id,omitempty"`   // 编码
	Name string `json:"name,omitempty"` // 名称
}

// Country 国家
type Country struct {
	areaBase             // 地区
	Provinces []Province `json:"provinces,omitempty"` // 省
}

// Province 省
type Province struct {
	areaBase        // 地区
	Cities   []City `json:"cities,omitempty"` // 市
}

// City 市
type City struct {
	areaBase          // 地区
	Counties []County `json:"counties,omitempty"` // 区县
}

// County 区县
type County struct {
	areaBase // 地区
}

// NewCountry 创建国家
func NewCountry() (country Country) {
	_, file, _, _ := runtime.Caller(0)
	path := filepath.Dir(file)

	if stream, err := os.ReadFile(path + "/area.json"); err != nil {
		panic(err.Error())
	} else if err = json.Unmarshal(stream, &country); err != nil {
		panic(err.Error())
	}

	return
}

// GetProvinces 获取省
func (country Country) GetProvinces() []Province {
	var provinces []Province
	for _, provinceOne := range country.Provinces {
		provinceOne.Cities = nil
		provinces = append(provinces, provinceOne)
	}

	return provinces
}

// GetCities 获取市
func (country Country) GetCities(provinceId int) []City {
	var province *Province
	for _, provinceOne := range country.Provinces {
		if provinceOne.Id == provinceId {
			province = &provinceOne
			break
		}
	}

	// 不存在的省默认返回 nil
	if province == nil {
		return nil
	}

	var cities []City
	for _, cityOne := range province.Cities {
		cityOne.Counties = nil
		cities = append(cities, cityOne)
	}

	return cities
}

// GetCounties 获取地区
func (country Country) GetCounties(cityId int) []County {
	var city *City
	for _, provinceOne := range country.Provinces {
		//  用前两位过滤掉省
		if provinceOne.Id/10000 != cityId/10000 {
			continue
		}

		for _, cityOne := range provinceOne.Cities {
			if cityOne.Id == cityId {
				city = &cityOne
				break
			}
		}
	}

	// 不存在的城市默认返回 nil
	if city == nil {
		return nil
	} else {
		return city.Counties
	}
}

// Get 获取省、市、区
func (country Country) Get(areaId int) (province *Province, city *City, county *County) {
	for _, provinceOne := range country.Provinces {
		//  用前两位过滤掉省
		if provinceOne.Id/10000 != areaId/10000 {
			continue
		}

		if provinceOne.Id == areaId {
			province = &provinceOne
			return
		}

		for _, cityOne := range provinceOne.Cities {
			//  用前四位过滤掉市
			if cityOne.Id/100 != areaId/100 {
				continue
			}

			if cityOne.Id == areaId {
				province = &provinceOne
				city = &cityOne
				return
			}

			for _, countyOne := range cityOne.Counties {
				if countyOne.Id == areaId {
					province = &provinceOne
					city = &cityOne
					county = &countyOne
					return
				}
			}
		}
	}

	return
}
