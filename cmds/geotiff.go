package cmds

import (
	"errors"
	"fmt"

	"github.com/mocheer/xena/gfs/raster"
	"github.com/urfave/cli/v2"
)

// GeoTiff
// nix geotiff ./data/25S_20200101-20210101.tif
var GeoTiff = &cli.Command{
	Name:  "geotiff",
	Usage: "读取geotiff文件",
	Action: func(c *cli.Context) error {
		args := c.Args()
		inFile := args.Get(0)

		rin, err := raster.CreateRasterFromFile(inFile)
		if err != nil {
			return err
		}
		fmt.Println(rin)
		tagInfo := rin.GetMetadataEntries()
		if len(tagInfo) > 0 {
			fmt.Println(tagInfo[0])
		} else {
			return errors.New("error reading metadata entries")
		}
		return nil
	},
}
