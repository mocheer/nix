package cmds

import (
	"errors"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/urfave/cli/v2"
)

// Esbuild
var Esbuild = &cli.Command{
	Name:  "esbuild",
	Usage: "相关信息",
	Action: func(c *cli.Context) error {
		args := c.Args()
		url := args.Get(0)

		result := api.Build(api.BuildOptions{
			// 输入
			EntryPoints: []string{url},
			// Target:      api.ES2015,
			// 压缩
			MinifyWhitespace:  true,
			MinifyIdentifiers: true,
			MinifySyntax:      true,
			// 输出
			Outfile: "out.js",
			Write:   true,
		})

		if len(result.Errors) > 0 {
			return errors.New(result.Errors[0].Text)
		}
		return nil
	},
}
