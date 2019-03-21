package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogoClient(t *testing.T) {

	logURL := "https://storage.googleapis.com/kdemo-logos/0.png"
	logo, err := GetLogoInfo(logURL)
	assert.Nil(t, err)
	assert.NotNil(t, logo)

}
