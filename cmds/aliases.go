package cmds

import (
	"github.com/mocheer/pluto/sys"
	"github.com/urfave/cli/v2"
)

//
var Dev = &cli.Command{
	Name:  "dev",
	Usage: "`nix run dev`的简写",
	Action: func(c *cli.Context) error {
		return sys.Exec("nix", "run", "dev")
	},
}

//
var Build = &cli.Command{
	Name:  "build",
	Usage: "nix run build的简写",
	Action: func(c *cli.Context) error {
		return sys.Exec("nix", "run", "build")
	},
}
