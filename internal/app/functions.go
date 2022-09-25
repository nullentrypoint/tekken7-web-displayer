package app

import (
	"fmt"
	"log"
	"time"

	"github.com/nullentrypoint/tekken7-web-displayer/internal/server"
	"github.com/nullentrypoint/tekken7-web-displayer/pkg/tekken"
)

type app struct {
	socket *server.Socket
	api    *tekken.Api
}

func (r *app) EventNewChallenger(info tekken.PlayerInfo) {
	r.socket.SendMessageToAll(server.MyItem{Info: info})
}

func (r *app) EventSteamWorksInit() {
	go func() {
		time.Sleep(time.Second * 15)
		r.api.OverlayOpenUrl(fmt.Sprintf("http://%s", r.socket.Addr))
	}()
}

func Run() {
	log.Println("[INFO] Start...")

	app := &app{
		socket: server.NewSocket("localhost:8080"),
		api:    tekken.NewApi(),
	}

	app.api.Subscribe(app)
	app.api.Run()
}
