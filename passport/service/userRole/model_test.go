/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-20 18:15:08
 */

package userRole

import (
	"framework/mysql"
	"testing"
)

func TestSave(t *testing.T) {
	userRole := UserRole{
		UserId: "userId-0",
		Type:   AdminType,
	}

	mysql.Save(&userRole)
}

func TestFind(t *testing.T) {

}
