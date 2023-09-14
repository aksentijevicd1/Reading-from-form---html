package models

import (
	"github.com/aksentijevicd1/reading-from-form-go/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Opinion struct {
	gorm.Model
	FirstName string `gorm:"column:firstname" json:"firstname"`
	LastName  string `gorm:"column:lastname" json:"lastname"`
	Address   string `gorm:"column:address" json:"address"`
	Opinion   string `gorm:"column:opinion" json:"opinion"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Opinion{})
}

func (o *Opinion) AddOpinion() {
	db.Create(&o)
}

func GetOpinions() []Opinion {
	var allOpinions []Opinion
	db.Find(&allOpinions)
	return allOpinions
}
