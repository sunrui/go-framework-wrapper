/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-05 22:24:45
 */

package mysql

import (
	"encoding/json"
	"fmt"
	"framework/app/glog"
	"testing"
)

type Basic struct {
	Name string `json:"name"  gorm:"primaryKey;unique;type:varchar(256); comment:名称"`
}

type User struct {
	Model
	Basic `gorm:"embedded"`
	Age   int    `json:"age"  gorm:"default:18; comment:年龄"` // 年龄
	Class string `json:"class" gorm:"type:varchar(32); not null;comment:班级"`
}

func (User) TableName() string {
	return "t_user"
}

type UserScore struct {
	Model
	Basic `gorm:"embedded"`
	Score int `json:"score" gorm:"not null;check:score>=0&&score<=100;comment:分数"` // 分数

	User   *User  `json:"user,omitempty" gorm:"foreignKey:UserId"`
	UserId string `json:"userId" gorm:"comment:用户 id"`
}

func (UserScore) TableName() string {
	return "t_user_score"
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

	// 删除数据库
	//db.Exec("DELETE FROM t_user")

	m.Run()
}

func TestMysql_Insert(t *testing.T) {
	user := User{
		Basic: Basic{
			Name: "张三",
		},
		Age: 19,
	}

	userRepository := NewRepository[User](db)
	if u := userRepository.FindOne("name = ? And age = ?", "张三", 19); u == nil {
		db.Save(&user)
	} else {
		user = *u
	}

	userScoreRepository := NewRepository[UserScore](db)
	count := userScoreRepository.Count()
	for i := count + 1; i < count+1+10; i++ {
		// 测试 beyond to
		score := UserScore{
			Basic: Basic{
				Name: fmt.Sprintf("语文 - %03d", +i),
			},
			UserId: user.Id,
			Score:  80,
		}
		db.Save(&score)
	}

	t.Log("ok")
}

func TestMysql_Find(t *testing.T) {
	userRepository := NewRepository[User](db)
	if user := userRepository.FindOne(&User{
		Basic: Basic{
			Name: "张三",
		},
	}); user == nil {
		t.Error("not have this id")
	} else {
		userJson, _ := json.Marshal(user)
		t.Log("\n" + string(userJson) + "\n")
	}
}

func TestMysql_Page(t *testing.T) {
	var userId string

	userRepository := NewRepository[User](db)
	if user := userRepository.FindOne(&User{
		Basic: Basic{
			Name: "张三",
		},
	}); user == nil {
		t.Error("not have this id")
		return
	} else {
		userId = user.Id
	}

	userScoreRepository := NewRepository[UserScore](db)

	userScorePage := userScoreRepository.FindPage(Page{
		Page:     1,
		PageSize: 10,
	}, "name ASC", &UserScore{
		UserId: userId,
	})

	userScoreJson, _ := json.Marshal(userScorePage)
	t.Log("\n" + string(userScoreJson) + "\n")

	for _, userScore := range userScorePage {
		var r bool
		//r = userScoreRepository.SoftDeleteById(userScore.Id)
		//t.Log("\n"+"SoftDeleteById userScore by id "+userScore.Id+", result =", r)
		r = userScoreRepository.DeleteById(userScore.Id)
		t.Log("\n"+"DeleteById userScore by id "+userScore.Id+", result =", r)
	}
}
