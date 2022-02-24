package cmds

import (
	"fmt"

	"github.com/mocheer/pluto/pkg/ts/clock"
	"github.com/mocheer/xena/pkg/tileset"
	"github.com/urfave/cli/v2"
)

// Tileset : nix tileset xxx/tileset.json
// nix tileset  http//122.112.175.6:8103/3dtile/tileset.json
var Tileset = &cli.Command{
	Name:  "tileset",
	Usage: "下载在线3dtiles模型数据",
	Action: func(c *cli.Context) error {

		args := c.Args()
		url := args.Get(0)
		name := args.Get(1) //保存的路径
		fmt.Println(url)
		if name == "" {
			name = "data/3dtiles-" + clock.Now().Fmt(clock.FmtCompactFullDate)
		}
		err := tileset.Load(url, name)
		fmt.Println(err)
		return nil
	},
}
