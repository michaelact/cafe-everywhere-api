package middleware

import (
	"net/http"
	"log"

	"github.com/michaelact/cafe-everywhere/helper"
	"github.com/michaelact/cafe-everywhere/app"
)

type AuthMiddleware struct {
	Handler http.Handler
	Key     string
}

func NewAuthMiddleware(handler http.Handler, c *app.ConfigApplication) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler, 
		Key:     c.API.Key, 
	}
}

func (self *AuthMiddleware) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if self.Key == req.Header.Get("x-api-key") {
		log.Println(req.Host, req.URL, "authorized!")
		self.Handler.ServeHTTP(res, req)
	} else {
		helper.WriteToResponseBodyError(res, http.StatusUnauthorized, "UNAUTHORIZED")
	}
}
