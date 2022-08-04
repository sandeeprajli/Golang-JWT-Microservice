package main

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
)

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %s!", samlsp.AttributeFromContext(r.Context(), "cn"))
// }

func main() {

	keyPair, err := tls.LoadX509KeyPair("myservice.cert", "myservice.key")
	if err != nil {
		panic(err) // TODO handle error
	}

	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		panic(err) // TODO handle error
	}

	idpMetadataURL, err := url.Parse("https://samltest.id/saml/idp")
	if err != nil {
		panic(err) // TODO handle error
	}

	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient, *idpMetadataURL)
	if err != nil {
		fmt.Print(err)
		panic(err) // TODO handle error
	}

	rootURL, err := url.Parse("http://localhost:8080")
	if err != nil {
		panic(err) // TODO handle error
	}

	samlSP, _ := samlsp.New(samlsp.Options{
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
	})
	router := gin.Default()
	router.Any("/saml/*action", gin.WrapH(samlSP))
	api := router.Group("/api")
	api.Use(adapter.Wrap(samlSP.RequireAccount))
	api.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "tested")
	})
	router.Run(":8080")
}
