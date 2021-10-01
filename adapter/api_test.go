package adapter

import (
	"github.com/Miyagawa-Ryohei/gode_conf"
	"ip-updater/infra"
	"log"
	"os"
	"path"
	"testing"
)

func TestUpdateIPAddrTest(t *testing.T) {
	wd , _ := os.Getwd()
	log.Print(wd)
	configOption := gode_conf.ConfigOption{
		FileName:  os.Getenv("RUNENV"),
		Directory: path.Join(wd, "..","config"),
		HotReload: false,
	}
	gode_conf.LoadTo(infra.CONF, &configOption)

	UpdateIPAddr(infra.CONF)

}
