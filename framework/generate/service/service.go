/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-07 10:58:49
 */

package service

import (
	"errors"
	"gorm.io/gorm"
	"medium-server-go/framework/db"
	"medium-server-go/framework/proto/result"
)

// 根据 id 查询
func findById(id string) *Generate {
	var generate Generate

	query := db.Mysql.Find(&generate, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &generate
}

// 根据 userId 查询
func findByUserId(userId string) *Generate {
	var generate Generate

	query := db.Mysql.Where(Generate{
		UserId: userId,
	}).Find(&generate)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &generate
}

// 根据 userId 查询所有
func findAllByUserId(userId string, page int, pageSize int) (generate []Generate, pagination result.Pagination) {
	return findAllByModel(Generate{
		UserId: userId,
	}, page, pageSize)
}

// 查询所有
func findAll(page int, pageSize int) (generate []Generate, pagination result.Pagination) {
	return findAllByModel(Generate{}, page, pageSize)
}

// 根据 Model 查询所有
func findAllByModel(where Generate, page int, pageSize int) (generate []Generate, pagination result.Pagination) {
	db.Mysql.Limit(pageSize).Offset((page - 1) * pageSize).Where(where).Find(&generate)

	var totalSize int64
	db.Mysql.Model(&Generate{}).Find(where).Count(&totalSize)

	totalPage := totalSize / int64(pageSize)
	if totalSize%int64(pageSize) != 0 {
		totalPage++
	}

	pagination = result.Pagination{
		Page:      page,
		PageSize:  len(generate),
		TotalPage: totalPage,
		TotalSize: totalSize,
	}

	return
}

// 根据 userId 获取总条数
func countAllByUserId(userId string) int64 {
	return countAllByModel(Generate{
		UserId: userId,
	})
}

// 获取总条数
func countAll() int64 {
	return countAllByModel(Generate{})
}

// 根据 Model 获取总条数
func countAllByModel(where Generate) int64 {
	var totalSize int64
	db.Mysql.Model(&Generate{}).Find(where).Count(&totalSize)
	return totalSize
}

// 根据 id 更新
func updateById(id string, generate Generate) bool {
	var one Generate

	query := db.Mysql.Find(&one, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return false
	}

	generate.Id = one.Id
	db.Mysql.Save(generate)
	return true
}
