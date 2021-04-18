package cmds

import (
	"github.com/mocheer/pluto/fs"
	"github.com/urfave/cli/v2"
)

// Fs 文件操作命令集
var Fs = &cli.Command{
	Name:  "fs",
	Usage: "文件操作命令集",
	Subcommands: []*cli.Command{
		// nix fs append -d b -s "import * as cesium from 'cesium'" -suffix text
		{
			Name:  "append",
			Usage: "往文件追加内容",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "suffix"},
				&cli.StringFlag{Name: "dir", Aliases: []string{"d"}},
				&cli.StringFlag{Name: "start", Aliases: []string{"s"}},
			},
			Action: func(c *cli.Context) error {
				fs.EachDirAppendHead(c.String("d"), c.String("s"), map[string]interface{}{
					"suffix": c.String("suffix"),
				})
				return nil
			},
		},
	},
}
