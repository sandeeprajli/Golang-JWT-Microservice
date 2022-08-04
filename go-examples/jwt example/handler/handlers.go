package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("secret")

var users = map[string]string{
	"root": "rootpassword",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signin(c *gin.Context) {
	var creds Credentials
	err := json.NewDecoder(c.Request.Body).Decode(&creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "err"})
		return
	}

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorixed"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internal server error"})
		return
	}

	c.SetCookie("token", tokenString, expirationTime.Minute(), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Welcome(g *gin.Context) {

	c, err := g.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			g.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		g.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	g.Header("result", c)

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(c, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			g.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		g.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	if !tkn.Valid {
		g.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})

		return
	}
	g.JSON(http.StatusOK, gin.H{"welcome": claims.Username})

}
