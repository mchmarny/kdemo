package handler

import (
	"encoding/json"
	"log"
	"time"
	"net/http"

	"github.com/mchmarny/kdemo/client"
	"github.com/mchmarny/kdemo/util"
    "github.com/mchmarny/kuser/message"
)

// LogoHandler handles posted queries
func LogoHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	uid := getCurrentUserID(r)
	if uid == "" {
		log.Println("User not authenticated")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	event := &message.KUserEvent{
		ID: util.MakeUUID(),
		On: time.Now(),
		UserID: uid,
		Data: []*message.KDataItem{
			&message.KDataItem{ Key:"logo-request", Value: imageURL },
			&message.KDataItem{ Key:"logo-response", Value: logo.Description },
		},
	}

	err = client.SaveUserEvent(event)
	if err != nil {
		log.Printf("Error while saving logo event: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}



}
