package config

import (
	"encoding/json"
	"io/ioutil"
	"libs/log"
)

var Conf map[string]interface{}

func init() {
	data, err := ioutil.ReadFile("config/loginserver.json")
	if err != nil {
		log.Fatal("Load conf/loginserver.json failed! %v", err)
		return
	}

	err = json.Unmarshal(data, &Conf)
	if err != nil {
		log.Fatal("Unmarshal! %v", err)
	}
}
