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

func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		
		c.Next()

	} else {
        
		// Everytime we receive an OPTIONS request, 
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real 
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}

func main() {
  	
	data.ConnectDatabase()

	router := gin.Default()
	router.Use(CORS)

	router.GET("/api/seed", seed)
	
	router.GET("/api/celestialbody", getCelestialBodies)
	router.GET("/api/celestialbody/:id", getCelestialBodyById)
	router.PUT("/api/celestialbody/:id", updateCelestialBody)
	router.POST("/api/celestialbody", addCelestialBody)
	router.DELETE("/api/celestialbody/:id", deleteCelestialBody)

	router.GET("/api/getpeopleinspace", GetPeopleInSpace)

	router.Run(":8080")

}
