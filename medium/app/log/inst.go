/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-14 23:31:56
 */

package log

// 当前等级
var levelType LevelType

// SetLevelType 设置等级
func SetLevelType(_levelType LevelType) {
	levelType = _levelType
}

// GetLevelType 获取等级
func GetLevelType() LevelType {
	return levelType
}
