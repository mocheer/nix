package cmds

import (
	"fmt"

	"github.com/mocheer/pluto/ds/ds_json"
	"github.com/mocheer/pluto/fn"
	"github.com/mocheer/pluto/sys"
	"github.com/mocheer/pluto/ts"
	"github.com/urfave/cli/v2"
)

type Config struct {
	Name    string
	Scripts map[string]string
}

// Run
var Run = &cli.Command{
	Name:  "run",
	Usage: "执行脚本",
	Action: func(c *cli.Context) error {
		var conf Config
		err := ds_json.Read("./nix.json", &conf)
		if err == nil {
			scriptName := c.Args().Get(0)
			scriptContent := conf.Scripts[scriptName]
			if scriptContent != "" {
				scriptContent = fn.FmtString(scriptContent, ts.Map{
					"appName":  conf.Name,
					"execPath": sys.GetCurrentPath(),
				})
				// 输出具体执行的内容
				fmt.Println("Powershell >", scriptContent)
				sys.Shell(scriptContent)
			}
		}
		return nil
	},
}
