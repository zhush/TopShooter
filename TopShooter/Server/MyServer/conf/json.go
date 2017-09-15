package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/name5566/leaf/log"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSocketAddr string
	TcpAddr     string
	MaxConnNum  int
}

func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("Load conf/server.json failed! %v", err)
		return
	}

	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("Unmarshal! %v", err)
	}

}
