package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
)

var db *gorm.DB
var app *iris.Application

const irisAddr string = ":8080"

func main() {
	defer db.Close()
	initDB()
	initIris()

	fmt.Println("ok")
}

func initDB() {
	db, _ = gorm.Open("mysql", "root:qnswjsqks@tcp(192.168.0.19:3306)/samp")
	db.AutoMigrate(&Permission{}, &ClubMember{}, &Comment{}, &Session{}, &Reservation{}, &Asset{}, &User{}, &Category{}, &Club{})
	db.Model(&ClubMember{}).AddForeignKey("club_id", "clubs(id)", "CASCADE", "CASCADE")
	db.Model(&ClubMember{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&ClubMember{}).AddForeignKey("permission_id", "permissions(id)", "CASCADE", "CASCADE")
	db.Model(&Comment{}).AddForeignKey("asset_id", "assets(id)", "CASCADE", "CASCADE")
	db.Model(&Comment{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Comment{}).AddForeignKey("parent_comment_id", "comments(id)", "CASCADE", "CASCADE")
	db.Model(&Session{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Reservation{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Reservation{}).AddForeignKey("asset_id", "assets(id)", "CASCADE", "CASCADE")
	db.Model(&Asset{}).AddForeignKey("club_id", "clubs(id)", "CASCADE", "CASCADE")
	db.Model(&Asset{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	db.Model(&Category{}).AddForeignKey("club_id", "clubs(id)", "CASCADE", "CASCADE")
}

func initIris() {
	app = iris.Default()
	app.Delete("/v1/account", memberUnregister)
	app.Run(iris.Addr(irisAddr))
}
