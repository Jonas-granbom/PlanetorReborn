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
	//data.FillDB()
	c.JSON(http.StatusOK, gin.H{"message":"DB filled with dummydata"})
}
func GetCelestialBodies(c *gin.Context) {	
	
	var celestialBodies []data.CelestialBody
	data.Db.Find(&celestialBodies)	

	c.JSON(http.StatusOK, celestialBodies)
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

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
func SetupRouter() *gin.Engine{
	router := gin.Default()
	router.Use(CORS())

	router.GET("/seed", seed)
	
	router.GET("/celestialbody", GetCelestialBodies)
	router.GET("/celestialbody/:id", getCelestialBodyById)
	router.PUT("/celestialbody/:id", updateCelestialBody)
	router.POST("/celestialbody", addCelestialBody)
	router.DELETE("/celestialbody/:id", deleteCelestialBody)

	router.GET("/getpeopleinspace", GetPeopleInSpace)
	return router
}

func main() {
  	
	data.ConnectDatabase()

	router := SetupRouter()


	router.Run(":8080")

}
