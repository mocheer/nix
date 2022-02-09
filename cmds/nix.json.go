package cmds

import (
	"fmt"

	"github.com/mocheer/pluto/pkg/ds/ds_text"
	"github.com/urfave/cli/v2"
)

type NixJSONConfig struct {
	Name    string            `json:"name"`
	Version string            `json:"version"`
	Scripts map[string]string `json:"scripts"`
}

// Run
var NixJSON = &cli.Command{
	Name:  "nix.json",
	Usage: "查看配置文件",
	Action: func(c *cli.Context) error {
		data, err := ds_text.Read("./nix.json")
		if err == nil {
			fmt.Println(data)
		}
		return nil
	},
}
