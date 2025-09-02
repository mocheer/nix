package cmds

import (
	"fmt"

	"github.com/mocheer/nix/cmds/types"
	"github.com/mocheer/pluto/pkg/ds/ds_json"
	"github.com/mocheer/pluto/pkg/sys"
	"github.com/urfave/cli/v2"
)

// Publish
// nix publish
// git tag v1.0.0
// git push --tags
var Publish = &cli.Command{
	Name:  "publish",
	Usage: "发布tag版本",
	Action: func(c *cli.Context) error {
		tagName := c.Args().Get(0)
		if tagName != "" {
			sys.Exec("git", "tag", tagName, "-f")
			sys.Exec("git", "push", "--tags")
			return nil
		}

		var conf types.PackageJSON
		err := ds_json.ReadFile("./package.json", &conf)
		if err == nil {
			sys.Exec("git", "tag", fmt.Sprintf("v%s", conf.Version))
			// 以下推送所有本地标签，如果要推送单个标签理应使用：git push origin v1.0.0
			// Git 默认只会将标签推送到与本地仓库关联的主远程仓库（通常是 origin），如果要推送指定仓库，使用git push gitlab --tags
			sys.Exec("git", "push", "--tags")
			// nix version
			sys.Exec("nix", "version")
		}
		return nil
	},
}
