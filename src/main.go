package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db, _ := gorm.Open("mysql", "root:qnswjsqks@tcp(192.168.0.19:3306)/samp")
	defer db.Close()
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

	fmt.Println("asdf")
}
