package example

import (
	"container/list"
	"encoding/json"
)

var collection = list.New()

// 创建
func create(phone string) string {
	sms := Sms{
		Phone: phone,
		Code:  "123456",
	}

	collection.PushBack(&sms)
	return sms.Code
}

// 更新
func update(phone string, code string) bool {
	for i := collection.Front(); i != nil; i = i.Next() {
		sms := i.Value.(*Sms)

		if sms.Phone == phone {
			sms.Code = code
			return true
		}
	}

	return false
}

// 删除
func remove(phone string) bool {
	for i := collection.Front(); i != nil; i = i.Next() {
		sms := i.Value.(*Sms)

		if sms.Phone == phone {
			collection.Remove(i)
			return true
		}
	}

	return false
}

// 查询
func find(phone string) *Sms {
	for i := collection.Front(); i != nil; i = i.Next() {
		sms := i.Value.(*Sms)

		if sms.Phone == phone {
			return sms
		}
	}

	return nil
}

// 查询所有
func all() []*Sms {
	var smsArray = make([]*Sms, collection.Len())
	index := 0
	for i := collection.Front(); i != nil; i = i.Next() {
		smsArray[index] = i.Value.(*Sms)
		index++
	}

	return smsArray
}

// dump
func dump() {
	bytes, _ := json.Marshal(all())
	println(string(bytes))
}
