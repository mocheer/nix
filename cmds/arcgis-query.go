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
// nix arcgis-query  <服务地址> <保存路径> <token>
var ArcgisQuery = &cli.Command{
	Name:  "arcgis-query",
	Usage: "获取arcgis的query服务数据",
	Action: func(c *cli.Context) error {

		url := c.Args().Get(0)

		if url == "" {
			reptileDefault()
		} else {
			fname := c.Args().Get(1)
			if fname == "" {
				fname = clock.Now().Fmt(clock.FmtCompactFullDate) + ".json"
			}
			token := c.Args().Get(2)
			go reptile(url, fname, token, false)
		}
		select {}
	},
}

func reptileDefault() {
	serviceURL := "http://10.135.210.71:6116/ags-proxy/server1/rest/services/fj_floodrisk/fj_floodrisk_map/MapServer/"
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhY2Nlc3NJZCI6IjE2MTA5MjQzMDczOTMwODEzNDYiLCJ1c2VyTmFtZSI6InNjMjAyNCIsImV4cCI6MTc1MjE5OTA5MX0.Bd2RuqJOQjbQBN2WHE73VYNp-AAGom37ANRGq9PbRL8"
	go reptile(serviceURL+"0/query", "断面.json", token, false)
	go reptile(serviceURL+"1/query", "5y.json", token, false)
	go reptile(serviceURL+"2/query", "10y.json", token, false)
	go reptile(serviceURL+"3/query", "20y.json", token, false)
	go reptile(serviceURL+"4/query", "50y.json", token, false)
	go reptile(serviceURL+"5/query", "100y.json", token, false)
	go reptile(serviceURL+"6/query", "200y.json", token, false)
	select {}
}

const (
	batchSize   = 1000  // 每批请求数量（不能超过服务端设置的最大值）, 很多服务端设置为2000，有些设置为1000
	whereClause = "1=1" // 获取所有数据的查询条件
)

type ArcGISResponse struct {
	Features []json.RawMessage `json:"features"`
	Error    *struct {
		Message string `json:"message"`
	} `json:"error"`
}

func reptile(serviceURL string, filename string, token string, flag bool) {
	offset := 0
	allFeatures := []json.RawMessage{}
	whereParam := whereClause
	idFields := []string{"OBJECTID", "FID"}
	i := 0

	for {
		if flag {
			if i >= len(idFields) {
				return
			}
			field := idFields[i]
			whereParam = fmt.Sprintf("%s>=%d AND %s<=%d", field, offset+1, field, offset+batchSize)
		}
		params := url.Values{}
		params.Set("f", "json")
		params.Set("where", whereParam)
		params.Set("outFields", "*")
		params.Set("resultOffset", strconv.Itoa(offset))
		params.Set("resultRecordCount", strconv.Itoa(batchSize))
		if token != "" {
			params.Set("token", token)
		}
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
			log.Println("请求失败", serviceURL+"?"+params.Encode())
			return
		}

		// 检查错误
		if data.Error != nil {
			log.Println("请求成功，但信息提示错误", data.Error.Message, len(data.Features), serviceURL+"?"+params.Encode())
			if flag {
				i++
				continue
			}
			return
		}
		if len(data.Features) == 0 {
			log.Println("请求成功，但数据为空", serviceURL+"?"+params.Encode())
			return
		}

		// 保存当前批次数据
		allFeatures = append(allFeatures, data.Features...)
		log.Println("当前总数", len(allFeatures), string(allFeatures[len(allFeatures)-1]))

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
