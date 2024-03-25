package services

import (
	"context"
	"net/http"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("aryaranggakusuma")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (DB *Db) Login(c *gin.Context) {

	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if !IsValid(loginData.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Data Invalid"})
		return
	}

	hashedpw, err := HashPassword(loginData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Request"})
		return
	}

	var password string
	err = DB.DB.QueryRow(context.Background(),
		"SELECT password FROM users WHERE username = $1",
		loginData.Username).Scan(&password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": hashedpw})
		return
	}

	checkhash := CheckPasswordHash(loginData.Password, password)
	if !checkhash {
		c.JSON(http.StatusInternalServerError, gin.H{"error": hashedpw})
		return
	}
	expirationTime := time.Now().Add(24 * time.Hour)
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

	c.JSON(http.StatusAccepted, gin.H{"token": tokenString})
}

func (DB *Db) Register(c *gin.Context) {

	var registerData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if !IsValid(registerData.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Data Invalid"})
		return
	}

	hashedpw, err := HashPassword(registerData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Request"})
		return
	}
	_, err = DB.DB.Exec(context.Background(),
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		registerData.Username, hashedpw)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Request"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "user register complete"})
}

func AuthMiddleware() gin.HandlerFunc {
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

func IsValid(input string) bool {
	//Check length input
	if len(input) < 8 {
		return false
	}
	// Define a regular expression to match alphanumeric characters and underscores
	reg := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	// Check if the input matches the pattern
	return reg.MatchString(input)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
