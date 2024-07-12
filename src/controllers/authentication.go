package controllers

import (
	"github.com/BinDruid/spotify-gin/db"
	"github.com/BinDruid/spotify-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *gin.Context) {
	var user models.User
	var userRecord models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Postgres.Where("USERNAME = ?", user.Username).First(&userRecord).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	userAuthenticated := CheckPasswordHash(user.Password, userRecord.Password)
	if !userAuthenticated {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "user not authorized"})
		return
	}
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &models.Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expirationTime.Unix(), 0))},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jwtOutput := models.JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}
	c.JSON(http.StatusOK, jwtOutput)
}

func RegisterUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	hashedPassword, err := HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	user := models.User{Username: newUser.Username, Password: hashedPassword}
	db.Postgres.Create(&user)
	c.IndentedJSON(http.StatusCreated, user)
}
