package controllers

import (
	"github.com/astaxie/beego"
	"GoodsMaster-Go/models"
	"encoding/json"
	//"fmt"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title UserExist
// @Description check user exist
// @Param	body		body 	models.User		true		"body for user content"
// @Success 200
// @Failure 403 body is empty
// @router /check [post]
func (u *UserController) CheckUser() {
	var request models.User

	json.Unmarshal(u.Ctx.Input.RequestBody, &request)
	_, err := models.GetUserByAccount(request.Account)

	if err == nil {
		u.Data["json"] = map[string]interface{}{
			"retCode": 20000,
			"retMsg": "用户存在",
			"retData": "",
			"subData": ""}
	} else {
		u.Data["json"] = map[string]interface{}{
			"retCode": 50000,
			"retMsg": "未找到用户",
			"retData": "",
			"subData": err.Error()}
	}

	u.ServeJSON()
}

