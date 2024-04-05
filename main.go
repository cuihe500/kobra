package main

import (
	"gitlab.eaip.top/gorm-gen-gin-learn-project/model"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/tools"
)

func main() {
	user := model.User{
		Email:    "1337985150@qq.com",
		Username: "cuihe500",
		Password: "Ch870219176!",
		Name:     "崔昌赫",
		Age:      24,
	}

	if table := tools.DB.Migrator().HasTable(&model.User{}); !table {
		tools.StdoutLog.Info("未能检测到数据库表存在,即将开始初始化数据库.")
		err := tools.DB.AutoMigrate(&model.User{})
		if err != nil {
			tools.StderrLog.Error("初始化数据库失败！", "reason", err)
		}
	}

	result := tools.DB.Create(&user)

	tools.StdoutLog.Info("成功插入插入数据!", "插入数量", result.RowsAffected)

	if result.Error != nil {
		tools.StderrLog.Error("插入数据失败！", "原因", result.Error.Error())
	}

	var user2 []model.User

	result1 := tools.DB.Find(&user2)
	tools.StdoutLog.Info("查询数据成功！", "查询结果", user2, "结果数量:", result1.RowsAffected)
}
