package agent

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rs/zerolog/log"
)

//RedisC RedisConf
type RedisC struct {
	RedisIP   string `json:"redis_ip"`
	RedisPort int    `json:"redis_port"`
	RedisAuth string `json:"redis_auth"`
}

//Conf agent
type Conf struct {
	ID    int    `json:"id"`
	IP    string `json:"ip"`
	Port  int    `json:"port"`
	Name  string `json:"name"`
	Redis RedisC `json:"redis"`
	Leve  string `json:"level"`
}

//ConfInfo conf
var confInfo *Conf

//Init conf
func (conf *Conf) Init(path string) bool {
	jdata, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal().Str("file", "can not open").Send()
		return false
	}
	err = json.Unmarshal(jdata, conf)
	if err != nil {
		log.Fatal().Str("Unmarshal", "Unmarshal error").Send()
		return false
	}
	return true
}

//GetConf return conf
func GetConf() *Conf {
	if confInfo == nil {
		confInfo = new(Conf)
	}
	return confInfo
}
