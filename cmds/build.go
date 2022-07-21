package cmds

import (
	"github.com/mocheer/pluto/pkg/sys"
	"github.com/urfave/cli/v2"
)

//
var Build = &cli.Command{
	Name:  "build",
	Usage: "nix run build 的简写",
	Action: func(c *cli.Context) error {
		return sys.Exec("nix", "run", "build")
	},
}
