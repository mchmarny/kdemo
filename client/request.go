package client

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"context"
	"log"
	"encoding/base64"
	"crypto/rand"
	"encoding/json"
	"time"
	"errors"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/mchmarny/kdemo/util"
	"github.com/mchmarny/kuser/message"

)

var (
    getUserURL = "http://kuser.demo.svc.cluster.local/user/"
)


// GetUser provides client lib for KUser
func GetUser(id string) (usr *message.KUser, err error) {

    url := getUserURL + id
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    log.Printf("Response status: %s", resp.Status)
    if resp.Status != http.StatusOK {
        return nil, errors.New("Error in depending service")
    }

    var usr message.KUser
    err := json.NewDecoder(r.Body).Decode(&usr)

    return usr, nil

}