package ffmpeg

import (
	"bytes"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func GetVideoCover(video []byte) (cover []byte, err error) {
	videoBuf := bytes.NewBuffer(video)
	coverBuf := bytes.NewBuffer(nil)
	err = ffmpeg.Input("pipe:0").Output("pipe:1", ffmpeg.KwArgs{"vframes": 1, "format": "image2"}).WithInput(videoBuf).WithOutput(coverBuf).Run()
	if err != nil {
		return nil, err
	}
	cover = coverBuf.Bytes()
	return
}
