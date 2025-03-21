package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ogarridojimenez/go-jwt/initializers"
	"github.com/ogarridojimenez/go-jwt/models"
)

func RequireAuth(c *gin.Context) {
	//TODO: Get the cookie off req

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//TODO: Decode/validate it
	// Parse takes the token string and a function for looking up the key. The latter is especially

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//TODO: Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)

		}
		//TODO: Find the user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//TODO: Attach to req
		c.Set("user", user)

		//TODO: Continue

		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
