package model

import (
	"github.com/ve-weiyi/go-examplse/src/core/db"
	"github.com/ve-weiyi/go-examplse/src/utils/jsonconv"
	"log"
	"testing"
)

const dsn = "root:mysql7914@(127.0.0.1:3306)/anker?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	db.Connect(dsn)
}

func TestCreate(t *testing.T) {
	model := getModel()
	if model.IsValid() {
		db.DB().Where(model).Create(model)
	}
}

func TestDelete(t *testing.T) {
	model := getModel()
	if model.IsValid() {
		db.DB().Where(model).Delete(model)
	}
}

func TestUpdate(t *testing.T) {
	model := getModel()
	if model.IsValid() {
		db.DB().Where("id = ?", model.ID).Updates(model)
	}
}

func TestList(t *testing.T) {
	var list []User
	db.DB().Where("1 = 1").Find(&list)
	log.Println(jsonconv.ObjectToJson(list))
}

func getModel() *User {
	return &User{
		ID:       0,
		Username: "test",
		Password: "123456",
	}
}