package dao

import (
	"fmt"
	"testing"
)

func Test_GetByUsername(t *testing.T) {
	db, err := getGormDB()
	if err != nil {
		t.Fatal(err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer sqlDB.Close()

	dao := NewUserDao(db)

	username := "chenhao6"
	user, err := dao.GetByUsername(username)
	if err != nil {
		t.Fatalf("获取用户失败, username: %v. %v", username, err.Error())
	}

	fmt.Println(user)
}
