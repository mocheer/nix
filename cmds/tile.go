package cmds

import (
	"github.com/mocheer/pluto/pkg/ts/clock"
	"github.com/mocheer/xena/pkg/tile"
	"github.com/urfave/cli/v2"
)

// Tile : nix tile xxx
// nix tile "http://map.geoq.cn/ArcGIS/rest/services/ChinaOnlineStreetPurplishBlue/MapServer/tile/{z}/{x}/{y}" data/tile-arcgis-dark
// nix tile "https://iserver.supermap.io/iserver/services/map-china400/rest/maps/ChinaDark/zxyTileImage.png?z={z}&x={x}&y={y}" data/tile-supermap-dark
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
		t := &tile.LoadConfig{
			URL:     url,
			DirName: name,
			MinZoom: 0,
			MaxZoom: 18,
		}
		t.LoadAndSave()
		return nil
	},
}
