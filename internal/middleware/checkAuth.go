package middleware

import (
	"notes/internal"
	"notes/pkg/logger"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var Log *logger.Loga = &logger.Log

func CheckAuth(c *gin.Context) {

	t := time.Now()

	// reading cookie
	tokenString, err := c.Cookie("Auth")
	if err != nil {
		res_code := 403
		res_msg := "no cookie - unauthorised"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

	// getting token
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SECRET")), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		res_code := 403
		res_msg := "no cookie - unauthorised"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

	// check if OK
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			res_code := 403
			res_msg := "old cookie"
			c.JSON(res_code, gin.H{
				"error": res_msg,
			})
			Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
			return
		} else {
			c.Set("userID", claims["sub"])
			c.Next()
		}
	} else {
		res_code := 400
		res_msg := "bad jwt"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}
}
