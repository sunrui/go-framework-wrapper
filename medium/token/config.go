/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-15 00:26:15
 */

package token

type config struct {
	JwtSecret string `json:"JwtSecret"` // jwt 密钥
}
