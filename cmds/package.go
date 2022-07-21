package cmds

import (
	"fmt"

	"github.com/mocheer/pluto/pkg/ds/ds_text"
	"github.com/urfave/cli/v2"
)

// Package
var Package = &cli.Command{
	Name:  "package",
	Usage: "查看配置文件",
	Action: func(c *cli.Context) error {
		data, err := ds_text.Read("./package.json")
		if err == nil {
			fmt.Println(data)
		}
		return nil
	},
}
