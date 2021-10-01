package infra

import "ip-updater/entity"

var CONF = &entity.IPUpdaterConfig{
	Auth: &entity.AuthOption{
		Username: "",
		Password: "",
	},
}
