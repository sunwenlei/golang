package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

var secretKey = "75c92a074c341e9964329c0550c2673730ed8479c885c43122c90a2843177d5ef21cb50cfadcccb20aeb730487c11e09ee4dbbb02387242ef264e74cbee97213"

func main() {
	r := gin.Default()

	r.GET("/api/", func(c *gin.Context) {
		/*
		   アルゴリズムの指定
		*/
		token := jwt.New(jwt.GetSigningMethod("HS256"))

		token.Claims = jwt.MapClaims{
			"user": "ゲスト",
			"auth": "123",
			"plan": "Abc",
			"exp":  time.Now().Add(time.Hour * 1).Unix(),
		}

		/*
		   トークンに対して署名の付与
		*/
		tokenString, err := token.SignedString([]byte(secretKey))
		if err == nil {
			c.JSON(200, gin.H{"token": tokenString})
		} else {
			c.JSON(500, gin.H{"message": "Could not generate token"})
		}
	})

	r.GET("/api/verification/", func(c *gin.Context) {
		/*
		   署名の検証
		*/
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := []byte(secretKey)
			return b, nil
		})

		if err == nil {
			claims := token.Claims.(jwt.MapClaims)
			var exp time.Time
			exp = time.Unix(int64(claims["exp"].(float64)), 0)

			msg := fmt.Sprintf("こんにちは、「 %s 」さん、あなたの権限コードは「%s」、プランコードは「%s」、有効期間は「%s」", claims["user"], claims["auth"], claims["plan"], exp.String())
			c.JSON(200, gin.H{"message": msg})
		} else {
			c.JSON(401, gin.H{"error": fmt.Sprint(err)})
		}
	})

	r.Run(":8080")
}
