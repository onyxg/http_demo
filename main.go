package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"runtime"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(delayImageMiddleware)
	e.Static("/", "static")

	e.GET("/health", healthHandler)
	e.GET("/metrics", metricsHandler)

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

func healthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func metricsHandler(c echo.Context) error {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	metrics := map[string]uint64{
		"Alloc":      m.Alloc,
		"TotalAlloc": m.TotalAlloc,
		"Sys":        m.Sys,
		"NumGC":      uint64(m.NumGC),
	}
	return c.JSON(http.StatusOK, metrics)
}
