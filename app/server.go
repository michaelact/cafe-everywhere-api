package app

import (
	"net/http"
	"fmt"
)

func NewServer(middleware http.Handler, c *ConfigApplication) *http.Server {
	address := fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
	return &http.Server{
		Addr:    address, 
		Handler: middleware, 
	}
}
