package cmds

import (
	"path"

	"github.com/mocheer/nix/global"
	"github.com/mocheer/pluto/ec"
	"github.com/urfave/cli/v2"
)

// nix rsa 生成密钥文件
var Rsa = &cli.Command{
	Name:  "rsa",
	Usage: "生成rsa文件",
	Action: func(c *cli.Context) error {
		dir := c.Args().Get(0)
		if dir == "" {
			dir = path.Join(global.ExportDir, "rsa")
		}
		ec.RSA_GenPemFiles(dir, 2048)
		return nil
	},
}
