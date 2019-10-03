package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	ev "github.com/mchmarny/gcputil/env"
	"github.com/mchmarny/kdemo/util"
	"github.com/mchmarny/klogo/message"
)

const (
	localLogoServiceURL = "http://logo.demo.svc.cluster.local"
)

var (
	logoServiceURL = ev.MustGetEnvVar("LOGO_SERVICE_URL", localLogoServiceURL)
)

// GetLogoInfo provides client lib for KLogo
func GetLogoInfo(url string) (logo *message.LogoResponse, err error) {

	imgReg := &message.LogoRequest{
		ID:       util.MakeUUID(),
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
