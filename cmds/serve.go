package cmds

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
)

// Serve : nix serve
var Serve = &cli.Command{
	Name:  "serve",
	Usage: "启动一个简易的web服务器",
	Action: func(c *cli.Context) error {
		http.Handle("/", http.FileServer(http.Dir(".")))
		fmt.Println("http://localhost:9212/")
		http.ListenAndServe(":9212", nil)
		return nil
	},
}
