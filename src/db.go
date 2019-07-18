package main

import (
	"bufio"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

var id string
var host string
var pw string

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func initDB(){
	f, err := os.Open("src/account.txt")
	check(err)
	defer f.Close()
	scanner:= bufio.NewScanner(f)
	scanner.Scan()
	host=scanner.Text()
	scanner.Scan()
	id=scanner.Text()
	scanner.Scan()
	pw=scanner.Text()
}

func openDB() *gorm.DB{
	r, err:=gorm.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s)/SAMP",id,pw,host))
	check(err)
	return r
}

func resetDB() {
	db := openDB()
	defer db.Close()
	db.DropTableIfExists(&Permission{}, &ClubMember{}, &Comment{}, &Session{}, &Reservation{}, &Asset{}, &User{}, &Category{}, &Club{})
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

//Permission .
type Permission struct {
	gorm.Model
	King             bool
	MemberAccept     bool
	MemberDelete     bool
	MemberPermission bool
	CategoryAdd      bool
	CategoryDelete   bool
	AssetAdd         bool
	AssetDelete      bool
}

//TableName .
func (Permission) TableName() string {
	return "permissions"
}

//Club .
type Club struct {
	gorm.Model
	Name        string
	Description string
}

//TableName .
func (Club) TableName() string {
	return "clubs"
}

//User .
type User struct {
	gorm.Model
	Name   string
	Email  string
	Passwd string
	Phone  string
}

//TableName .
func (User) TableName() string {
	return "users"
}

//ClubMember .
type ClubMember struct {
	gorm.Model
	ClubID       uint
	UserID       uint
	IsAccepted   bool
	PermissionID uint
}

//TableName .
func (ClubMember) TableName() string {
	return "club_members"
}

//Comment .
type Comment struct {
	gorm.Model
	AssetID         uint
	UserID          uint
	ParentCommentID uint
	WrittenTime     time.Time
	Content         string
}

//TableName .
func (Comment) TableName() string {
	return "comments"
}

//Session .
type Session struct {
	gorm.Model
	LoginTime time.Time
	Mac       string
	UserID    uint
	Key       string
}

//TableName .
func (Session) TableName() string {
	return "sessions"
}

//Reservation .
type Reservation struct {
	gorm.Model
	AssetID         uint
	UserID          uint
	ReservationTime time.Time
	ReturnTime      time.Time
	IsReturned      bool
}

//TableName .
func (Reservation) TableName() string {
	return "reservations"
}

//Asset .
type Asset struct {
	gorm.Model
	ClubID        uint
	CategoryID    uint
	Name          string
	Description   string
	ImageFileName string
}

//TableName .
func (Asset) TableName() string {
	return "assets"
}

//Category .
type Category struct {
	gorm.Model
	ClubID           uint
	ParentCategoryID uint
	Name             string
}

//TableName .
func (Category) TableName() string {
	return "categories"
}
