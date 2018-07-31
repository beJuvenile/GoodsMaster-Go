package controllers

import (
	"GoodsMaster-Go/models"

	"github.com/astaxie/beego"
)

// Operations about init
type InitsController struct {
	beego.Controller
}

// @Title WebBaseInit
// @Description home base config data
// @Success 200 object models.inits
// @Failure 403 bad request
// @router /web [get]
func (i *InitsController) WebBaseConfig() {

	set, err := models.GetSettingByName("web_config")
	if err != nil {
		i.Data["json"] = map[string]interface{}{
			"retCode": 50000,
			"retMsg": "服务器内部错误",
			"retData": "",
			"subData": err.Error()}
	} else {
		i.Data["json"] = map[string]interface{}{
			"retCode": 20000,
			"retMsg": "请求成功",
			"retData": set,
			"subData": ""}
	}

	i.ServeJSON()
}