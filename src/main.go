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
	db.Model(&ClubMember{}).AddForeignKey("club_id", "club(id)", "CASCADE", "CASCADE")
	fmt.Println("asdf")
}
