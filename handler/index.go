package handler

import (
	"log"
	"net/http"

	"github.com/mchmarny/kdemo/util"
)

// DefaultHandler handles index page
func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	data["version"] = util.MustGetEnv("RELEASE", "NOT SET")

	if err := templates.ExecuteTemplate(w, "index", data); err != nil {
		log.Printf("Error in index template: %s", err)
	}

}
