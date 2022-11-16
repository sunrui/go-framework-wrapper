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

type myFormatter struct {
}

func (m *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	b.WriteString(fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message))

	return b.Bytes(), nil
}
