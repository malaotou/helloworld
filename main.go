package main

import (
	"encoding/xml"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		c.XML(200, Person{FirstName: "Mohamed", LastName: "name" + name})
	})
	router.Run(":3000")
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

type Recipe struct {
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredents`
	Instructions []string  `json:"instructions"`
	PulishedAt   time.Time `json:"published"`
}
