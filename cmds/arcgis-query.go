package cmds

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/mocheer/pluto/pkg/ds/ds_json"
	"github.com/mocheer/pluto/pkg/ts/clock"

	"github.com/urfave/cli/v2"
)

// arcgis-query
var ArcgisQuery = &cli.Command{
	Name:  "arcgis-query",
	Usage: "获取arcgis的query服务数据",
	Action: func(c *cli.Context) error {

		url := c.Args().Get(0)
		if url == "" {
			serviceURL := "http://10.135.210.71:6116/ags-proxy/server1/rest/services/fj_floodrisk/fj_floodrisk_map/MapServer/"
			token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhY2Nlc3NJZCI6IjE2MTA5MjQzMDczOTMwODEzNDYiLCJ1c2VyTmFtZSI6InNjMjAyNCIsImV4cCI6MTc1MjE5OTA5MX0.Bd2RuqJOQjbQBN2WHE73VYNp-AAGom37ANRGq9PbRL8"
			go reptile(serviceURL+"0/query", "断面.json", token)
			go reptile(serviceURL+"1/query", "5y.json", token)
			go reptile(serviceURL+"2/query", "10y.json", token)
			go reptile(serviceURL+"3/query", "20y.json", token)
			go reptile(serviceURL+"4/query", "50y.json", token)
			go reptile(serviceURL+"5/query", "100y.json", token)
			go reptile(serviceURL+"6/query", "200y.json", token)
			select {}
		} else {
			fname := c.Args().Get(1)
			if fname == "" {
				fname = clock.Now().Fmt(clock.FmtCompactFullDate) + ".json"
			}
			token := c.Args().Get(2)
			reptile(url, fname, token)
		}
		return nil
	},
}

const (
	batchSize   = 2000  // 每批请求数量（不能超过服务端设置的最大值）
	whereClause = "1=1" // 获取所有数据的查询条件
)

type ArcGISResponse struct {
	Features []json.RawMessage `json:"features"`
	Error    *struct {
		Message string `json:"message"`
	} `json:"error"`
}

func reptile(serviceURL string, filename string, token string) {
	offset := 0
	allFeatures := []json.RawMessage{}

	for {
		params := url.Values{}
		params.Set("f", "json")
		params.Set("where", whereClause)
		params.Set("outFields", "*")
		params.Set("resultOffset", strconv.Itoa(offset))
		params.Set("resultRecordCount", strconv.Itoa(batchSize))
		params.Set("token", token)
		//
		resp, err := http.Get(serviceURL + "?" + params.Encode())

		if err != nil {
			log.Println("请求错误,请检查服务是否正常运行")
			break
		} else {
			log.Printf("%s请求成功，当前索引:%d\n", filename, offset)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		var data ArcGISResponse
		if err := json.Unmarshal(body, &data); err != nil {
			panic(err)
		}

		// 检查错误
		if data.Error != nil {
			panic(data.Error.Message)
		}

		// 保存当前批次数据
		allFeatures = append(allFeatures, data.Features...)

		// 若返回数量不足批次大小，说明已到末尾
		if len(data.Features) < batchSize {
			break
		}
		offset += batchSize // 更新偏移量
	}

	if len(allFeatures) > 0 {
		fmt.Printf("下载：%s，成功获取 %d 条数据\n", filename, len(allFeatures))
		ds_json.Save(filename, allFeatures)
	}
	// 此处处理所有数据 (allFeatures)
}
