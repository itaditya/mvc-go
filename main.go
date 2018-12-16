package main

import (
  _ "github.com/joho/godotenv/autoload"
  _ "mvc-go/routers"
  models "mvc-go/models"
  "github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

var ORM orm.Ormer

func init() {
  ORM = models.GetOrmObject()
}

func main() {
	beego.Run()
}

