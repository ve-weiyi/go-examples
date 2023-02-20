package model

import (
	"github.com/ve-weiyi/go-examplse/src/core/database"
	"github.com/ve-weiyi/go-examplse/src/utils/jsonconv"
	"log"
	"testing"
)

func init() {
	database.Connect()
}

func TestCreate(t *testing.T) {
	model := getModel()
	if model.IsValid() {
		database.DB().Where(model).Create(model)
	}
}

func TestDelete(t *testing.T) {
	model := getModel()
	if model.IsValid() {
		database.DB().Where(model).Delete(model)
	}
}

func TestUpdate(t *testing.T) {
	model := getModel()
	if model.IsValid() {
		database.DB().Where("id = ?", model.ID).Updates(model)
	}
}

func TestList(t *testing.T) {
	var list []User
	database.DB().Where("1 = 1").Find(&list)
	log.Println(jsonconv.ObjectToJson(list))
}

func getModel() *User {
	return &User{
		ID:       0,
		Username: "test",
		Password: "123456",
	}
}
