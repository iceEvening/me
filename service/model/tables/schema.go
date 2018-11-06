package tables

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Schemas is a struct which contains all tables
type Schemas struct {
	User      *User
	Role      *Role
	Job       *Job
	Education *Education
	Article   *Article
}

//User table
type User struct {
	gorm.Model
	Email        string `gorm:"type:varchar(64);not null;unique"`
	Nickname     string `gorm:"type:varchar(16);not null;unique"`
	Passwd       string `gorm:"type:varchar(32);not null"`
	Name         string `gorm:"type:varchar(16);not null;unique"`
	Gender       uint   `gorm:"default:0"`
	ContactEmail string `gorm:"type:varchar(64)"`
	Birthday     time.Time
	Hometown     string `gorm:"type:varchar(32)"`
	City         string `gorm:"type:varchar(32)"`
	Img          []byte `gorm:"type:bytea"`
	Me           string `gorm:"type:text"`
	UserRole     uint   `gorm:"not null"`
}

//Role table
type Role struct {
	gorm.Model
	RoleName string `gorm:"type:varchar(16);not null"`
}

//Job table
type Job struct {
	gorm.Model
	UserID     uint      `gorm:"not null;"`
	Company    string    `gorm:"type:varchar(32);not null"`
	Department string    `gorm:"type:varchar(32);not null"`
	Post       string    `gorm:"type:varchar(32);not null"`
	JobLevel   string    `gorm:"type:varchar(16);not null"`
	JobDesc    string    `gorm:"type:text;not null"`
	StartTime  time.Time `gorm:"not null"`
	EndTime    time.Time `gorm:"not null"`
	IsCurrent  bool      `gorm:"not null"`
}

//Education table
type Education struct {
	gorm.Model
	UserID    uint      `gorm:"not null;"`
	School    string    `gorm:"type:varchar(32);not null"`
	Major     string    `gorm:"type:varchar(32);not null"`
	Degree    string    `gorm:"type:varchar(8);not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
	IsCurrent bool      `gorm:"not null"`
}

//Article table
type Article struct {
	gorm.Model
	UserID   uint `gorm:"not null;"`
	Category uint
	NickName string `gorm:"type:varchar(16);not null"`
	Title    string `gorm:"type:varchar(64);not null"`
	Markdown string `gorm:"type:text;not null"`
	HTML     string `gorm:"type:text;not null"`
}
