package cmds

import (
	"github.com/mocheer/pluto/pkg/ts/clock"
	"github.com/mocheer/xena/pkg/tile"
	"github.com/urfave/cli/v2"
)

// TileXYZ : nix tile-xyz xxx
// nix tile-xyz "http://webrd03.is.autonavi.com/appmaptile?lang=zh_cn&style=8&x={x}&y={y}&z={z}" data/GaodeMap.Normal
// nix tile-xyz "http://wprd02.is.autonavi.com/appmaptile?x={x}&y={y}&z={z}&lang=zh_cn&size=1&scl=2&style=6" data/GaodeMap.Satellite
// nix tile-xyz "http://wprd01.is.autonavi.com/appmaptile?x={x}&y={y}&z={z}&lang=zh_cn&size=1&scl=1&style=8" data/GaodeMap.Satellite_A
// nix tile-xyz "http://t5.tianditu.gov.cn/DataServer?T=vec_c&x={x}&y={y}&l={z}&tk=16554181e1d8f9f3b82ce84fe953c164" data/Tianditu.Normal
// nix tile-xyz "http://t1.tianditu.gov.cn/DataServer?T=img_c&x={x}&y={y}&l={z}&tk=16554181e1d8f9f3b82ce84fe953c164" data/Tianditu.Satellite
var TileXYZ = &cli.Command{
	Name:  "tile-xyz",
	Usage: "下载在线瓦片地图",
	Action: func(c *cli.Context) error {
		args := c.Args()
		url := args.Get(0)
		name := args.Get(1) //保存的路径
		if name == "" {
			name = "data/tile-" + clock.Now().Fmt(clock.FmtCompactFullDate)
		}
		tile.Load(&tile.LoadConfig{
			URL:     url,
			DirName: name,
			MinZoom: 0,
			MaxZoom: 18,
		})
		return nil
	},
}
