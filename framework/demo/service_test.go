package demo

import "testing"

func TestHelloWorld(t *testing.T) {
	// 创建
	id := create("13012340001")
	dump()

	// 更新
	update("13012340001", "000000")
	dump()

	// 创建
	id = create("13012340002")
	dump()

	// 创建
	id = create("13012340003")
	dump()

	// 删除
	remove(id)
	dump()
}
