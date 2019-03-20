package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mchmarny/kdemo/handler"
	"github.com/mchmarny/kdemo/util"
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
	mux.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	})

	// Server
	port := util.MustGetEnv("PORT", "8080")
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	log.Printf("Server starting on port %s \n", port)
	log.Fatal(server.ListenAndServe())

}
