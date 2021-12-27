package cmds

import (
	"github.com/urfave/cli/v2"
)

// git clone
var Clone = &cli.Command{
	Name:  "clone",
	Usage: "相关信息",
	Action: func(c *cli.Context) error {
		return nil
	},
}
