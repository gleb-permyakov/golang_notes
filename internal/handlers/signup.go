package handlers

import (
	"net/http"
	"notes/inits"
	"notes/internal"
	"notes/internal/models"
	"notes/pkg/logger"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var Log *logger.Loga = &logger.Log

func Signup(c *gin.Context) {

	t := time.Now()

	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// read body
	err := c.Bind(&body)
	if err != nil {
		res_code := 400
		res_msg := "failed to read body"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

	// ckeck if exists
	var user models.User
	inits.DB.First(&user, "username = ?", body.Username)

	if user.ID != 0 {
		res_code := 403
		res_msg := "this username is already taken"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

	// hash pwd
	hashed_pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		res_code := 400
		res_msg := "failed to hash password"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
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
		res_code := 400
		res_msg := "error in signing jwt"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

	// Give cookie with jwt
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)

	// OK
	res_code := 200
	c.JSON(res_code, gin.H{})
	Log.Info("", internal.LoggerParams(c, res_code, t)...)

}
