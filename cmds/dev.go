package cmds

import (
	"github.com/mocheer/pluto/pkg/sys"
	"github.com/urfave/cli/v2"
)

//
var Dev = &cli.Command{
	Name:  "dev",
	Usage: "nix run dev įįŽå",
	Action: func(c *cli.Context) error {
		return sys.Exec("nix", "run", "dev")
	},
}
