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
	r.socket.SendMessageToAll(server.Message{Info: info})
}

func (r *app) EventSteamWorksInit() {
	go func() {
		time.Sleep(time.Second * 2)
		r.api.OverlayOpenUrl(fmt.Sprintf("http://localhost:%d", r.socket.Port))
	}()
}

func Run(httpPort int) {
	log.Println("[INFO] Start...")

	app := &app{
		socket: server.NewSocket(fmt.Sprintf(":%d", httpPort)),
		api:    tekken.NewApi(),
	}

	app.api.Subscribe(app)
	app.api.Run()
}
