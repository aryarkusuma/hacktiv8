package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type ReservSeat struct {
	ShuttleId  string `json:"shuttle_id" binding:"required"`
	ReservName string `json:"name" binding:"required"`
	SeatNumber int    `json:"seat_number" binding:"required"`
	UserId     string `json:"user_id" binding:"required"`
}

type UserReservSeatList struct {
	ShuttleId  string `json:"shuttle_id"`
	ReservName string `json:"name"`
	SeatNumber int    `json:"seat_number" `
	UserId     string `json:"user_id" `
	CreatedAt  string `json:"created_at"`
}

type Db struct {
	DB *pgx.Conn
}

func (DB *Db) ReservSeatGin(c *gin.Context) {

	var v ReservSeat
	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check seat availability -> db required
	var count int
	err := DB.DB.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM reservations WHERE shuttle_id = $1 AND seat_number = $2",
		v.ShuttleId, v.SeatNumber).Scan(&count)

	if err != nil {
		fmt.Println((v))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Request"})
		return
	}

	//reserved available seat -> db required
	_, err = DB.DB.Exec(context.Background(),
		"INSERT INTO reservations (shuttle_id, reserv_name, seat_number, user_id) VALUES ($1, $2, $3, $4)",
		v.ShuttleId, v.ReservName, v.ReservName, v.SeatNumber, v.UserId)

	if err != nil {
		fmt.Println((v))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Request"})
		return
	}

	//

	c.JSON(http.StatusAccepted, gin.H{"count": count})

}
