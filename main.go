package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/mchmarny/kdemo/handler"

	ev "github.com/mchmarny/gcputil/env"
)

var (
	logger = log.New(os.Stdout, "", 0)
	port   = ev.MustGetEnvVar("PORT", "8080")
)

func main() {

	handler.InitHandlers()
	mux := http.NewServeMux()

	// Static
	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	// Handlers
	mux.HandleFunc("/", handler.DefaultHandler)
	mux.HandleFunc("/auth/login", handler.OAuthLoginHandler)
	mux.HandleFunc("/auth/callback", handler.OAuthCallbackHandler)
	mux.HandleFunc("/auth/logout", handler.LogOutHandler)
	mux.HandleFunc("/view", handler.ViewHandler)
	mux.HandleFunc("/logo", handler.LogoHandler)
	mux.HandleFunc("/_health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	})

	// Server
	addr := net.JoinHostPort("0.0.0.0", port)
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("Server starting on port %s \n", port)
	log.Fatal(server.ListenAndServe())

}
