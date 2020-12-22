package frame

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"os/exec"
)

type FFmpeg struct {
	path    string
	frame   int
	quality int
	size    FFmpegImageSize
	format  string
	command *exec.Cmd
}

type Command struct {
	ffmpeg *FFmpeg
}

type FFmpegImageSize struct {
	With   int
	Height int
}

// DefaultFFmpeg set default options
func DefaultFFmpeg(path string) *FFmpeg {
	if path == "" {
		panic("format cannot be empty")
	}

	return &FFmpeg{
		path:    path,
		frame:   1,
		quality: 2,
		size: FFmpegImageSize{
			With:   640,
			Height: 360,
		},
		format: "singlejpeg",
	}
}

// NewFFmpeg set custom
func NewFFmpeg(path string, opts ...FFmpegOption) *FFmpeg {
	if path == "" {
		panic("path cannot be empty")
	}

	ffmpeg := DefaultFFmpeg(path)

	for _, opt := range opts {
		opt(ffmpeg)
	}

	return ffmpeg
}

func (c *FFmpeg) Build() *Command {
	c.command = exec.Command(c.path,
		"-i", "-", // to read from stdin
		"-vframes", fmt.Sprint(c.frame),
		"-s", fmt.Sprintf("%dx%d", c.size.With, c.size.Height),
		"-q:v", fmt.Sprint(c.quality),
		"-f", c.format,
		"-", // to read from stdout
	) // TODO make builder

	return &Command{ffmpeg: c}
}

// String get command string
func (c *FFmpeg) String() {
	c.command.String()
}

func (c *Command) FFrame(data []byte) (*bufio.Reader, error) {
	c.ffmpeg.command.Stdin = bytes.NewBuffer(data)

	var buffer bytes.Buffer
	c.ffmpeg.command.Stdout = &buffer

	if c.ffmpeg.command.Run() != nil {
		return nil, fmt.Errorf("Cannot run or found ffmpeg file by path: %s", c.ffmpeg.path)
	}

	reader := bufio.NewReader(&buffer)
	return reader, nil
}

func (c *Command) WriteJpegFile(path string, reader *bufio.Reader) error {
	i, err := jpeg.Decode(reader)

	if err != nil {
		return err
	}

	file, err := os.Create(path)
	defer file.Close()

	if err != nil {
		return err
	}

	return jpeg.Encode(file, i, &jpeg.Options{Quality: 100})
}

func (c *Command) Jpeg(path string, reader *bufio.Reader) (image.Image, error) {
	i, err := png.Decode(reader)

	if err != nil {
		return nil, err
	}

	return i, nil
}
