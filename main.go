package main

import (
	"os"
	"log"
	"strconv"
	"net/http"
	"github.com/shideon/toml-go"
)

func main() {
	cfgFile := os.Getenv("WX3_CONFIG");
	if cfgFile == "" {
		cfgFile = "/etc/wx3.conf"
	}

	parser := toml.Parser{}
	cfg  := parser.ParseFile(cfgFile)

	if err := http.ListenAndServe(":"+strconv.Itoa(cfg.GetInt("server.port")), nil); err != nil {
		log.Panic("Wx3 web server not started.")
	}
}