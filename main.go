package main

import (
	"github.com/labstack/echo/v4"
)

func router() *echo.Echo {
	e := echo.New()
	e.GET("/get_id", getSaveImagesID)

	e.POST("/save_concent", saveConcentration)
	return e
}

func main() {
	envLoad()
	e := router()

	e.Logger.Fatal(e.Start(":1323"))
	// e.Logger.Fatal(e.StartTLS(":1323", "./fullchain.pem", "./privkey.pem"))
}
