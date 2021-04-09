package service

import (
	"context"
	"fmt"

	"github.com/labstack/echo"
	"github.com/thanhlam/sso-service/model"
	"gopkg.in/mgo.v2/bson"
)

type ArrayRole struct {
	ThingRole []string
}

func GetThingRole(c echo.Context) error {
	thingRoleBody := new(model.ThingRoleBody)
	err := c.Bind(thingRoleBody)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, map[string]interface{}{"code": "6", "message": "Body is Invalid", "data": nil})
	}
	thingID := thingRoleBody.Thingid
	//query
	collection := CNX.Database("users").Collection("things_role")
	if err != nil {
		return c.JSON(400, map[string]interface{}{"code": "9", "message": "Error connect mongoDB", "data": nil})
	}
	var arrayRole ArrayRole
	ctx := context.Background()
	err = collection.FindOne(ctx, bson.M{"thingid": thingID}).Decode(&arrayRole)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(arrayRole)
	return c.JSON(400, map[string]interface{}{"code": "0", "message": "Success", "data": arrayRole.ThingRole})
}
