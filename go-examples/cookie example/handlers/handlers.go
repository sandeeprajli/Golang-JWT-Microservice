package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

var cred = credentials{
	Password: "password",
	Username: "username",
}

func SignUp(c *gin.Context) {
	sessionToken := cred.Username + " " + cred.Password
	c.SetCookie("session_token", sessionToken, 120, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "sign up"})
}

func Login(c *gin.Context) {
	var logCred credentials
	if err := c.BindJSON(&logCred); err != nil {
		c.Error(err)
		return
	}
	string, err := c.Cookie("session_token")
	if err != nil {
		c.Error(err)
		return
	}
	credCookie := strings.Split(string, " ")
	if credCookie[0] == cred.Username && credCookie[1] == cred.Password {
		c.JSON(http.StatusOK, gin.H{"status": "logged in"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	}
}
