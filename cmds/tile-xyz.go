package cmds

import (
	"strings"

	"github.com/mocheer/pluto/pkg/ts/clock"
	"github.com/mocheer/xena/pkg/tile"
	"github.com/urfave/cli/v2"
)

// TileXYZ : nix tile-xyz xxx
// nix tile-xyz "http://webrd0{s}.is.autonavi.com/appmaptile?lang=zh_cn&style=8&x={x}&y={y}&z={z}" data/GaodeMap.Normal
// nix tile-xyz "http://wprd0{s}.is.autonavi.com/appmaptile?x={x}&y={y}&z={z}&lang=zh_cn&size=1&scl=2&style=6" data/GaodeMap.Satellite
// nix tile-xyz "http://wprd0{s}.is.autonavi.com/appmaptile?x={x}&y={y}&z={z}&lang=zh_cn&size=1&scl=1&style=8" data/GaodeMap.Satellite_A
// nix tile-xyz "http://t{s}.tianditu.gov.cn/DataServer?T=vec_c&x={x}&y={y}&l={z}&tk=070a93160eddd5f891599e51a6b764ac" data/Tianditu.Normal
// nix tile-xyz "http://t{s}.tianditu.gov.cn/DataServer?T=cva_w&x={x}&y={y}&l={z}&tk=070a93160eddd5f891599e51a6b764ac" data/Tianditu.Normal_A
// nix tile-xyz "http://t{s}.tianditu.gov.cn/DataServer?T=img_c&x={x}&y={y}&l={z}&tk=60714a8ef3e4491df43827cd34c2aa22" data/Tianditu.Satellite
// nix tile-xyz "http://t{s}.tianditu.gov.cn/DataServer?T=cia_w&x={x}&y={y}&l={z}&tk=60714a8ef3e4491df43827cd34c2aa22" data/Tianditu.Satellite_A
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
			URL:        url,
			DirName:    name,
			MinZoom:    14,
			MaxZoom:    18,
			Subdomains: strings.Split("1234", ""),
		})
		return nil
	},
}
