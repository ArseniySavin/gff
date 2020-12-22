package frame

// FFmpegOption
type FFmpegOption func(*FFmpeg)

// WithFrame frame number
func WithFrame(frame int) FFmpegOption {
	if frame <= 0 {
		panic("frame cannot be 0 or less than 0")
	}

	return func(h *FFmpeg) {
		h.frame = frame
	}
}

// WithFrame set quality
func WithQuality(quality int) FFmpegOption {
	if 0 >= quality || quality > 10 {
		panic("quality cannot be 0 or less than 10")
	}

	return func(h *FFmpeg) {
		h.quality = quality
	}
}

// WithSize set image size
func WithSize(with, height int) FFmpegOption {
	if with == 0 || height == 0 {
		panic("size cannot be zero")
	}
	return func(h *FFmpeg) {
		h.size = FFmpegImageSize{
			With:   with,
			Height: height,
		}
	}
}

// WithFormat singlejpeg, image2
// https://ffmpeg.org/ffmpeg.html#Video-and-Audio-file-format-conversion
func WithFormat(format string) FFmpegOption {
	if format == "" {
		panic("format cannot be empty")
	}

	return func(h *FFmpeg) {
		h.format = format
	}
}
