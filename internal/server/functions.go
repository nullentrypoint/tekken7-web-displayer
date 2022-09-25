package server

import (
	"context"
	"fmt"
	"log"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/nullentrypoint/tekken7-web-displayer/client"
	"github.com/nullentrypoint/tekken7-web-displayer/pkg/livegollection"
)

type Socket struct {
	*livegollection.LiveGollection[IdType, MyItem]
	Addr string
}

func NewSocket(addr string) *Socket {
	serveStatic()

	myColl, err := NewMyCollection()
	if err != nil {
		log.Fatal(fmt.Errorf("error when creating MyCollection: %v", err))
	}

	// When we create the LiveGollection we need to specify the type parameters for items' id and items themselves.
	socket := Socket{
		LiveGollection: livegollection.NewLiveGollection[IdType, MyItem](context.TODO(), myColl, log.Default()),
		Addr:           addr,
	}

	// After we created liveGoll, to enable it we just need to register a route handled by liveGoll.Join.
	http.HandleFunc("/livegollection", socket.Join)

	go func(xAddr string) {
		log.Fatal(http.ListenAndServe(xAddr, nil))
	}(socket.Addr)

	return &socket
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
