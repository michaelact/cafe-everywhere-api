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
	Origin  string
}

func NewAuthMiddleware(handler http.Handler, c *app.ConfigApplication) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler, 
		Key:     c.API.Key, 
		Origin:  c.API.Origin,
	}
}

func (self *AuthMiddleware) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-api-key")
	res.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH, DELETE")
	res.Header().Set("Access-Control-Allow-Origin", self.Origin)
	
	if req.Method == "OPTIONS" {
		helper.WriteToResponseBodyError(res, http.StatusOK, "OK")
	} else if self.Key == req.Header.Get("x-api-key") {
		log.Println(req.Host, req.URL, req.UserAgent(), "authorized!")
		self.Handler.ServeHTTP(res, req)
	} else {
		helper.WriteToResponseBodyError(res, http.StatusUnauthorized, "UNAUTHORIZED")
	}
}
