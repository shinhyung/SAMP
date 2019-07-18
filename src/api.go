package main

import "github.com/kataras/iris"

func userPost(ctx iris.Context) {
	db :=openDB()
	ctx.para
	ctx.JSON(iris.Map({"success":true}))
}

func userDelete(ctx iris.Context) {

}

func clubGet(ctx iris.Context) {

}

func clubPut(ctx iris.Context) {

}

func clubPost(ctx iris.Context) {

}

func clubDelete(ctx iris.Context) {

}

func memberDelete(ctx iris.Context) {

}

func memberGet(ctx iris.Context) {

}

func memberGetList(ctx iris.Context) {

}

func memberPutPermission(ctx iris.Context) {

}

func categoryGet(ctx iris.Context) {

}

func categoryGetList(ctx iris.Context) {

}

func categoryPut(ctx iris.Context) {

}

func categoryPost(ctx iris.Context) {

}

func categoryDelete(ctx iris.Context) {

}

func assetPut(ctx iris.Context) {

}

func assetDelete(ctx iris.Context) {

}

func assetGet(ctx iris.Context) {

}

func assetGetList(ctx iris.Context) {

}

func assetPost(ctx iris.Context) {

}

func login(ctx iris.Context) {

}

func logout(ctx iris.Context) {

}

func waiterGetList(ctx iris.Context) {

}

func waiterAccept(ctx iris.Context) {

}

func waiterDelete(ctx iris.Context) {

}

func commentGetList(ctx iris.Context) {

}

func commentPost(ctx iris.Context) {

}

func commentDelete(ctx iris.Context) {

}

func commentOnCommentPost(ctx iris.Context) {

}

func ownGetList(ctx iris.Context) {

}

func ownPost(ctx iris.Context) {

}

func ownPut(ctx iris.Context) {

}

func ownDelete(ctx iris.Context) {

}

func likeGetList(ctx iris.Context) {

}

func likePost(ctx iris.Context) {

}

func likeDelete(ctx iris.Context) {

}
