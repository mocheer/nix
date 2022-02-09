package cmds

import (
	"github.com/mocheer/pluto/pkg/ts/clock"
	"github.com/mocheer/xena/pkg/tileset"
	"github.com/urfave/cli/v2"
)

// Tileset : nix tileset xxx/tileset.json
var Tileset = &cli.Command{
	Name:  "tileset",
	Usage: "下载在线3dtiles模型数据",
	Action: func(c *cli.Context) error {
		args := c.Args()
		url := args.Get(0)
		name := args.Get(1) //保存的路径
		if name == "" {
			name = "data/3dtiles-" + clock.Now().Fmt(clock.FmtCompactFullDate)
		}
		tileset.Load(url, name)
		return nil
	},
}
