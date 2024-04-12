package cmds

import (
	"strconv"
	"strings"

	"github.com/mocheer/pluto/pkg/ds/ds_json"
	"github.com/mocheer/pluto/pkg/fn"
	"github.com/urfave/cli/v2"

	"github.com/mocheer/nix/cmds/types"
)

var Version = &cli.Command{
	Name:  "version",
	Usage: "修改版本号",
	Action: func(c *cli.Context) error {
		var conf types.PackageJSON
		err := ds_json.ReadFile("./package.json", &conf)
		if err == nil {
			versions := strings.Split(conf.Version, ".")
			pathchVersion := versions[len(versions)-1]
			pathchVersionInt := fn.ParseInt(pathchVersion) + 1
			versions[len(versions)-1] = strconv.Itoa(pathchVersionInt)
			conf.Version = strings.Join(versions, ".")
			ds_json.Save("./package.json", conf)
		}
		return nil
	},
}
