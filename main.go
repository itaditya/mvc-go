package main

import (
  _ "github.com/joho/godotenv/autoload"
  _ "mvc-go/routers"
  _ "mvc-go/models"
	"github.com/astaxie/beego"
  "strconv"
  "os"
)

func main() {
  port, err := strconv.Atoi(os.Getenv("PORT"))
  if err == nil {
    beego.BConfig.Listen.HTTPPort = port //set port
  }
	beego.Run()
}

