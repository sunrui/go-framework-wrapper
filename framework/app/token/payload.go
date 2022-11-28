/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-28 17:07:17
 */

package token

import "encoding/json"

// Payload 负荷
type Payload struct {
	UserId string `json:"userId"` // 用户 id
}

func (payload Payload) toJson() []byte {
	marshal, _ := json.MarshalIndent(payload, "", "\t")
	return marshal
}
