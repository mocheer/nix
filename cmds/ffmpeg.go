package cmds

import (
	_ "embed"

	"github.com/mocheer/pluto/pkg/sys"
	"github.com/urfave/cli/v2"
)

//go:embed tools/ffmpeg.exe
var embedFFmpeg []byte

// FFmpeg
// ffmpeg -i input.mp4 output.avi
// ffmpeg -i input.mp4 -c:v copy -c:a copy -f hls -hls_time 5 -hls_list_size 0 -hls_flags delete_segments test/output.m3u8
// ffmpeg -i input.webm -c:v libx264 -crf 25 -preset veryfast -c:a aac -b:a 128k -ac 2 -ar 44100 -f hls -hls_time 5 -hls_list_size 30 -hls_flags delete_segments -hls_segment_type mpegts test/output.m3u8
var FFmpeg = &cli.Command{
	Name:  "ffmpeg",
	Usage: "",
	Action: func(c *cli.Context) error {
		sys.MemExec(embedFFmpeg, c.Args().Slice()...)
		return nil
	},
}

// https://github.com/u2takey/ffmpeg-go
