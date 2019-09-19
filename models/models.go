package models

import (
    u "go-cars/utils"
	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
	Brand string `json:"brand"`
	Vehicle string `json:"vehicle"`
	Price uint `json:"price"` 
    Status string `json:"status"`
	Mileage uint `json:"mileage"`
}

var status = []string{ "В пути", "На складе", "Продан", "Снят с продажи" }

func (model *Model) Validate() (map[string] interface{}, bool) {
     var i int
	for _, s := range status {
        if (model.Status == s){
			i = 1    
		}
	}
	if ( i==0 ) {
	return u.Message(false, "Status not correct"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (model *Model) Create() (map[string] interface{}) {


	if resp, ok := model.Validate(); 
	!ok {
		return resp
	}
	GetDB().Create(model)
	resp := u.Message(true, "success")
	resp["model"] = model
	return resp
}


func (model *Model) Delete(id string) (map[string] interface{}) {
	 
	err := GetDB().Table("models").Where("id = ?", id).First(model).Error
	if err != nil {
		resp := u.Message(false, "Model not found")
		return resp
	}
	
	GetDB().Table("models").Where("Id= ?", id).Delete(model)
	resp := u.Message(true, "success")
	return resp
}

func (model *Model) GetModels(id string) (map[string] interface{}) {
	 
	err := GetDB().Table("models").Where("id = ?", id).First(model).Error
	if err != nil {
		resp := u.Message(false, "Model not found")
		return resp
	}
	
	GetDB().Table("models").Where("Id= ?", id).First(&model)
	resp := u.Message(true, "success")
	resp["model"] = model
	return resp
}