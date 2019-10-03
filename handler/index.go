package handler

import (
	"log"
	"net/http"

	ev "github.com/mchmarny/gcputil/env"
)

// DefaultHandler handles index page
func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	data["version"] = ev.MustGetEnvVar("RELEASE", "NOT SET")

	if err := templates.ExecuteTemplate(w, "index", data); err != nil {
		log.Printf("Error in index template: %s", err)
	}

}
