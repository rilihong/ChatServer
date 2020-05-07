package main

import (
	"encoding/json"
	"flag"
	"os"
	"strconv"

	"github.com/rilihong/ChatServer/src/chatagent/agent"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
)

func init() {
	path := flag.String("path", "./conf.js", "server conf")
	flag.Parse()
	conf := agent.GetConf()
	ok := conf.Init(*path)
	if ok == false {
		log.Fatal().Str("error", "open conf file err").Send()
		os.Exit(-1)
	}
	logfile := "agent." + strconv.Itoa(os.Getpid()) + ".log"
	file, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Str("error", "open log file err").Send()
		os.Exit(-1)
	}
	log.Logger = log.Output(file)
	level, _ := zerolog.ParseLevel(conf.Leve)
	log.Logger.Level(level)
}

func main() {
	conf := agent.GetConf()
	msg, _ := json.Marshal(conf)
	log.Info().Str("conf", string(msg)).Send()
}
