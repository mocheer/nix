package cmds

import (
	"fmt"

	"github.com/mocheer/pluto/pkg/ds/ds_tif"
	"github.com/urfave/cli/v2"
)

// tif
// nix tif ./data/25S_20200101-20210101.tif
var Tif = &cli.Command{
	Name:  "tif",
	Usage: "查看tif文件标签属性",
	Action: func(c *cli.Context) error {
		args := c.Args()
		fileName := args.Get(0)
		fmt.Println(fileName)
		t, _ := ds_tif.ReadFile(fileName)
		fmt.Println(t.Tif.IFDs())
		return nil
	},
}
