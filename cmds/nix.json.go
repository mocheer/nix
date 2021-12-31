package cmds

import (
	"fmt"

	"github.com/mocheer/pluto/ds/ds_json"
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
		var conf NixJSONConfig
		err := ds_json.Read("./nix.json", &conf)
		if err == nil {
			fmt.Printf("%+v", conf)
		}
		return nil
	},
}
