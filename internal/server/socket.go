package server

import (
	"context"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nullentrypoint/tekken7-web-displayer/client"
	"github.com/nullentrypoint/tekken7-web-displayer/pkg/livegollection"
)

type Socket struct {
	*livegollection.LiveGollection[MessageID, Message]
	Addr string
	Port int
}

func NewSocket(addr string) *Socket {
	serveStatic()

	// When we create the LiveGollection we need to specify the type parameters for items' id and items themselves.
	socket := Socket{
		LiveGollection: livegollection.NewLiveGollection[MessageID, Message](context.TODO(), NewCollection(), log.Default()),
		Addr:           addr,
		Port:           getPort(addr),
	}

	// After we created liveGoll, to enable it we just need to register a route handled by liveGoll.Join.
	http.HandleFunc("/livegollection", socket.Join)

	go func(xAddr string) {
		log.Fatal(http.ListenAndServe(xAddr, nil))
	}(socket.Addr)

	return &socket
}

func getPort(addr string) int {
	parts := strings.Split(addr, ":")
	if len(parts) < 2 {
		return 0
	}

	port, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return 0
	}

	return port
}

func serveStatic() {
	for r, fP := range client.GetStaticFiles() {
		// Prevents passing loop variables to a closure.
		route, fileContent := r, fP
		http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			ctype := mime.TypeByExtension(filepath.Ext(route))
			w.Header().Set("Content-Type", ctype)
			w.Write(fileContent)
		})
	}
}
