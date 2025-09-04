package handlers

import (
	"net/http"
	"notes/inits"
	"notes/internal/models"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

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
		c.JSON(res_code, gin.H{
			"error": "failed to read body",
		})
		Log.Error("failed to read body", c.Request.Method, c.Request.URL.Path, "-", res_code, errors[res_code], strconv.FormatInt(time.Since(t).Milliseconds(), 10), "ms")
		return
	}

	// ckeck if exists
	var user models.User
	inits.DB.First(&user, "username = ?", body.Username)

	if user.ID != 0 {
		res_code := 403
		c.JSON(res_code, gin.H{
			"error": "this username is already taken",
		})
		Log.Error("this username is already taken", c.Request.Method, c.Request.URL.Path, "-", res_code, errors[res_code], strconv.FormatInt(time.Since(t).Milliseconds(), 10), "ms")
		return
	}

	// hash pwd
	hashed_pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		resp_code := 400
		c.JSON(resp_code, gin.H{
			"error": "failed to hash password",
		})
		Log.Error("failed to hash password", c.Request.Method, c.Request.URL.Path, "-", resp_code, errors[resp_code], strconv.FormatInt(time.Since(t).Milliseconds(), 10), "ms")
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
		resp_code := 400
		c.JSON(resp_code, gin.H{
			"error": "error in signing jwt",
		})
		Log.Error("error in signing jwt", c.Request.Method, c.Request.URL.Path, "-", resp_code, errors[resp_code], strconv.FormatInt(time.Since(t).Milliseconds(), 10), "ms")
		return
	}

	// Give cookie with jwt
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)

	// OK
	resp_code := 200
	c.JSON(resp_code, gin.H{})
	Log.Info("", c.Request.Method, c.Request.URL.Path, "-", resp_code, errors[resp_code], strconv.FormatInt(time.Since(t).Milliseconds(), 10), "ms")

}
