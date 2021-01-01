package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func saveConcentration(c echo.Context) error {
	var token string = ""
	if token = c.Request().Header["Access-Token"][0]; token == "" {
		return c.JSON(500, "access token not found")
	}
	check := os.Getenv("TOKEN")
	if check != token {
		return c.JSON(500, "access token mistaken")
	}
	conc := new(SaveConcentration)
	if err := c.Bind(conc); err != nil {
		return c.JSON(500, "concentration not found")
	}
	mc, ctx := mongoConnect()
	defer mc.Disconnect(ctx)
	dbColl := mc.Database(conc.TypeData).Collection(conc.Measurement)
	res, err := dbColl.InsertOne(context.Background(), conc)
	if err != nil {
		return c.JSON(500, "insert error")
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok != true {
		fmt.Println(oid)
		return c.JSON(http.StatusInternalServerError, "Not objectid.ObjectID, do what you want")
	}
	return c.JSON(200, "ok")
}
