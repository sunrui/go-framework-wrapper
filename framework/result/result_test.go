/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-14 22:48:37
 */

package result

import (
	"fmt"
	"testing"
)

func TestIdData(t *testing.T) {
	fmt.Println(Result[any]{
		Code:    OK,
		Message: "增加成功",
		Data:    IdData("my id"),
	})
}

func TestKeyValueData(t *testing.T) {
	fmt.Println(Result[any]{
		Code:    CONFLICT,
		Message: "冲突了",
		Data:    KeyValueData("my key", "mey value"),
	})
}
