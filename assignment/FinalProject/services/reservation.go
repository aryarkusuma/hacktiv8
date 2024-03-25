package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type ReservSeat struct {
	ShuttleId  string `json:"shuttle_id" binding:"required"`
	ReservName string `json:"name" binding:"required"`
	SeatNumber int    `json:"seat_number" binding:"required"`
	// UserId     string `json:"user_id" binding:"required"` -> Savin n Using Claims.UserId in Context instead From JWT
}

type UserReservedSeats struct {
	Id         string    `json:"id"`
	ShuttleId  string    `json:"shuttle_id"`
	ReservName string    `json:"name"`
	SeatNumber int       `json:"seat_number" `
	UserId     string    `json:"user_id" `
	CreatedAt  time.Time `json:"created_at"`
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

	var shuttleStart time.Time
	err := DB.DB.QueryRow(context.Background(),
		"SELECT start_date FROM shuttles WHERE id = $1",
		v.ShuttleId).Scan(&shuttleStart)

	fmt.Println(shuttleStart.Format(time.RFC3339), time.Now())

	if err != nil || shuttleStart.Before(time.Now()) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Date Expired"})
		return
	}
	//check seat availability -> db required
	var count int
	err = DB.DB.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM reservations WHERE shuttle_id = $1 AND seat_number = $2",
		v.ShuttleId, v.SeatNumber).Scan(&count)

	if err != nil || count >= 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": count})
		return
	}

	//reserved available seat -> db required
	_, err = DB.DB.Exec(context.Background(),
		"INSERT INTO reservations (shuttle_id, reserv_name, seat_number, user_id) VALUES ($1, $2, $3, $4)",
		v.ShuttleId, v.ReservName, v.SeatNumber, c.GetString("Id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Exec"})
		return
	}
	//

	c.JSON(http.StatusAccepted, gin.H{"status": "Reservation Accepted"})

}

func (DB *Db) GetReservedSeats(c *gin.Context) {

	rows, err := DB.DB.Query(context.Background(), "SELECT id, shuttle_id, reserv_name, seat_number, user_id, created_at FROM reservations where user_id = $1", c.GetString("Id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Slice to hold the query results
	var reservedSeats []UserReservedSeats

	// Iterate over the rows and scan data into struct
	for rows.Next() {
		var reservedSeat UserReservedSeats
		if err := rows.Scan(&reservedSeat.Id, &reservedSeat.ShuttleId, &reservedSeat.ReservName, &reservedSeat.SeatNumber, &reservedSeat.UserId, &reservedSeat.CreatedAt); err != nil {
			fmt.Println("Error scanning row data:", err)
			return
		}
		// Append the scanned row data to shuttles slice
		reservedSeats = append(reservedSeats, reservedSeat)
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

	c.SecureJSON(http.StatusOK, gin.H{"data": reservedSeats})
}
