package cmds

import (
	"github.com/mocheer/pluto/pkg/sys"
	"github.com/urfave/cli/v2"
)

// go get -u all
var Install = &cli.Command{
	Name:  "install",
	Usage: "`升级安装所有依赖包",
	Action: func(c *cli.Context) error {
		return sys.Exec("go", "get", "-u")
	},
}
