package client

import (
	"testing"
	"log"

	"github.com/stretchr/testify/assert"
)

func TestRestHandler(t *testing.T) {

	url := "http://kuser.demo.knative.tech/user/handler-123"
	usr, err := getUserFromService(url)

	assert.Nil(t, err)
	assert.NotNil(t, usr)
	log.Printf(usr.ID)
}
