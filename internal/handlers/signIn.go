package handlers

import (
	"net/http"
	"notes/inits"
	"notes/internal"
	"notes/internal/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c *gin.Context) {
	t := time.Now()

	// get username and pwd
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	err := c.Bind(&body)
	if err != nil {
		res_code := 400
		res_msg := "invalid body"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

	// get hached pwd from db
	var user models.User
	result := inits.DB.Where("username = ?", body.Username).
		Find(&user)

	if result != nil {
		res_code := 400
		res_msg := "incorrect username or password"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

	hashed_pwd := user.Password

	// compare pwd
	err = bcrypt.CompareHashAndPassword([]byte(hashed_pwd), []byte(body.Password))
	if err != nil {
		res_code := 400
		res_msg := "incorrect username or password"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

	// make jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(os.Getenv("SECRET"))
	if err != nil {
		res_code := 500
		res_msg := "internal error"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error("error in signing jwt", internal.LoggerParams(c, res_code, t)...)
		return
	}

	// give cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)

	// OK
	res_code := 200
	res_msg := "signed in successfully"
	c.JSON(res_code, gin.H{
		"message": res_msg,
		"user_id": user.ID,
	})
	Log.Info("", internal.LoggerParams(c, res_code, t)...)

}
