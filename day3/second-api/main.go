package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/second-api/Config"
	"github.com/second-api/Models"
	"github.com/second-api/Routes"
)
var err error
func main() {
	//fmt.Println(time.)
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)

	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Student_table{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}