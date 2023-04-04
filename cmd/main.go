package main

import (
	"gin_mall/conf"
	"gin_mall/routers"
)

func main() {
	conf.Init()
	r := routers.NewRouter()
	_ = r.Run(conf.HttpPort)
}
