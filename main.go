package main

import (
	"errors"
	"net/http"
	"planetor-reborn/data"
	"strconv"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func seed(c *gin.Context) {
	
	data.FillDB()
	c.JSON(http.StatusOK, "db filled with random data")
}
func getCelestialBodies(c *gin.Context) {	
	enableCors(c)

	var celestialBodies []data.CelestialBody
	data.Db.Find(&celestialBodies)	

	c.IndentedJSON(http.StatusOK, celestialBodies)
}
func addCelestialBody(c *gin.Context) {
	var celestialBody data.CelestialBody
	if err := c.BindJSON(&celestialBody); err != nil {
		return
	}
	celestialBody.Id = 0
	err := data.Db.Create(&celestialBody).Error
	if err != nil {

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, celestialBody)
	}

}
func getCelestialBodyById(c *gin.Context) {
	id := c.Param("id")
	var celestialBody data.CelestialBody
	err := data.Db.First(&celestialBody, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found in database"})
	} else {
		c.IndentedJSON(http.StatusOK, celestialBody)
	}

}
func updateCelestialBody(c *gin.Context) {
	id := c.Param("id")
	var celestialBody data.CelestialBody
	err := data.Db.First(&celestialBody, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found in database"})
	} else {
		if err := c.BindJSON(&celestialBody); err != nil {
			return
		}
		celestialBody.Id, _ = strconv.Atoi(id)
		data.Db.Save(&celestialBody)
		c.IndentedJSON(http.StatusOK, celestialBody)
	}

}
func deleteCelestialBody(c *gin.Context) {
	id := c.Param("id")
	var celestialBody data.CelestialBody
	err := data.Db.First(&celestialBody, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found in database"})
	}
	celestialBody.Id, _ = strconv.Atoi(id)
	data.Db.Delete(&celestialBody)
	c.IndentedJSON(http.StatusOK, celestialBody)

}
func enableCors(c *gin.Context) {
	(*c).Header("Access-Control-Allow-Origin", "*")
}


func main() {
  	
	data.ConnectDatabase()

	router := gin.Default()
	

	router.GET("/api/seed", seed)
	
	router.GET("/api/celestialbodies", getCelestialBodies)
	router.GET("/api/celestialbody/:id", getCelestialBodyById)
	router.PUT("/api/celestialbody/:id", updateCelestialBody)
	router.POST("/api/celestialbody", addCelestialBody)
	router.DELETE("/api/celestialbody/:id", deleteCelestialBody)

	router.GET("/api/getpeopleinspace", GetPeopleInSpace)

	router.Run(":8080")

}
