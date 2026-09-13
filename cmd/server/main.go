package main

import (
	"net/http"
	"os"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/healthz", func(c echo.Context) error { return c.JSON(http.StatusOK, map[string]string{"status": "ok"}) })
	e.POST("/incidents", func(c echo.Context) error { return c.JSON(http.StatusNotImplemented, map[string]string{"error": "incident workflow pending"}) })
	port := os.Getenv("PORT")
	if port == "" { port = "8080" }
	e.Logger.Fatal(e.Start(":" + port))
}
