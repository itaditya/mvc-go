package models

import (
  "github.com/astaxie/beego/orm"
  _ "github.com/lib/pq"
  "os"
)

var ormObject orm.Ormer

func init() {
  dbUrl := os.Getenv("DB_URL")
  orm.RegisterDriver("postgres", orm.DRPostgres)
  orm.RegisterDataBase("default", "postgres", dbUrl, 30)
  orm.RegisterModel(new(Users))
  orm.RunSyncdb("default", false, true)
  ormObject = orm.NewOrm()
}

func GetOrmObject() orm.Ormer {
    return ormObject
}
