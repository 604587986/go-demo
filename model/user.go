package model

import (
	"github.com/gin-gonic/gin"
	. "go/database"
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key" json:"id"` 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Account string `json:"account"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

func AllUsers(c *gin.Context) (users []User,err error) {
	err = Db.Find(&users).Error
	return
}

