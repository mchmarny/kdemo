package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

    "github.com/mchmarny/klogo/message"
    "github.com/mchmarny/kdemo/util"
)

const (
	localLogoServiceURL = "http://klogo.demo.svc.cluster.local"
)

var (
	logoServiceURL  = util.MustGetEnv("KLOGO_SERVICE_URL", localLogoServiceURL)
)

// GetLogoInfo provides client lib for KLogo
func GetLogoInfo(url string) (logo *message.LogoResponse, err error) {

	imgReg := &message.LogoRequest{
		ID: util.MakeUUID(),
		ImageURL: url,
	}

	b, err := json.Marshal(imgReg)
	req, err := http.NewRequest("POST", logoServiceURL, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.Printf("Response status: %s", resp.Status)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Invalid response code from KUser service")
	}

	var imrResp message.LogoResponse
	err = json.NewDecoder(resp.Body).Decode(&imrResp)

	return &imrResp, nil

}
