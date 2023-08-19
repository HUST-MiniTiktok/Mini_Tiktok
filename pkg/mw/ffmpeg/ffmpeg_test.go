package ffmpeg

import (
	"io/ioutil"
	"testing"
)

func TestGetVideoCover(t *testing.T) {
	video, err := ioutil.ReadFile("bear.mp4")
	if err != nil {
		t.Fatal(err)
	}
	cover, err := GetVideoCover(video)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("bear.jpg", cover, 0644)
	if err != nil {
		t.Fatal(err)
	}
}
