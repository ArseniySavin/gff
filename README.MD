### Firs frame with ffmpeg
> This is sample and experimental application for calling ffmpeg and getting first frame.
> If you are searching really power full lib for streaming and working with ffmpeg on golang. [Use gmf](https://github.com/3d0c/gmf)

### Examples
#### Default
```go
data, err := ioutil.ReadFile("video.mp4")
	if err != nil {
		panic(err)
	}

	ffmpeg := frame.DefaultFFmpeg("ffmpeg").Build()
	reader, err :=	ffmpeg.FFrame(data)

	if err != nil {
		panic(err)
	}

	ffmpeg.WriteJpegFile("suragate-image.jpeg", reader)
```
#### Function options
```go
	data, err := ioutil.ReadFile("video.mp4")
	if err != nil {
		panic(err)
	}

	ffmpeg := frame.NewFFmpeg("ffmpeg",
		frame.WithSize(1024, 800),
		frame.WithQuality(2),
	).Build()
	reader, err :=	ffmpeg.FFrame(data)

	if err != nil {
		panic(err)
	}

	ffmpeg.WriteJpegFile("suragate-image.jpeg", reader)
```

### TODO
- to write test
- to extended examples