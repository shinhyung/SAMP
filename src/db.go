package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

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
