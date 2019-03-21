package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mchmarny/kdemo/client"
)

// LogoHandler handles posted queries
func LogoHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		log.Println("Nil request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	imageURL := r.URL.Query().Get("imageUrl")
	log.Printf("Logo request: %s", imageURL)

	logo, err := client.GetLogoInfo(imageURL)
	if err != nil {
		log.Printf("Error while quering logo service: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(logo)
	if err != nil {
		log.Printf("Error while encoding logo response: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
