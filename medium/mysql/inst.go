/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-07 00:26:26
 */

package mysql

import "medium/configer"

var Inst *Mysql

func init() {
	Inst = newMysql(configer.Load[config]())
}
