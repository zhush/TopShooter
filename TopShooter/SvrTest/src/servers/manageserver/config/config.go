package config

import (
	"encoding/json"
	"io/ioutil"
	"libs/log"
)

var Conf map[string]interface{}

func init() {
	data, err := ioutil.ReadFile("config/manageserver.json")
	if err != nil {
		log.Fatal("Load conf/manageserver1.json failed! %v", err)
		return
	}

	err = json.Unmarshal(data, &Conf)
	if err != nil {
		log.Fatal("Unmarshal! %v", err)
	}
}
