package cmds

import (
	"github.com/mocheer/nix/cmds/types"
	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ds/ds_json"
	"github.com/mocheer/pluto/pkg/sys"
	"github.com/urfave/cli/v2"
)

//
var Init = &cli.Command{
	Name:  "init",
	Usage: "`初始化",
	Action: func(c *cli.Context) error {
		if ds.IsExist("package.json") {

		} else {

			ds_json.Save("package.json", &types.PackageJSON{
				Name:    sys.GetCurrentDirname(),
				Version: "1.0.0",
				Author:  "mocheer",
				License: "MIT",
				Scripts: map[string]string{},
			})
		}
		return nil
	},
}
