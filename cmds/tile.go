package cmds

import (
	"github.com/mocheer/pluto/ts/clock"
	"github.com/mocheer/xena/gfs/tile"
	"github.com/urfave/cli/v2"
)

// Tile : nix tile xxx
// nix tile http://map.geoq.cn/ArcGIS/rest/services/ChinaOnlineStreetPurplishBlue/MapServer/tile/ data/tile-dark
var Tile = &cli.Command{
	Name:  "tile",
	Usage: "下载在线瓦片地图",
	Action: func(c *cli.Context) error {
		args := c.Args()
		url := args.Get(0)
		name := args.Get(1) //保存的路径
		if name == "" {
			name = "data/tile-" + clock.Now().Fmt(clock.FmtCompactFullDate)
		}
		tile.Load(url, name)

		return nil
	},
}
