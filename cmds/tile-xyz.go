package cmds

import (
	"strings"

	"github.com/mocheer/pluto/pkg/ts/clock"
	"github.com/mocheer/xena/pkg/tile"
	"github.com/urfave/cli/v2"
)

// TileXYZ : nix tile-xyz xxx
// nix tile-xyz "http://webrd0{s}.is.autonavi.com/appmaptile?lang=zh_cn&style=8&x={x}&y={y}&z={z}" public/GaodeMap.Normal
// nix tile-xyz "http://wprd0{s}.is.autonavi.com/appmaptile?x={x}&y={y}&z={z}&lang=zh_cn&size=1&scl=2&style=6" public/GaodeMap.Satellite
// nix tile-xyz "http://wprd0{s}.is.autonavi.com/appmaptile?x={x}&y={y}&z={z}&lang=zh_cn&size=1&scl=1&style=8" public/GaodeMap.Satellite_A
// nix tile-xyz "http://t{s}.tianditu.gov.cn/DataServer?T=vec_c&x={x}&y={y}&l={z}&tk=070a93160eddd5f891599e51a6b764ac" public/Tianditu.Normal
// nix tile-xyz "http://t{s}.tianditu.gov.cn/DataServer?T=cva_w&x={x}&y={y}&l={z}&tk=070a93160eddd5f891599e51a6b764ac" public/Tianditu.Normal_A
// nix tile-xyz "http://t{s}.tianditu.gov.cn/DataServer?T=img_c&x={x}&y={y}&l={z}&tk=60714a8ef3e4491df43827cd34c2aa22" public/Tianditu.Satellite
// nix tile-xyz "http://t{s}.tianditu.gov.cn/DataServer?T=cia_w&x={x}&y={y}&l={z}&tk=60714a8ef3e4491df43827cd34c2aa22" public/Tianditu.Satellite_A
// nix tile-xyz "https://tile{s}.tianditu.gov.cn/vts?t=vt&z={z}&x={x}&y={y}&tk=34568012b0e7be57119fa5124bd7bdd6"  public/Tianditu.Normal_VT
var TileXYZ = &cli.Command{
	Name:  "tile-xyz",
	Usage: "下载在线瓦片地图",
	Action: func(c *cli.Context) error {
		args := c.Args()
		url := args.Get(0)
		name := args.Get(1) //保存的路径

		if name == "" {
			name = "public/tile-" + clock.Now().Fmt(clock.FmtCompactFullDate)
		}
		c2 := &tile.LoadConfig{
			URL:        url,
			DirName:    name,
			MinZoom:    0,
			MaxZoom:    18,
			Subdomains: strings.Split("123", ""),
		}
		c2.LoadAndSave()
		return nil
	},
}
