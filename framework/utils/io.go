/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/13 23:49:13
 */

package utils

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// CopyFile 拷贝文件
func CopyFile(dst, src string) error {
	var srcFile *os.File
	var dstFile *os.File
	var err error

	if srcFile, err = os.Open(src); err != nil {
		return err
	}

	if dstFile, err = os.Open(dst); err != nil {
		return err
	}

	_, err = io.Copy(dstFile, srcFile)

	defer func() {
		_ = srcFile.Close()
		_ = dstFile.Close()
	}()

	return err
}

// CopyDirectory 拷贝文件夹
func CopyDirectory(src, dst string) error {
	fileInfo, e := os.Stat(src)
	if e != nil {
		return e
	}

	if fileInfo.IsDir() {
		// src 是文件夹，那么定义 dst 也是文件夹
		if list, e := ioutil.ReadDir(src); e == nil {
			for _, item := range list {
				src = filepath.Join(src, item.Name())
				dst = filepath.Join(dst, item.Name())

				return CopyDirectory(src, dst)
			}
		} else {
			return e
		}
	} else {
		// src 是文件，那么创建 dst 的文件夹
		dir := filepath.Dir(dst)

		if _, e = os.Stat(dir); e != nil {
			return e
		}

		if e = os.MkdirAll(dir, os.ModeDir); e != nil {
			return e
		}
	}

	return CopyFile(dst, src)
}
