package main

import (
	"net/http"
	_ "net/http/pprof"
	_ "runtime"
	_ "showgirl/models/utils"
	_ "showgirl/routers"

	"github.com/astaxie/beego"
	_ "github.com/golang/protobuf/jsonpb"
	_ "github.com/golang/protobuf/proto"
	_ "gopkg.in/ini.v1"
)

func main() {
	go func() {
		if true {
			beego.Warn(http.ListenAndServe("0.0.0.0:8060", nil))
		}
	}()

	beego.Run()
}
