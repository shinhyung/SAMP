package main

import (
	"fmt"

	
	"github.com/kataras/iris"
)

var app *iris.Application

const irisAddr string = ":8080"

func main() {
	initDB()
	initIris()

	fmt.Println("ok")
}

func initIris() {
	
	app = iris.Default()
	app.Delete("/v1/account", userDelete) //회원 탈퇴
	app.Post("/v1/account", userPost)   //회원 가입
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
	app.Delete("/v1/club/{clubID:uint}/categories/{categoryID:uint}", categoryDelete)
	app.Get("/v1/club/{clubID:uint}/assets", assetGetList)
	app.Post("/v1/club/{clubID:uint}/assets", assetPost)
	app.Get("/v1/club/{clubID:uint}/assets/{assetID:uint}", assetGet)
	app.Delete("/v1/club/{clubID:uint}/assets/{assetID:uint}", assetDelete)
	app.Put("/v1/club/{clubID:uint}/assets/{assetID:uint}", assetPut)
	app.Get("/v1/club/{clubID:uint}/assets/{assetID:uint}/comments", commentGetList)
	app.Post("/v1/club/{clubID:uint}/assets/{assetID:uint}/comments", commentPost)
	app.Delete("/v1/club/{clubID:uint}/assets/{assetID:uint}/comments/{commentID:uint}", commentDelete)
	app.Post("/v1/club/{clubID:uint}/assets/{assetID:uint}/comments/{commentID:uint}", commentOnCommentPost)

	_ = app.Run(iris.Addr(irisAddr))
}

//app.Get("/profile/{name:alphabetical max(255)}",
