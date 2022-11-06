/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-07 00:26:26
 */

package mysql

import "medium/config"

var Inst *Mysql

func init() {
	Inst = NewMysql(config.Load[conf]())
}
