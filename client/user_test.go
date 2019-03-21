package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/mchmarny/kuser/message"
)

func TestUserClient(t *testing.T) {

	usrIn := &message.KUser{
		ID: "test-user-ok-to-delete",
		Email: "test-user@domain.com",
		Name: "Test User",
		Created: time.Now(),
		Updated: time.Now(),
		Picture: "http://my.pic.com/123",
	}

	err := SaveUser(usrIn)
	assert.Nil(t, err)

	usrOut, err := GetUser(usrIn.ID)
	assert.Nil(t, err)
	assert.NotNil(t, usrOut)

}

func TestUserClientWithInvalidUser(t *testing.T) {

	usr, err := GetUser("badID")

	assert.NotNil(t, err)
	assert.Nil(t, usr)
}
