package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type House struct {
	Name        string `json:"name"`
	Sigil       string `json:"sigil"`
	Men_at_Arms int    `json:"men"`
}

var Houses = []House{
	{Name: "Stark", Sigil: "Direwolf", Men_at_Arms: 3000},
	{Name: "Glover", Sigil: "Waves", Men_at_Arms: 450},
	{Name: "Mormont", Sigil: "Bear", Men_at_Arms: 800},
	{Name: "Bolton", Sigil: "Flayed Man", Men_at_Arms: 1200},
}

func main() {
	router := gin.Default()
	router.GET("/houses", allHouses)
	router.POST("/houses", listHouse)
	router.GET("/houses:id", getHouse)

	router.Run("localhost:3000")
	fmt.Println(("Houses are up!"))
}

func allHouses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Houses)
}

func listHouse(c *gin.Context) {
	var newHouse House
	if err := c.BindJSON(&newHouse); err != nil {
		return
	}
	Houses = append(Houses, newHouse)
	c.IndentedJSON(http.StatusCreated, newHouse)
	fmt.Println("Added")
}

func getHouse(c *gin.Context) {
	name := c.Param("name")
	for _, a := range Houses {
		if a.Name == name {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "House not found"})
}
