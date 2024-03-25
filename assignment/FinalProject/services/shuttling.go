package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PostShuttle struct {
	ShuttleType string `json:"shuttle_type"  binding:"required"`
	Seats       int    `json:"seats"  binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	RouteStart  string `json:"route_start"  binding:"required"`
	RouteEnd    string `json:"route_end"  binding:"required"`
}

type Shuttle struct {
	ID          string    `json:"id"  binding:"required"`
	ShuttleType string    `json:"shuttle_type"  binding:"required"`
	Seats       int       `json:"seats"  binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	RouteStart  string    `json:"route_start"  binding:"required"`
	RouteEnd    string    `json:"route_end"  binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

func (DB *Db) PostShuttleGin(c *gin.Context) {

	var v PostShuttle
	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "req error"})

	}

	//check seat availability -> db required

	//reserved available seat -> db required

	//

	c.Status(http.StatusCreated)

	//post shuttle to reserv -> db required

}

func (DB *Db) GetShuttlesGin(c *gin.Context) {

	rows, err := DB.DB.Query(context.Background(), "SELECT id, shuttle_type, seats, start_date, route_start, route_end, created_At FROM shuttles")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Slice to hold the query results
	var shuttles []Shuttle

	// Iterate over the rows and scan data into struct
	for rows.Next() {
		var shuttle Shuttle
		if err := rows.Scan(&shuttle.ID, &shuttle.ShuttleType, &shuttle.Seats, &shuttle.StartDate, &shuttle.RouteStart, &shuttle.RouteEnd, &shuttle.CreatedAt); err != nil {
			fmt.Println("Error scanning row data:", err)
			return
		}
		// Append the scanned row data to shuttles slice
		shuttles = append(shuttles, shuttle)
	}
	// Check for errors during rows iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return
	}

	// Marshal shuttles into JSON
	// jsonData, err := json.Marshal(shuttles)
	// if err != nil {
	// 	fmt.Println("Error marshaling JSON:", err)
	// 	return
	// }

	// Print JSON data

	c.SecureJSON(http.StatusOK, gin.H{"data": shuttles})
}
