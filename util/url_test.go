package util

import (
	"log"
	"testing"
)

func TestServerSizeResizePlusPic(t *testing.T) {

	testImage := "https://lh6.googleusercontent.com/-65SFt9rUmD0/AAAAAAAAAAI/AAAAAAAB698/8pIgz0b5NG8/photo.jpg"
	expectedImg := "https://lh6.googleusercontent.com/-65SFt9rUmD0/AAAAAAAAAAI/AAAAAAAB698/8pIgz0b5NG8/s100/photo.jpg"

	sizedImg := ServerSizeResizePlusPic(testImage, 100)
	log.Printf("Sized image: %s", sizedImg)

	if sizedImg != expectedImg {
		t.Errorf("Failed to resize valid Google Plus image")
	}

}

func TestServerSizeResizeInvalidPic(t *testing.T) {

	testImage := "https://test.domain.com/someotherpic.jpg"

	sizedImg := ServerSizeResizePlusPic(testImage, 100)
	log.Printf("Sized image: %s", sizedImg)

	if testImage != sizedImg {
		t.Errorf("Resize invalid Google Plus image")
	}

}
