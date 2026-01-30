package cmds

import (
	"fmt"
	"path/filepath"

	"github.com/mocheer/pluto/pkg/ts/clock"
	"github.com/mocheer/xena/pkg/tileset"
	"github.com/urfave/cli/v2"
)

// Tileset : nix tileset xxx/tileset.json
// nix tileset http//122.112.175.6:8103/3dtile/tileset.json
// nix tileset https://www.thingjs.com/static/tilesData/tileset.json
// nix tileset https://tile.googleapis.com/v1/3dtiles/root.json?key=AIzaSyCnRPXWDIj1LuX6OWIweIqZFHHoXVgdYss googlephotorealistic3dtileset
// nix tileset http://data1.mars3d.cn/3dtiles/qx-hfdxy/tileset.json mars3d_video3d
// nix tileset http://112.81.89.197:9900/data/3dtiles/tileset.json data/3dtiles-20240708143021
// 以下是变种，龙文区九十九湾
// nix tileset http://172.16.2.52/models-rest/rest/models/preview/LW99DG1/tileset.json data/lwq99w1
// nix tileset http://172.16.2.52/models-rest/rest/models/preview/LW99DS2/tileset.json data/lwq99w2
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
		name = filepath.Join(name, "tileset.json")
		err := tileset.Load(url, name)
		fmt.Println(err)
		return err
	},
}
