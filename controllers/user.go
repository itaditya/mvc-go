package controllers

import (
  models "mvc-go/models"
  "github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
  "fmt"
  "encoding/json"
)

var ORM orm.Ormer

type UserController struct {
	beego.Controller
}

func init() {
  ORM = models.GetOrmObject()
}

func (ctrl *UserController) Get() {
  var users[]*models.Users
  ORM.QueryTable("users").All(&users, "UserId", "Email", "UserName")

	ctrl.Data["Users"] = users
	ctrl.TplName = "user.tpl"
}

func (ctrl *UserController) List() {
  defer ctrl.ServeJSON()
  jsonResponse := make(map[string] interface{})

  var users[]*models.Users
  numRows, err := ORM.QueryTable("users").All(&users, "UserId", "Email", "UserName")
  if err == nil {
    fmt.Println("Rows Found: ", numRows)
    jsonResponse["message"] = "success"
    jsonResponse["data"] = &users
  } else {
    jsonResponse["message"] = "failed"
  }

  ctrl.Data["json"] = jsonResponse
}

func (ctrl *UserController) Post() {
  defer ctrl.ServeJSON()

  var newUser models.Users
  jsonResponse := make(map[string] interface{})

  json.Unmarshal(ctrl.Ctx.Input.RequestBody, &newUser)

  _, err := ORM.Insert(&newUser)
  if err == nil {
    jsonResponse["message"] = "success"
    jsonResponse["data"] = &newUser
  } else {
    jsonResponse["message"] = "failed"
  }

  ctrl.Data["json"] = jsonResponse
}

func (ctrl *UserController) Delete() {
  defer ctrl.ServeJSON()

  jsonResponse := make(map[string] string)

  var userId int
  ctrl.Ctx.Input.Bind(&userId, "user_id")
  user := models.Users{UserId: userId}

  numRowsAffected, err := ORM.Delete(&user)
  if(err == nil && numRowsAffected != 0) {
    fmt.Println("Rows Affected: ", numRowsAffected)
    jsonResponse["message"] = "success"
  } else {
    jsonResponse["message"] = "failed"
  }

  ctrl.Data["json"] = jsonResponse
}
