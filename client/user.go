package client

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/mchmarny/kuser/message"
)

var (
	getUserURL = "http://kuser.demo.svc.cluster.local/user/"
)

// GetUser provides client lib for KUser
func GetUser(id string) (usr *message.KUser, err error) {
	url := getUserURL + id
	return getUserFromService(url)
}

// GetUser provides client lib for KUser
func getUserFromService(url string) (usr *message.KUser, err error) {

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.Printf("Response status: %s", resp.Status)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Error in depending service")
	}

	var u message.KUser
	err = json.NewDecoder(resp.Body).Decode(&u)

	return &u, nil

}
