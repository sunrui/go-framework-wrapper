/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/12 17:16:12
 */

package sms

import (
	"fmt"
)

// Send 短信发送
func Send(phone string, smsType Type, code string) error {
	echo := fmt.Sprintf("Send - %s, %s, %s", phone, smsType, code)
	fmt.Println(echo)

	return nil
}
