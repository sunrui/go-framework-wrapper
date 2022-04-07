/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-07 16:12:55
 */

package template

import (
	"errors"
	"gorm.io/gorm"
	"medium-server-go/framework/db"
	"medium-server-go/framework/proto/result"
)

// FindById 根据 id 查询
func FindById(id string) *Template {
	var template Template

	query := db.Mysql.Find(&template, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &template
}

// FindByIdAndUserId 根据 id、userId 查询
func FindByIdAndUserId(id string, userId string) *Template {
	var template Template

	query := db.Mysql.Find(&template, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	if template.UserId != userId {
		return nil
	}

	return &template
}

// FindByUserId 根据 userId 查询
func FindByUserId(userId string) *Template {
	var template Template

	query := db.Mysql.Where(Template{
		UserId: userId,
	}).Find(&template)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &template
}

// FindAllByUserId 根据 userId 查询所有
func FindAllByUserId(userId string, page int, pageSize int) (template []Template, pagination result.Pagination) {
	return FindAllByModel(Template{
		UserId: userId,
	}, page, pageSize)
}

// FindAll 查询所有
func FindAll(page int, pageSize int) (template []Template, pagination result.Pagination) {
	return FindAllByModel(Template{}, page, pageSize)
}

// FindAllByModel 根据 Model 查询所有
func FindAllByModel(where Template, page int, pageSize int) (template []Template, pagination result.Pagination) {
	db.Mysql.Limit(pageSize).Offset((page - 1) * pageSize).Where(where).Find(&template)

	var totalSize int64
	db.Mysql.Model(&Template{}).Find(where).Count(&totalSize)

	totalPage := totalSize / int64(pageSize)
	if totalSize%int64(pageSize) != 0 {
		totalPage++
	}

	pagination = result.Pagination{
		Page:      page,
		PageSize:  len(template),
		TotalPage: totalPage,
		TotalSize: totalSize,
	}

	return
}

// CountAllByUserId 根据 userId 获取总条数
func CountAllByUserId(userId string) int64 {
	return CountAllByModel(Template{
		UserId: userId,
	})
}

// CountAll 获取总条数
func CountAll() int64 {
	return CountAllByModel(Template{})
}

// CountAllByModel 根据 Model 获取总条数
func CountAllByModel(where Template) int64 {
	var totalSize int64
	db.Mysql.Model(&Template{}).Find(where).Count(&totalSize)
	return totalSize
}

// UpdateById 根据 id 更新
func UpdateById(id string, template Template) bool {
	var one Template

	query := db.Mysql.Find(&one, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return false
	}

	template.Id = one.Id
	db.Mysql.Save(template)
	return true
}

// UpdateByIdAndUserId 根据 id、userId 更新
func UpdateByIdAndUserId(id string, userId string, template Template) bool {
	var one Template

	query := db.Mysql.Find(&one, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return false
	}

	if one.UserId != userId {
		return false
	}

	template.Id = one.Id
	db.Mysql.Save(template)
	return true
}

// UpdateById 根据 id 删除
func DeleteById(id string) bool {
	var one Template

	query := db.Mysql.Find(&one, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return false
	}

	db.Mysql.Delete(query)
	return true
}

// DeleteByIdAndUserId 根据 id 删除
func DeleteByIdAndUserId(id string, userId string) bool {
	var one Template

	query := db.Mysql.Find(&one, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return false
	}

	if one.UserId != userId {
		return false
	}

	db.Mysql.Delete(query)
	return true
}
