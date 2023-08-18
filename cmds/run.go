package cmds

import (
	"fmt"

	"github.com/mocheer/nix/cmds/types"
	"github.com/mocheer/pluto/pkg/ds/ds_json"
	"github.com/mocheer/pluto/pkg/fn"
	"github.com/mocheer/pluto/pkg/sys"
	"github.com/mocheer/pluto/pkg/ts"
	"github.com/urfave/cli/v2"
)

// Run
var Run = &cli.Command{
	Name:  "run",
	Usage: "执行脚本",
	Action: func(c *cli.Context) error {
		var conf types.PackageJSON
		err := ds_json.ReadFile("./package.json", &conf)
		if err == nil {
			scriptName := c.Args().Get(0)
			scriptContent := conf.Scripts[scriptName]
			if scriptContent != "" {
				scriptContent = fn.Format(scriptContent, ts.Map[any]{
					"appName":  conf.Name,
					"name":     conf.Name,
					"version":  conf.Version,
					"execPath": sys.GetCurrentPath(),
				})
				// 输出具体执行的内容
				fmt.Println("Powershell >", scriptContent)
				sys.Shell(scriptContent)

				// hooks
				postScriptContent := conf.Scripts[scriptName+"#post"]
				if postScriptContent != "" {
					sys.Shell(postScriptContent)
					fmt.Println("Powershell >", scriptContent)
				}

			}
		}
		return nil
	},
}
