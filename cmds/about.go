package cmds

import (
	"fmt"

	"github.com/mocheer/pluto/sys"
	"github.com/urfave/cli/v2"
)

//
var About = &cli.Command{
	Name:  "about",
	Usage: "相关信息",
	Action: func(c *cli.Context) error {
		fmt.Println(sys.GetCurrentPath())
		return nil
	},
}
