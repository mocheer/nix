package cmds

import (
	"fmt"
	"github.com/mocheer/pluto/pkg/ds"
	"github.com/urfave/cli/v2"
	"log"
	"github.com/mocheer/pluto/pkg/ts/ctp"
)

// arcgis-query
// 一次性采集所有mapserver下的图层
// nix arcgis-mapserver-query  <服务地址> <保存路径> <token>
// nix arcgis-mapserver-query http://10.135.6.100:6080/arcgis/rest/services/FDMS/layercontrol/MapServer ./arcgis-data
// nix arcgis-mapserver-query-test 
var ArcgisMapServerQueryTest = &cli.Command{
	Name:  "arcgis-mapserver-query-test",
	Usage: "获取arcgis的query服务数据",
	Action: func(c *cli.Context) error {
		for index := range 10000 {
			u := fmt.Sprintf("http://10.135.6.100:6080/arcgis/rest/services/FDMS/base_query/MapServer/2/%d?f=json", index)
			data, err := ctp.Get(u)
			if err != nil {
				panic(err)
			}
			log.Println(u,len(data))
			ds.Save(fmt.Sprintf("./arcgis_data/小流域/%d.json", index), data)
		}
		return nil
		
	},
}
