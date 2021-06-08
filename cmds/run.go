package cmds

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mocheer/pluto/fn"
	"github.com/mocheer/pluto/fs"
	"github.com/mocheer/pluto/reg"
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
		err := fs.ReadJSON("./nix.json", &conf)
		if err == nil {
			scriptName := c.Args().Get(0)
			scriptContent := conf.Scripts[scriptName]
			if scriptContent != "" {
				scriptContent = fn.FmtString(scriptContent, ts.Map{
					"appName":  conf.Name,
					"execPath": sys.GetAbsolutePath(),
				})
				// 输出具体执行的内容
				fmt.Println(scriptContent)
				// 多个脚本命令
				scriptArray := strings.Split(scriptContent, "&&")
				for _, script := range scriptArray {
					// 空格分割
					args := regexp.MustCompile(reg.CommandParams).FindAllString(script, 10)
					// 参数集合
					params := []string{}
					name := args[0]
					switch name {
					case "nssm":
						args = append([]string{"nix"}, args...)
					}
					// 命令行参数在窗口输入的时候需要带引号，但这里的参数不需要，反而要去掉
					for _, val := range args[1:] {
						params = append(params, strings.ReplaceAll(val, `"`, ``))
					}
					//
					err := sys.Exec(args[0], params...)
					if err != nil {
						fmt.Println(err)
					}
				}

			}
		}
		return nil
	},
}
