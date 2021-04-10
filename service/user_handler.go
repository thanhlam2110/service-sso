package service

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo"
	"github.com/thanhlam/sso-service/model"
)

func UserPushCommand(c echo.Context) error {
	userRoleBody := new(model.UserRoleBody)
	err := c.Bind(userRoleBody)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, map[string]interface{}{"code": "6", "message": "Body is Invalid", "data": nil})
	}
	token := userRoleBody.Token
	thingId := userRoleBody.Thingid
	//role := userRoleBody.Role
	//check user status
	authResponse, err := BasicAuth(token)
	if err != nil {
		return c.JSON(200, map[string]interface{}{"code": "10", "message": "Connection refused", "data": nil})
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(authResponse), &result)
	if result["error"] != nil {
		if result["message"] != nil {
			return c.JSON(200, map[string]interface{}{"code": "7", "message": (result["message"]).(string), "data": nil})
		}
		return c.JSON(200, map[string]interface{}{"code": "5", "message": (result["error"]).(string), "data": nil})
	}
	attributes := result["attributes"].(map[string]interface{})
	if attributes["userstatus"].(string) != "ACTIVE" {
		return c.JSON(200, map[string]interface{}{"code": "8", "message": "User Disable", "data": nil})
	}
	//check user status
	//check assign things
	userid := attributes["userid"].(string)
	//fmt.Println("userid: " + userid)
	_, arrayUserAssign := CheckUserAssignThing(thingId)
	//fmt.Println(arrayUserAssign)
	if contains(arrayUserAssign, userid) != true {
		return c.JSON(200, map[string]interface{}{"code": "3", "message": "Don't Allow", "data": nil})
	}
	//check assign things
	//check role
	_, userRole := CheckUserRole(userid)
	//fmt.Println(userRole)
	/*if contains(userRole, role) != true {
		return c.JSON(200, map[string]interface{}{"code": "3", "message": "Don't Have Role", "data": nil})
	}
	//check role*/
	return c.JSON(200, map[string]interface{}{"code": "0", "message": "Success", "data": userRole})
}

//check element exits in array
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
