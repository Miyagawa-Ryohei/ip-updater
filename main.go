package main

import (
	"github.com/Miyagawa-Ryohei/gode_conf"
	"ip-updater/infra"
	"log"
	"net/http"
	"net/http/cgi"
	"os"
	"path"
)

func handler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	wd , err := os.Getwd()
	configOption := gode_conf.ConfigOption{
		FileName:  os.Getenv("RUNENV"),
		Directory: path.Join(wd, "config"),
		HotReload: false,
	}
	gode_conf.LoadTo(infra.CONF, &configOption)
	if err != nil {
		panic(err)
	}
	if err := cgi.Serve(http.HandlerFunc(handler)); err != nil {
		log.Fatalln(err)
	}
}
