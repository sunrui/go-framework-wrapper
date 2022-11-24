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

func TestResult_WithData(t *testing.T) {
	fmt.Println(Ok.WithData(M{
		"hello": "world",
	}))
}

func TestNoAuth(t *testing.T) {
	r := NoAuth.WithMessage("hello world")
	fmt.Println(r)
}
