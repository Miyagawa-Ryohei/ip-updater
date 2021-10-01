package adapter

import (
	"fmt"
	"ip-updater/entity"
	"log"
	"net/http"
)

func UpdateIPAddr(conf *entity.IPUpdaterConfig) {

	client := new(http.Client)

	resp, err := client.Get(fmt.Sprintf("https://%s:%s@ipv4.mydns.jp/login.html", conf.Auth.Username,conf.Auth.Password))

	if err != nil {
		log.Fatal(err)
		return
	}
	res := []byte{}
	resp.Body.Read(res)
	if resp.StatusCode == 200 {
		log.Printf("OK")
		log.Print(string(res))
	}
}