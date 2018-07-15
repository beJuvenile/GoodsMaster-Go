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
	set, err := models.GetItemByName("web_config")
	if err != nil {
		i.Data["json"] = err.Error()
	} else {
		i.Data["json"] = set
	}

	i.ServeJSON()
}