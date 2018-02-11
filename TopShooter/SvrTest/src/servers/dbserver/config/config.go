package config

import (
	"encoding/json"
	"io/ioutil"
	"libs/log"
)

var DBConf map[string]interface{}

func init() {
	data, err := ioutil.ReadFile("config/dbserver.json")
	if err != nil {
		log.Fatal("Load conf/server.json failed! %v", err)
		return
	}

	err = json.Unmarshal(data, &DBConf)
	if err != nil {
		log.Fatal("Unmarshal! %v", err)
	}
}
