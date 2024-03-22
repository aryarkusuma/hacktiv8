package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("aryaranggakusuma")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	router := gin.Default()

	router.POST("/login", login)

	authGroup := router.Group("/api")
	authGroup.Use(authMiddleware())
	authGroup.GET("/protected", protectedHandler)
	authGroup.GET("/getorder", GetOrder)
	authGroup.POST("/postorder", PostOrder)

	router.Run(":8080")
}

func login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if loginData.Username != "user" || loginData.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: loginData.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(
			tokenString,
			&Claims{}, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You're authorized"})
}

func PostOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You're authorized to get order"})
}

func protectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You're authorized to post order"})
}
