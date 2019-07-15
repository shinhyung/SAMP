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
	app.Delete("/v1/account", memberDelete) //회원 탈퇴
	app.Post("/v1/account", memberDelete)   //회원 가입
	app.Post("/v1/login", login)
	app.Get("/v1/logout", logout)
	app.Post("/v1/club", clubPost)
	app.Get("/v1/club", clubGet)
	app.Put("/v1/club/{clubID:uint}", clubPut)
	app.Delete("/v1/club/{clubID:uint}", clubDelete)
	app.Get("/v1/club/{clubID:uint}", waiterGetList)
	app.Put("/v1/club/{clubID:uint}/waiters/{userID:uint}", waiterAccept)
	app.Delete("/v1/club/{clubID:uint}/waiters/{userID:uint}", waiterDelete)
	app.Get("/v1/club/{clubID:uint}/members", memberGetList)
	app.Get("/v1/club/{clubID:uint}/members/{userID:uint}", memberGet)
	app.Delete("/v1/club/{clubID:uint}/members/{userID:uint}", memberDelete)
	app.Put("/v1/club/{clubID:uint}/members/{userID:uint}", memberPutPermission)
	app.Get("/v1/club/{clubID:uint}/members/{userID:uint}/owns", ownGetList)
	app.Post("/v1/club/{clubID:uint}/members/{userID:uint}/owns", ownPost)
	app.Put("/v1/club/{clubID:uint}/members/{userID:uint}/owns/{assetID:uint}", ownPut)
	app.Delete("/v1/club/{clubID:uint}/members/{userID:uint}/owns/{assetID:uint}", ownDelete)
	app.Get("/v1/club/{clubID:uint}/members/{userID:uint}/likes", likeGetList)
	app.Post("/v1/club/{clubID:uint}/members/{userID:uint}/likes", likePost)
	app.Delete("/v1/club/{clubID:uint}/members/{userID:uint}/likes/{assetID:uint}", likeDelete)
	app.Get("/v1/club/{clubID:uint}/categories", categoryGetList)
	app.Post("/v1/club/{clubID:uint}/categories", categoryPost)
	app.Get("/v1/club/{clubID:uint}/categories/{categoryID:uint}", categoryGet)
	app.Put("/v1/club/{clubID:uint}/categories/{categoryID:uint}", categoryPut)
	app.Get("/v1/club/{clubID:uint}/assets", assetGetList)
	app.Post("/v1/club/{clubID:uint}/assets", assetPost)
	app.Get("/v1/club/{clubID:uint}/assets/{assetID:uint}", assetGet)
	app.Delete("/v1/club/{clubID:uint}/assets/{assetID:uint}", assetDelete)
	app.Put("/v1/club/{clubID:uint}/assets/{assetID:uint}", assetPut)
	app.Get("/v1/club/{clubID:uint}/assets/{assetID:uint}/comments", commentGetList)
	app.Post("/v1/club/{clubID:uint}/assets/{assetID:uint}/comments", commentPost)
	app.Delete("/v1/club/{clubID:uint}/assets/{assetID:uint}/comments/{commentID:uint}", commentDelete)
	app.Post("/v1/club/{clubID:uint}/assets/{assetID:uint}/comments/{commentID:uint}", commentOnCommentPost)

	app.Run(iris.Addr(irisAddr))
}

//app.Get("/profile/{name:alphabetical max(255)}",
