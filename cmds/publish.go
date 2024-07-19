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
		var conf types.PackageJSON
		err := ds_json.ReadFile("./package.json", &conf)
		if err == nil {
			sys.Exec("git", "tag", fmt.Sprintf("v%s", conf.Version))
			sys.Exec("git", "push", "--tags")
			// nix version
			sys.Exec("nix", "version")
		}
		return nil
	},
}
