package main

import (
	"github.com/astaxie/beego"
	"github.com/qor/qor"
	"github.com/qor/qor/admin"
	"github.com/sunfmin/beego_with_qor/base/db"
	"net/http"
)

func main() {

	adm := admin.New(&qor.Config{DB: &db.DB})
	adm.AddResource(&db.User{}, &admin.Config{Menu: []string{"管理"}})

	mux := http.NewServeMux()
	adm.MountTo("/admin", mux)

	beego.Handler("/admin", mux)
	beego.Handler("/admin/*", mux)
	beego.Run()

}
