package cmds

import (
	"log"
	"net/http"
	"net/url"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/urfave/cli/v2"
)

// Serve : nix serve
// nix serve -p=9912
// 可用于搭配爬虫服务中下载所有网站资源的功能，api也是一种资源，这里解决api重定向的问题
var Serve = &cli.Command{
	Name:  "serve",
	Usage: "启动一个简易的web服务器",
	Action: func(c *cli.Context) error {
		handle := http.FileServer(http.Dir("."))
		http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 检查URL中是否有查询参数
			if r.URL.RawQuery != "" {
				name := r.URL.Path + "@" + url.QueryEscape(r.URL.RawQuery)
				if name[0] == '/' {
					name = "." + name
				}
				if ds.IsExist(name) {
					log.Println(name)
					http.ServeFile(w, r, name)
					return
				}
				// 有查询参数，重定向到指定文件 //
				name = r.URL.Path + url.QueryEscape(r.URL.RawQuery)
				if name[0] == '/' {
					name = "." + name
				}
				if ds.IsExist(name) {
					log.Println(name)
					http.ServeFile(w, r, name)
					return
				}
				log.Println(name, r.URL.Path, r.URL.RawQuery)
				if ds.IsExist(r.URL.Path) {
					http.ServeFile(w, r, r.URL.Path)
					return
				}
				http.NotFound(w, r)
				return
			}
			// 没有查询参数，使用文件服务器处理请求
			handle.ServeHTTP(w, r)
		}))
		//
		log.Println("http://localhost:9212/")
		//
		http.ListenAndServe(":9212", nil)
		return nil
	},
}
