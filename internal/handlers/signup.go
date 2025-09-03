package handlers

import (
	"log/slog"
	"net/http"
	"notes/inits"
	"notes/internal/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// read body
	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed to read body",
		})

		slog.Error(c.FullPath(), slog.String("method", c.Request.Method))

		return
	}

	// ckeck if exists
	var user models.User
	inits.DB.First(&user, "username = ?", body.Username)

	if user.ID != 0 {
		c.JSON(403, gin.H{
			"error": "this username is already taken",
		})
		return
	}

	// hash pwd
	hashed_pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	// add new user
	newUser := models.User{Username: body.Username, Password: string(hashed_pwd)}
	inits.DB.Create(&newUser)

	// add jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": newUser.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "error in signing jwt",
		})
		return
	}

	// Give cookie with jwt
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)

	// OK
	c.JSON(200, gin.H{})

}
