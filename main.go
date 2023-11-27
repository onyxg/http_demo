package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(delayImageMiddleware)
	e.Static("/", "static")

	if os.Getenv("HTTP2") != "" {
		log.Printf("starting http2 server...")
		//net/http.Server is a http1.1 server with optional http2 support
		log.Fatal(http.ListenAndServeTLS(":9002", "./ssl/localhost.crt", "./ssl/localhost.key", e))
	} else {
		log.Printf("starting http server...")
		log.Fatal(http.ListenAndServe(":9001", e))
	}
}

func delayImageMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.Contains(c.Request().RequestURI, "/images/") {
			log.Printf("delaying image requests...")
			time.Sleep(time.Second * 2)
		}

		return next(c)
	}
}
