package cmds

import (
	"fmt"

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
			name = "3dtiles"
		}
		fmt.Println(url)
		tileset.Load(url, name)
		return nil
	},
}
