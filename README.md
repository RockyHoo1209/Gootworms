<p align="center">
  <a href="" rel="noopener"></a>
</p>

<h1 align="center">Gootworms</h1>

<div align="center"> 

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> Few lines describing your project.
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Built Using](#built_using)
- [TODO](./TODO.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>
这是一个用golang编写的主从分布式爬虫

## 🏁 Getting Started <a name = "getting_started"></a>
### Project Tree
```
Gootworms
├─ conf
│  └─ config.yml //爬虫规则及数据库的配置文件
├─ Dockerfile   
├─ go.mod
├─ go.sum
├─ LICENSE
├─ main
├─ main.go
├─ README.md
├─ src
│  ├─ Data
│  │  ├─ Enum  //存放常量信息
│  │  │  └─ NodeConstants
│  │  │     └─ NodeConstants.go 
│  │  ├─ Job.go
│  │  └─ Result.go
│  ├─ Master
│  │  ├─ Master.go
│  │  └─ Master_test.go
│  ├─ Slave
│  │  ├─ Downloader
│  │  │  ├─ Downloader.go
│  │  │  └─ Downloader_test.go
│  │  ├─ Node
│  │  │  └─ NodeManager.go
│  │  └─ Spider
│  │     ├─ Spider.go
│  │     └─ Spider_test.go
│  └─ Utils  //工具类
│     ├─ ChromeDriverUtil
│     │  ├─ ChromeDriverUtil.go
│     │  └─ ChromeDriverUtil_test.go
│     ├─ ConfigUtil
│     │  ├─ ConfigUtil.go
│     │  └─ ConfigUtil_test.go
│     ├─ CronUtil
│     │  ├─ CronUtil.go
│     │  └─ CronUtilTest.go
│     ├─ DBUtils
│     │  ├─ MongoUtil
│     │  │  ├─ config.yml
│     │  │  ├─ MongoUtil.go
│     │  │  └─ MongoUtil_test.go
│     │  └─ RedisUtil
│     │     ├─ RedisUtil.go
│     │     └─ RedisUtil_test.go
│     ├─ HttpUtil  //Http解析
│     │  └─ DoRequest
│     │     ├─ DoRequest.go
│     │     └─ DoRequest_test.go
│     ├─ LinksUtil  //提取文本中的链接
│     │  ├─ config.yml
│     │  ├─ LinksUtil.go
│     │  └─ LinksUtil_test.go
│     └─ TaskUtil
│        ├─ TaskUtil.go
│        └─ TaskUtil_test.go
└─ TODO.md

```

### Prerequisites

启动分为从节点(worker)和主节点(master)两种模式。开启redis-server和mongodb服务，且先启动主节点(需安装并启动redis和mongodb)，再启动从节点


主节点启动
```
go run main master
```

从节点启动
```
go run main worker
```

## ⛏️ Built Using <a name = "built_using"></a>

- [MongoDB](https://www.mongodb.com/)-4.4.1 -存储数据库
- [Redis](https://expressjs.com/)-6.2.1 - 节点RPC通信依赖
- [golang](https://vuejs.org/)-1.5 - 开发语言


## ✍️ Authors <a name = "authors"></a>

- [@RockyHoo1209](https://github.com/RockyHoo1209) - Idea & Initial work
- 欢迎讨论和交流,qq&email:1284219022@qq.com

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Hat tip to anyone whose code was used
- Inspiration
- References
  
