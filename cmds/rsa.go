package cmds

import (
	"path"

	"github.com/mocheer/nix/global"
	"github.com/mocheer/pluto/ecc"
	"github.com/urfave/cli/v2"
)

//
var Rsa = &cli.Command{
	Name:  "rsa",
	Usage: "生成rsa文件",
	Action: func(c *cli.Context) error {
		dir := c.Args().Get(0)
		if dir == "" {
			dir = path.Join(global.ExportDir, "rsa")
		}
		ecc.RSA_GenPemFiles(dir, 2048)
		return nil
	},
}
