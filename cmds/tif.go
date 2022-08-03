package cmds

import (
	"fmt"

	"github.com/mocheer/xena/pkg/gtif"
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
		t, _ := gtif.Read(fileName)
		fmt.Println(t.Tif.IFDs())
		return nil
	},
}
