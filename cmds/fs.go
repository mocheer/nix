package cmds

import (
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/urfave/cli/v2"
)

// Fs 文件操作命令集
var Fs = &cli.Command{
	Name:  "fs",
	Usage: "文件操作命令集",
	Subcommands: []*cli.Command{
		// nix fs append -d ./dir -s "import * as cesium from 'cesium'" -suffix text
		{
			Name:  "append",
			Usage: "批量往文件追加内容",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "suffix"},
				&cli.StringFlag{Name: "dir", Aliases: []string{"d"}},
				&cli.StringFlag{Name: "start", Aliases: []string{"s"}},
			},
			Action: func(c *cli.Context) error {
				ds.EachFilesToAppendHead(c.String("d"), c.String("s"), map[string]any{
					"suffix": c.String("suffix"),
				})
				return nil
			},
		},
		// nix fs rename -d ./dir
		{
			Name:  "rename",
			Usage: "批量重命名文件",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "dir", Aliases: []string{"d"}},
			},
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		// nix fs split -n file.ext -cs 1024
		{
			Name:  "split",
			Usage: "切割大文件",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
				&cli.StringFlag{Name: "chunkSize", Aliases: []string{"cs"}},
			},
			Action: func(c *cli.Context) error {
				chunkSize := c.Int("cs")
				if chunkSize == 0 {
					chunkSize = 1024 * 1000 * 1000 * 16 // 16G
				}
				return ds.SplitFile(c.String("name"), chunkSize)
			},
		},
		// nix fs merge -d file.ext
		{
			Name:  "merge",
			Usage: "合并被切割的大文件",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "dir", Aliases: []string{"d"}},
			},
			Action: func(c *cli.Context) error {
				name := c.String("d")
				files, err := ds.GetFiles(name)
				slices.SortFunc(files, func(a string, b string) int {
					ai, _ := strconv.Atoi(filepath.Base(a))
					bi, _ := strconv.Atoi(filepath.Base(b))
					return ai - bi
				})
				if err != nil {
					return err
				}
				return ds.MergeFiles(files, strings.Split(filepath.Base(name), "_")[0])
			},
		},
	},
}
