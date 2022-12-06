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
)

type Basic struct {
	Name string `json:"name"  gorm:"primaryKey;unique;type:varchar(6); comment:名称"`
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

func TestMysql(t *testing.T) {
	// 启动控制台日志
	log := glog.NewGLog(glog.DefaultLayout{}, []glog.Appender{
		glog.ConsoleAppender{},
	})

	// 测试数据库连接
	db, err := New(Config{
		User:         "root",
		Password:     "honeysenselt",
		Host:         "127.0.0.1",
		Port:         3306,
		Database:     "medium_dev",
		MaxOpenConns: 1,
		MaxIdleConns: 1,
	}, log)
	if err != nil {
		t.Fatal(err.Error())
	}

	// 测试多个数据库初始化
	db.AutoMigrate(User{}, UserScore{})

	// 软删除两个数据库
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

	// 测试 beyond to
	user := UserScore{
		Basic: Basic{
			Name: "语文",
		},
		User: &User{
			Basic: Basic{
				Name: "张三",
			},
			Age: 19,
		},
		Score: 80,
	}
	db.Save(&user)

	// 测试外键
	score := UserScore{
		Basic: Basic{
			Name: "数学",
		},
		Score:  100,
		UserId: user.User.Id,
	}
	db.Save(&score)

	userScoreRepository := NewRepository[UserScore](db)
	if user := userScoreRepository.FindById(score.Id); user == nil {
		t.Fatal("have this id")
	} else {
		userJson, _ := json.Marshal(user)
		t.Log("\n" + string(userJson) + "\n")
	}
}
