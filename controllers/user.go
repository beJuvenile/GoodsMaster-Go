package controllers

import (
	"github.com/astaxie/beego"
	"GoodsMaster-Go/models"
	"encoding/json"
	"crypto/md5"
	"io"
	"fmt"
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
	ret, err := models.GetUserByAccount(request.Account)

	if err == nil {
		u.Data["json"] = map[string]interface{}{
			"retCode": 20000,
			"retMsg": "用户存在",
			"retData": ret.PasswordSalt,
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

// @Title UserLogin
// @Description do user login
// @Param	body		body 	models.User		true		"body for user content"
// @Success 200
// @Failure 403 body is empty
// @router /login [post]
func (u *UserController) UserLogin() {
	var request models.User

	json.Unmarshal(u.Ctx.Input.RequestBody, &request)
	ret, err := models.GetUserByAccount(request.Account)

	if err == nil {
		// 验证密码
		passwordSalt := beego.AppConfig.Strings("passwordSalt")
		w := md5.New()
		io.WriteString(w, request.Password + passwordSalt[ret.PasswordSalt])
		password := fmt.Sprintf("%x", w.Sum(nil))

		if password == ret.Password {

		}

		u.Data["json"] = map[string]interface{}{
			"retCode": 20000,
			"retMsg": "登录成功",
			"retData": passwordSalt,
			"subData": ""}
	} else {
		u.Data["json"] = map[string]interface{}{
			"retCode": 50000,
			"retMsg": "登录失败",
			"retData": "",
			"subData": err.Error()}
	}

	u.ServeJSON()
}