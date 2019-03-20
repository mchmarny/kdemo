package handler

import (
	"log"
	"net/http"

	"github.com/mchmarny/kdemo/util"
)

// ViewHandler handles view page
func ViewHandler(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	uid := getCurrentUserID(r)
	if uid != "" {
		log.Printf("User has ID: %s, getting data...", uid)
		// TODO: Call user service to save
		// userData, err := stores.GetData(r.Context(), uid)
		// if err != nil {
		// 	log.Printf("Error while getting user data: %v", err)
		// }else{
		// 	data = userData
		// }
	}

	data["version"] = util.MustGetEnv("RELEASE", "NOT SET")

	if err := templates.ExecuteTemplate(w, "view", data); err != nil {
		log.Printf("Error in view template: %s", err)
	}

}
