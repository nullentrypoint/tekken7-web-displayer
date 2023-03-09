package main

import (
	"log"
	"os"

	"gopkg.in/ini.v1"

	"github.com/nullentrypoint/tekken7-web-displayer/internal/app"
)

const defaultHttpPort = 8080

func main() {

	cfg, err := ini.Load("config.ini")
    if err != nil {
        log.Printf("Fail to read file: %v", err)
        os.Exit(1)
    }

	httpPort := cfg.Section("server").Key("http_port").MustInt(defaultHttpPort)

	app.Run(httpPort)
}
