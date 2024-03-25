package main

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/aryarkusuma/hacktiv8/assignment/FinalProject/repository"
	"github.com/aryarkusuma/hacktiv8/assignment/FinalProject/services"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db := repository.DbInit(context.Background())

	s := &services.Db{DB: db}

	router := gin.Default()

	router.POST("/login", s.Login)

	router.POST("/reg", s.Register)

	authGroup := router.Group("/api")

	authGroup.Use(services.AuthMiddleware())

	authGroup.POST("/reserveseat", s.ReservSeatGin)

	authGroup.POST("/postshuttle", s.PostShuttleGin)

	authGroup.GET("/shuttlelist", s.GetShuttlesGin)

	router.Run(":8080")

	db.Close(context.Background())
}
