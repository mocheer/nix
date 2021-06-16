package cmds

import (
	"github.com/mocheer/pluto/ts/clock"
	"github.com/mocheer/pluto/ts/tileset"
	"github.com/urfave/cli/v2"
)

// Tileset : nix tileset xxx/tileset.json
var Tileset = &cli.Command{
	Name:  "tileset",
	Usage: "远程下载第三方3dtiles模型数据",
	Action: func(c *cli.Context) error {
		args := c.Args()
		url := args.Get(0)
		name := args.Get(1)
		if name == "" {
			name = "data/3dtiles-" + clock.Now().Fmt(clock.FmtCompactFullDate)
		}
		tileset.Load(url, name)
		return nil
	},
}
