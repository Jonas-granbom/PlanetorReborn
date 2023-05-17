package data

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var server = "containers-us-west-184.railway.app"
var port = 7843
var user = "root"
var password = "XVkDp2jGDoH7EgAt4ZGT"
var database = "railway" //put in config file

func ConnectDatabase() {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, server, port, database)

	Db, _ = gorm.Open(mysql.Open(connString), &gorm.Config{})
	

	Db.AutoMigrate(&CelestialBody{})

}

func FillDB() {

	Db.Create(&CelestialBody{
		Id:                 1,
		Name:               "Mercury",
		Mass:               0.330,
		Density:            5427,
		Diameter:           4879,
		Gravity:            3.7,
		DayInEarthHours:    4222.6,
		YearInEarthDays:    87.9,
		Moons:              0,
		AverageTemperature: 167})
}
