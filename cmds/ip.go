package cmds

import (
	"github.com/urfave/cli/v2"
)

var IP = &cli.Command{
	Name:  "ip",
	Usage: "查看当前ip信息",
	Action: func(c *cli.Context) error {
		// fmt.Println(sys_ip.GetLocalIPNet().IP)
		return nil
	},
}
