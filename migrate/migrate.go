package main

import (
	"fmt"
	"my-favorite-pokemon-rest-api/db"
	"my-favorite-pokemon-rest-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	// 例：userとtaskテーブルを作成したい場合
	dbConn.AutoMigrate(&model.User{}, &model.Star{}) //作成したいモデルのstructを0値で引数に渡す
}
