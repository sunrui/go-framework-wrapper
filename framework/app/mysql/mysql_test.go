/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-05 22:24:45
 */

package mysql

import (
	"encoding/json"
	"framework/app/glog"
	"gorm.io/gorm"
	"testing"
	"time"
)

type Basic struct {
	Name string `json:"name"  gorm:"primaryKey;unique;type:varchar(32); comment:名称"`
}

type User struct {
	Model
	Basic `gorm:"embedded"`
	Age   int    `json:"age"  gorm:"default:18; comment:年龄"` // 年龄
	Class string `json:"class" gorm:"type:varchar(32); not null;comment:班级"`
}

func (User) TableName() string {
	return "_t_user"
}

type UserScore struct {
	Model
	Basic `gorm:"embedded"`
	Score int `json:"score" gorm:"not null;check:score>=0&&score<=100;comment:分数"` // 分数

	User   *User  `json:"user,omitempty" gorm:"foreignKey:UserId"`
	UserId string `json:"userId" gorm:"comment:用户 id"`
}

func (UserScore) TableName() string {
	return "_t_user_score"
}

var log *glog.GLog
var db *Mysql

// TestMain 初始化前准备
func TestMain(m *testing.M) {
	var err error

	// 启动控制台日志
	log = glog.NewGLog(glog.DefaultLayout{}, []glog.Appender{
		glog.ConsoleAppender{},
	})

	// 测试数据库连接
	db, err = New(Config{
		User:         "root",
		Password:     "honeysenselt",
		Host:         "127.0.0.1",
		Port:         3306,
		Database:     "medium_dev",
		MaxOpenConns: 1,
		MaxIdleConns: 1,
	}, log)
	if err != nil {
		panic(err.Error())
	}

	// 测试多个数据库初始化
	db.AutoMigrate(User{}, UserScore{})

	// 删除两个数据库
	db.Where("1 = 1").Delete(&User{})
	db.Where("1 = 1").Delete(&UserScore{})
	db.Exec("DELETE FROM _t_user")
	db.Exec("DELETE FROM _t_user_score")
	db.Exec("OPTIMIZE TABLE _t_user")
	db.Exec("OPTIMIZE TABLE _t_user_score")
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&UserScore{})

	// 强制删除数据库
	db.Unscoped().Delete(&User{})
	db.Unscoped().Delete(&UserScore{})

	m.Run()
}

func TestMysql_Insert(t *testing.T) {
	tm := time.Now().Format("2006-01-02 15:04:05")

	user := User{
		Basic: Basic{
			Name: "张三",
		},
		Age: 19,
	}

	userScoreRepository := NewRepository[UserScore](db)
	if user := userScoreRepository.FindOne(&User{
		Basic: Basic{
			Name: "张三",
		},
	}); user == nil {
		t.Fatal("have this id")
	} else {
		userJson, _ := json.Marshal(user)
		t.Log("\n" + string(userJson) + "\n")
	}

	db.Save(&user)

	// 测试 beyond to
	score := UserScore{
		Basic: Basic{
			Name: "语文-" + tm,
		},
		UserId: user.Id,
		Score:  80,
	}
	db.Save(&score)
}

func TestMysql_Find(t *testing.T) {
	userScoreRepository := NewRepository[UserScore](db)
	if user := userScoreRepository.FindOne(&User{
		Basic: Basic{
			Name: "张三",
		},
	}); user == nil {
		t.Fatal("have this id")
	} else {
		userJson, _ := json.Marshal(user)
		t.Log("\n" + string(userJson) + "\n")
	}
}
