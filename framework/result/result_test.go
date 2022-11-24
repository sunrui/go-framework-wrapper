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

func TestOk(t *testing.T) {
	fmt.Println(Ok())
}

func TestOkWithData(t *testing.T) {
	fmt.Println(OkWithData[string]("hello world"))
}

func TestOkWithIdData(t *testing.T) {
	fmt.Println(OkWithMapData(M{
		"id": "my id",
	}))
}

func TestOkWithDataAndPagination(t *testing.T) {
	fmt.Println(OkWithDataAndPagination[string]("hello world", &Pagination{
		Page:      0,
		PageSize:  0,
		TotalPage: 0,
		TotalSize: 0,
	}))
}

func TestNoAuth(t *testing.T) {
	r := NoAuth.WithMessage("hello world")
	fmt.Println(r)
}
