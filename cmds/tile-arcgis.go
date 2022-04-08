package cmds

import (
	"github.com/mocheer/pluto/pkg/ts/clock"
	"github.com/mocheer/xena/pkg/tile"
	"github.com/urfave/cli/v2"
)

// Tile : nix tile xxx
// nix tile-arcgis "./dmap-arcgis/dq" test/tile-arcgis-dark
var TileArcgis = &cli.Command{
	Name:  "tile-arcgis",
	Usage: "下载arcgis本地瓦片地图",
	Action: func(c *cli.Context) error {
		args := c.Args()
		url := args.Get(0)
		name := args.Get(1) //保存的路径
		if name == "" {
			name = "data/tile-" + clock.Now().Fmt(clock.FmtCompactFullDate)
		}
		return tile.LoadArcgis(url, name)
	},
}
