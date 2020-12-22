package main

import (
	"gff/frame"
	"io/ioutil"
)

func main() {

	b, err := ioutil.ReadFile("test.mp4")
	if err != nil {
		panic(err)
	}

	ffmpeg := frame.DefaultFFmpeg("ffmpeg").Build()
	reader, err := ffmpeg.FFrame(b)

	if err != nil {
		panic(err)
	}

	ffmpeg.WriteJpegFile("suragate-image.jpeg", reader)
}

func Options() {

	b, err := ioutil.ReadFile("test.mp4")
	if err != nil {
		panic(err)
	}

	ffmpeg := frame.NewFFmpeg("ffmpeg",
		frame.WithSize(1024, 800),
		frame.WithQuality(2),
	).Build()
	reader, err := ffmpeg.FFrame(b)

	if err != nil {
		panic(err)
	}

	ffmpeg.WriteJpegFile("suragate-image.jpeg", reader)
}
