package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

    "github.com/mchmarny/kuser/message"
    "github.com/mchmarny/kdemo/util"
)

const (
	localUserServiceURL = "http://kuser.demo.svc.cluster.local"
)

var (
	userServiceURL  = util.MustGetEnv("KUSER_SERVICE_URL", localUserServiceURL)
)

// GetUser provides client lib for KUser
func GetUser(id string) (usr *message.KUser, err error) {

	url := fmt.Sprintf("%s/user/%s", userServiceURL, id)
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

// SaveUser provides client lib for KUser
func SaveUser(usr *message.KUser) error {

    url := fmt.Sprintf("%s/user", userServiceURL)
	b, err := json.Marshal(usr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Printf("Response status: %s", resp.Status)
	if resp.StatusCode != http.StatusOK {
		return errors.New("Invalid response code from KUser service")
	}

	return nil

}
