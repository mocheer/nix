package cmds

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/mocheer/pluto/pkg/sys"
	"github.com/urfave/cli/v2"
)

//
var Exec = &cli.Command{
	Name:  "exec",
	Usage: "执行脚本",
	Action: func(c *cli.Context) error {
		fmt.Println(sys.GetCurrentPath())
		scriptName := c.Args().Get(0)
		command := exec.Command(filepath.Join(`./scripts`, scriptName))
		result, err := command.Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
		return nil
	},
}
