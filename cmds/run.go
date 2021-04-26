package cmds

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mocheer/pluto/fn"
	"github.com/mocheer/pluto/fs"
	"github.com/mocheer/pluto/reg"
	"github.com/mocheer/pluto/sys"
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
		err := fs.ReadJSON("./nix.json", &conf)
		if err == nil {
			scriptName := c.Args().Get(0)
			scriptContent := conf.Scripts[scriptName]
			if scriptContent != "" {
				scriptContent = fn.FmtString(scriptContent, map[string]interface{}{
					"appName":  conf.Name,
					"execPath": sys.GetAbsolutePath(),
				})
				fmt.Println(scriptContent)

				scriptArray := strings.Split(scriptContent, "&&")
				for _, script := range scriptArray {
					args := regexp.MustCompile(reg.CommandParams).FindAllString(script, 10)
					params := []string{}
					// 命令行参数在窗口输入的时候需要带引号，但这里的参数不需要，反而要去掉
					for _, val := range args[1:] {
						params = append(params, strings.ReplaceAll(val, `"`, ``))
					}
					command := exec.Command(args[0], params...)
					err := command.Run()
					if err != nil {
						fmt.Println(err)
					}
				}

			}
		}
		return nil
	},
}
