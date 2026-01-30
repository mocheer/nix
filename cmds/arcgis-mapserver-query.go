package cmds

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/urfave/cli/v2"
)

// arcgis-query
// 一次性采集所有mapserver下的图层
// nix arcgis-mapserver-query  <服务地址> <保存路径> <token>
// nix arcgis-mapserver-query http://10.135.6.100:6080/arcgis/rest/services/FDMS/layercontrol/MapServer ./arcgis-data
var ArcgisMapServerQuery = &cli.Command{
	Name:  "arcgis-mapserver-query",
	Usage: "获取arcgis的query服务数据",
	Action: func(c *cli.Context) error {

		url := c.Args().Get(0)
		fname := c.Args().Get(1)
		token := c.Args().Get(2)
		//
		resp, err := http.Get(url + "?f=json")

		if err != nil {
			log.Println("请求错误,请检查服务是否正常运行")
		} else {
			log.Println("请求成功", url)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		var data ArcgisMapServerMeta
		if err := json.Unmarshal(body, &data); err != nil {
			panic(err)
		}
		log.Println(len(data.Layers))
		for _, layer := range data.Layers {
			if len(layer.SubLayerIds) > 0 {
				log.Println(layer.Name, len(layer.SubLayerIds))
				continue
			}
			layerFName := fname + "/" + layer.Name
			queryURL := fmt.Sprintf("%s/%d"+"/query", url, layer.ID)
			log.Println(queryURL)
			if ds.IsExist(layerFName) {
				continue
			}
			go reptile(queryURL, layerFName, token, true)
		}
		select {}
	},
}

type ArcgisMapServerMeta struct {
	CurrentVersion            float64          `json:"currentVersion"`
	ServiceDescription        string           `json:"serviceDescription"`
	MapName                   string           `json:"mapName"`
	Description               string           `json:"description"`
	CopyrightText             string           `json:"copyrightText"`
	SupportsDynamicLayers     bool             `json:"supportsDynamicLayers"`
	Layers                    []Layers         `json:"layers"`
	Tables                    []any            `json:"tables"`
	SpatialReference          SpatialReference `json:"spatialReference"`
	SingleFusedMapCache       bool             `json:"singleFusedMapCache"`
	InitialExtent             InitialExtent    `json:"initialExtent"`
	FullExtent                FullExtent       `json:"fullExtent"`
	MinScale                  int              `json:"minScale"`
	MaxScale                  int              `json:"maxScale"`
	Units                     string           `json:"units"`
	SupportedImageFormatTypes string           `json:"supportedImageFormatTypes"`
	DocumentInfo              DocumentInfo     `json:"documentInfo"`
	Capabilities              string           `json:"capabilities"`
	SupportedQueryFormats     string           `json:"supportedQueryFormats"`
	ExportTilesAllowed        bool             `json:"exportTilesAllowed"`
	MaxRecordCount            int              `json:"maxRecordCount"`
	MaxImageHeight            int              `json:"maxImageHeight"`
	MaxImageWidth             int              `json:"maxImageWidth"`
	SupportedExtensions       string           `json:"supportedExtensions"`
}
type Layers struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	ParentLayerID     int    `json:"parentLayerId"`
	DefaultVisibility bool   `json:"defaultVisibility"`
	SubLayerIds       []int  `json:"subLayerIds"`
	MinScale          int    `json:"minScale"`
	MaxScale          int    `json:"maxScale"`
}
type SpatialReference struct {
	Wkid       int `json:"wkid"`
	LatestWkid int `json:"latestWkid"`
}
type InitialExtentSpatialReference struct {
	Wkid       int `json:"wkid"`
	LatestWkid int `json:"latestWkid"`
}
type InitialExtent struct {
	Xmin                          float64                       `json:"xmin"`
	Ymin                          float64                       `json:"ymin"`
	Xmax                          float64                       `json:"xmax"`
	Ymax                          float64                       `json:"ymax"`
	InitialExtentSpatialReference InitialExtentSpatialReference `json:"spatialReference"`
}
type FullExtentSpatialReference struct {
	Wkid       int `json:"wkid"`
	LatestWkid int `json:"latestWkid"`
}
type FullExtent struct {
	Xmin                       float64                    `json:"xmin"`
	Ymin                       float64                    `json:"ymin"`
	Xmax                       float64                    `json:"xmax"`
	Ymax                       float64                    `json:"ymax"`
	FullExtentSpatialReference FullExtentSpatialReference `json:"spatialReference"`
}
type DocumentInfo struct {
	Title                string `json:"Title"`
	Author               string `json:"Author"`
	Comments             string `json:"Comments"`
	Subject              string `json:"Subject"`
	Category             string `json:"Category"`
	AntialiasingMode     string `json:"AntialiasingMode"`
	TextAntialiasingMode string `json:"TextAntialiasingMode"`
	Keywords             string `json:"Keywords"`
}
