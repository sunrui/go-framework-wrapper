/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-16 23:07:10
 */

package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
)

// 自定义格式器
type myFormatter struct {
}

// Format 格式化
func (m *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteString(fmt.Sprintf("%s - %s", entry.Level, entry.Message))

	return b.Bytes(), nil
}
