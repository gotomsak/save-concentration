package main

import (
	"context"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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

	oldData := new(SaveConcentration)
	newData := new(SaveConcentration)
	if err := c.Bind(newData); err != nil {
		return c.JSON(500, "concentratioon not found")
	}

	mc, ctx := mongoConnect()
	defer mc.Disconnect(ctx)

	filter := bson.M{"_id": newData.ID}

	dbColl := mc.Database(newData.Type).Collection(newData.Measurement)

	err := dbColl.FindOne(context.Background(), filter).Decode(&oldData)
	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "find error")
	}

	oldData.Concentration = append(oldData.Concentration, newData.Concentration)

	res, err := dbColl.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"concentration": oldData.Concentration}})
	fmt.Println(res)
	return c.JSON(200, "ok")
}
