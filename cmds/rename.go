package cmds

import (
	"github.com/urfave/cli/v2"
)

var Rename = &cli.Command{
	Name:  "rename",
	Usage: "重命名",
	Action: func(c *cli.Context) error {

		return nil
	},
}
