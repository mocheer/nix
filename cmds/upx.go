package cmds

import (
	_ "embed"
	"path"

	"github.com/mocheer/nix/global"
	"github.com/mocheer/pluto/fs"
	"github.com/mocheer/pluto/sys"
	"github.com/urfave/cli/v2"
)

//go:embed tools/upx/upx.exe
var upxBytes []byte

// Upx
// upx -9 nix.exe
var Upx = &cli.Command{
	Name:            "upx",
	Usage:           "执行 upx 命令",
	SkipFlagParsing: true,
	Action: func(c *cli.Context) error {
		upxExePath := path.Join(global.ExportDir, "upx.exe")
		isExist := fs.IsExist(upxExePath)
		if !isExist {
			fs.Save(upxExePath, upxBytes)
		}
		sys.Exec(upxExePath, c.Args().Slice()...)
		return nil
	},
}
