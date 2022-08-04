package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://127.0.0.1:8000/callback",
	ClientID:     "133980786539-n376srcg19fcq41n1glr05r4icbjk2nc.apps.googleusercontent.com",
	ClientSecret: "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",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func Login(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL("random")
	c.Redirect(301, url)
	c.JSON(301, gin.H{"logging in...": ""})
}

func Callback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"redirected to callback...": ""})
}
