---
name: nix
description: 个人开发工具箱 - 集成了GIS数据处理、文件管理、媒体处理、开发构建等实用工具
---

# nix 开发工具箱

**nix** 是一个 Go 编写的命令行工具集，提供开发、GIS、文件处理等方面的常用功能。

## GIS 与地图相关

### nix tile
下载在线瓦片地图（支持 ArcGIS 瓦片服务、SuperMap 等格式）
```
nix tile <服务URL> [保存路径]
```
- 参数1: 瓦片服务URL（支持 {z}/{x}/{y} 占位符）
- 参数2: 保存路径（可选，默认 data/tile-日期）
- 缩放范围: 0-18

### nix tile-xyz
下载在线瓦片地图（支持多子域名）
```
nix tile-xyz <服务URL> [保存路径]
```
- 参数1: 瓦片服务URL（支持 {z}/{x}/{y}/{s} 占位符）
- 参数2: 保存路径（可选，默认 public/tile-日期）
- 子域名: 自动轮询 1,2,3
- 适用场景: 高德地图、天地图等 XYZ 规范的瓦片服务

### nix tile-arcgis
下载 ArcGIS 本地瓦片地图
```
nix tile-arcgis <服务URL> [保存路径]
```
- 参数1: ArcGIS 瓦片服务URL
- 参数2: 保存路径（可选，默认 data/tile-日期）

### nix tileset
下载在线 3DTiles 模型数据
```
nix tileset <tileset.json URL> [保存路径]
```
- 参数1: tileset.json 文件的 URL
- 参数2: 保存路径（可选，默认 data/3dtiles-日期）
- 适用场景: Cesium 3DTiles、Google 3D Tiles、Mars3D 等

### nix tileset-dem
下载在线 3DTiles DEM 高程数据
```
nix tileset-dem <layer.json URL> [保存路径]
```
- 参数1: 高程数据服务 URL
- 参数2: 保存路径（可选，默认 data/3dtiles_dem-日期）

### nix tif
查看 GeoTIFF 文件的标签属性信息
```
nix tif <文件名>
```

### nix arcgis-query
获取 ArcGIS 的 query 服务数据，支持分页批量拉取
```
nix arcgis-query <服务地址> [保存路径] [token]
```
- 参数1: ArcGIS query 服务 URL
- 参数2: 保存文件名（可选，默认 日期.json）
- 参数3: token（可选，用于需要认证的服务）
- 特点: 自动分批请求（每批1000条），支持断点续传，保存为 JSON 格式

### nix arcgis-mapserver-query
一次性采集 ArcGIS MapServer 下所有图层的数据
```
nix arcgis-mapserver-query <服务地址> <保存路径> [token]
```
- 参数1: MapServer 服务 URL（如 http://host/arcgis/rest/services/xxx/MapServer）
- 参数2: 保存目录路径
- 参数3: token（可选）
- 特点: 自动遍历所有图层，跳过已存在的文件，并发拉取数据

### nix arcgis-mapserver-query-test
测试 ArcGIS MapServer 单个要素的查询（开发调试用）
```
nix arcgis-mapserver-query-test
```

### nix point-cloud
从点云数据（LAS）生成 3DTiles 模型
```
nix point-cloud
```
（基于 gocesiumtiler 库，当前为骨架实现）

## 3D模型处理

### nix FBX2glTF
将 FBX 模型格式转换为 glTF/glB 格式
```
nix FBX2glTF <参数...>
```
- 内置自动添加 `-b -d` 参数（二进制输出、启用 Draco 压缩）
- 底层工具: Facebook FBX2glTF

### nix basis
将 PNG 等图片转为 Basis Universal 纹理格式（KTX2）
```
nix basis <图片文件...>
```
- 自动添加 `-ktx2` 参数使用 KTX2 容器格式
- 底层工具: basisu.exe
- 参考: `basis.ps1` 提供更多参数示例（-linear、-comp_level 等）

### nix gltf-texture
提取 glTF 模型的纹理图片
```
nix gltf-texture <glTF文件路径>
```

## 视频与多媒体

### nix ffmpeg
执行 FFmpeg 命令进行视频/音频处理
```
nix ffmpeg <ffmpeg参数...>
```
- 内置 ffmpeg.exe，可执行任意 ffmpeg 命令
- 适用场景: 格式转换、视频截图、HLS 切片、压缩等

## 文件操作

### nix fs append
批量往文件追加内容
```
nix fs append -d <目录> -s <内容> -suffix <后缀>
```
- `-d, --dir`: 目标目录
- `-s, --start`: 要追加的内容
- `--suffix`: 文件后缀过滤

### nix fs rename
批量重命名文件（当前为骨架实现）
```
nix fs rename -d <目录>
```

### nix fs split
切割大文件为多个小文件
```
nix fs split -n <文件名> -cs <块大小(字节)>
```
- `-n, --name`: 要切割的文件
- `-cs, --chunkSize`: 块大小，默认 16GB

### nix fs merge
合并被切割的大文件
```
nix fs merge -d <目录>
```
- `-d, --dir`: 被切割文件所在目录

## 系统与服务

### nix nssm
使用 NSSM 注册和管理 Windows 服务
```
nix nssm install <服务名> <执行路径>
nix nssm start <服务名>
nix nssm stop <服务名>
nix nssm remove <服务名>
```
- 内置 nssm.exe（Non-Sucking Service Manager）

### nix upx
使用 UPX 压缩可执行文件
```
nix upx -9 <文件名>
```
- 内置 upx.exe

### nix active
模拟鼠标和键盘操作，防止电脑进入休眠/锁屏
```
nix active
```

### nix serve
启动一个简单的 Web 文件服务器（默认端口 9212）
```
nix serve
```
- 支持 URL 查询参数映射到文件名（用于 API 资源的静态化）
- 默认监听: `http://localhost:9212`

### nix scoop
安装 Scoop 包管理器（Windows）
```
nix scoop
```

## 开发与构建

### nix run
执行 package.json 中定义的脚本
```
nix run <脚本名称>
```
- 支持 `{appName}`、`{name}`、`{version}`、`{execPath}` 模板变量替换
- 支持 `脚本名#post` 钩子脚本（脚本执行后自动运行）

### nix build
`nix run build` 的简写

### nix dev
`nix run dev` 的简写

### nix version
自动递增 package.json 中的版本号（最后一位 +1）
```
nix version
```

### nix publish
发布 Git tag 版本
```
nix publish [tag名称]
```
- 无参数: 从 package.json 的 version 创建 tag 并推送
- 有参数: 使用指定 tag 名称

### nix esbuild
使用 esbuild 打包 JavaScript/TypeScript 文件
```
nix esbuild <入口文件>
```
- 输出: out.js（压缩、混淆）

### nix install
升级安装所有 Go 依赖包
```
nix install
```

### nix init
在当前目录初始化 package.json（如果不存在）
```
nix init
```

### nix exec
执行 scripts 目录下的脚本
```
nix exec <脚本名>
```

## 信息查询

### nix about
显示程序相关信息
```
nix about
```
- 显示: 可执行文件路径、当前工作目录

### nix package
查看当前项目的 package.json
```
nix package
```

### nix ip
查看当前 IP 信息（当前为骨架实现）
```
nix ip
```

## 数据库

### nix struct
从 PostgreSQL 数据库中读取表结构，转换为 Go struct
```
nix struct <schema名>
```
- 从 pg_tables/pg_class/pg_attribute/pg_type 系统表读取
- 读取 `./assets/config/app.json` 中的 DSN 配置
- 支持 JSON 类型字段自动解析

## 加密与安全

### nix rsa
生成 RSA 2048 位密钥文件
```
nix rsa [输出目录]
```
- 输出目录可选，默认 `.nix/rsa/`

## 其他工具

### nix rename
重命名（当前为骨架实现）

### nix clone
Git 克隆（当前为骨架实现，gittools 包已有 SSH 认证的 Pull 功能）

## 项目结构

```
d:\code\go\nix\
├── main.go              # 入口，注册所有命令
├── cmds/                # 命令实现
│   ├── *.go             # 各命令实现
│   ├── gittools/        # Git 工具（SSH 认证拉取）
│   ├── robot/           # 电脑机器人（鼠标键盘自动化框架）
│   ├── types/           # 类型定义（PackageJSON）
│   └── tools/           # 嵌入的二进制工具
├── global/              # 全局常量
├── tests/               # 测试数据
├── scripts/             # 脚本
└── package.json         # 构建脚本
```

## 注意事项

1. 大部分命令工作在**当前工作目录**下
2. tif/tileset/arcgis-query 等 GIS 命令涉及网络请求，注意网络连通性
3. upx/nssm/fbkx2gltf/ffmpeg/basis 命令通过 `sys.MemExec` 嵌入执行，无需额外安装
4. arcgis 系列命令默认批次大小为 1000 条，可通过代码中的 `batchSize` 常量调整