package cmds

import (
	_ "embed"

	"github.com/mocheer/pluto/pkg/sys"
	"github.com/urfave/cli/v2"
)

//go:embed tools/upx.exe
var embedUpx []byte

// Upx
// upx -9 nix.exe
var Upx = &cli.Command{
	Name:            "upx",
	Usage:           "执行 upx 命令",
	SkipFlagParsing: true,
	Action: func(c *cli.Context) error {
		sys.MemExec(embedUpx, c.Args().Slice()...)
		return nil
	},
}
