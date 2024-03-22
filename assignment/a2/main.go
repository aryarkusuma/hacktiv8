package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Teacher struct {
	ID        string `json:"id"`
	OrderedAt string `json:"orderedAt"`
}

func main() {
	r := gin.Default()
	// define the routes
	r.GET("/order", GetOrder)
	r.POST("/order", PostOrder)
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}
}

func GetOrder(c *gin.Context) {
	// retrieve teacher from db
	id := c.Query("id")
	order, err := GetOrderId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "impossible to retrieve teacher"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func PostOrder(c *gin.Context) {
	// retrieve teacher from db
	id := c.Query("id")
	order, err := GetOrderId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "impossible to retrieve teacher"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func GetOrderId(id string) (Teacher, error) {
	// TODO : lookup in db
	t := time.Now()
	trfc := (t.Format(time.RFC3339))

	return Teacher{
		ID:        id,
		OrderedAt: trfc,
	}, nil
}
