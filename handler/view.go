package handler

import (
	"log"
	"net/http"

	ev "github.com/mchmarny/gcputil/env"
	"github.com/mchmarny/kdemo/client"
)

// ViewHandler handles view page
func ViewHandler(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	uid := getCurrentUserID(r)
	if uid == "" {
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}

	log.Printf("User has ID: %s, getting data...", uid)
	usr, err := client.GetUser(uid)
	if err != nil {
		log.Printf("Error while getting user data: %v", err)
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}

	data["name"] = usr.Name
	data["email"] = usr.Email
	data["pic"] = usr.Picture
	data["version"] = ev.MustGetEnvVar("RELEASE", "NOT SET")

	log.Printf("Data: %v", data)

	if err := templates.ExecuteTemplate(w, "view", data); err != nil {
		log.Printf("Error in view template: %s", err)
	}

}
