package main

import (
	"encoding/json"
	"fmt"

	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type peopleInSpace struct {
	Message string `json:"message"`
	Number  int    `json:"number"`
	People  []struct {
		Craft string `json:"craft"`
		Name  string `json:"name"`
	} `json:"people"`
}

func GetPeopleInSpace(c *gin.Context) {
	response, err := http.Get("http://api.open-notify.org/astros.json")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	var res peopleInSpace

	json.NewDecoder(response.Body).Decode(&res)
	c.IndentedJSON(http.StatusOK, res.People)
	

}
