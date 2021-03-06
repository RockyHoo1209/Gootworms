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

## ๐ Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Built Using](#built_using)
- [TODO](./TODO.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## ๐ง About <a name = "about"></a>
่ฟๆฏไธไธช็จgolang็ผๅ็ไธปไปๅๅธๅผ็ฌ่ซ

## ๐ Getting Started <a name = "getting_started"></a>
### Project Tree
```
Gootworms
โโ conf
โ  โโ config.yml //็ฌ่ซ่งๅๅๆฐๆฎๅบ็้็ฝฎๆไปถ
โโ Dockerfile   
โโ go.mod
โโ go.sum
โโ LICENSE
โโ main
โโ main.go
โโ README.md
โโ src
โ  โโ Data
โ  โ  โโ Enum  //ๅญๆพๅธธ้ไฟกๆฏ
โ  โ  โ  โโ NodeConstants
โ  โ  โ     โโ NodeConstants.go 
โ  โ  โโ Job.go
โ  โ  โโ Result.go
โ  โโ Master
โ  โ  โโ Master.go
โ  โ  โโ Master_test.go
โ  โโ Slave
โ  โ  โโ Downloader
โ  โ  โ  โโ Downloader.go
โ  โ  โ  โโ Downloader_test.go
โ  โ  โโ Node
โ  โ  โ  โโ NodeManager.go
โ  โ  โโ Spider
โ  โ     โโ Spider.go
โ  โ     โโ Spider_test.go
โ  โโ Utils  //ๅทฅๅท็ฑป
โ     โโ ChromeDriverUtil
โ     โ  โโ ChromeDriverUtil.go
โ     โ  โโ ChromeDriverUtil_test.go
โ     โโ ConfigUtil
โ     โ  โโ ConfigUtil.go
โ     โ  โโ ConfigUtil_test.go
โ     โโ CronUtil
โ     โ  โโ CronUtil.go
โ     โ  โโ CronUtilTest.go
โ     โโ DBUtils
โ     โ  โโ MongoUtil
โ     โ  โ  โโ config.yml
โ     โ  โ  โโ MongoUtil.go
โ     โ  โ  โโ MongoUtil_test.go
โ     โ  โโ RedisUtil
โ     โ     โโ RedisUtil.go
โ     โ     โโ RedisUtil_test.go
โ     โโ HttpUtil  //Http่งฃๆ
โ     โ  โโ DoRequest
โ     โ     โโ DoRequest.go
โ     โ     โโ DoRequest_test.go
โ     โโ LinksUtil  //ๆๅๆๆฌไธญ็้พๆฅ
โ     โ  โโ config.yml
โ     โ  โโ LinksUtil.go
โ     โ  โโ LinksUtil_test.go
โ     โโ TaskUtil
โ        โโ TaskUtil.go
โ        โโ TaskUtil_test.go
โโ TODO.md

```

### Prerequisites

ๅฏๅจๅไธบไป่็น(worker)ๅไธป่็น(master)ไธค็งๆจกๅผใๅผๅฏredis-serverๅmongodbๆๅก๏ผไธๅๅฏๅจไธป่็น(้ๅฎ่ฃๅนถๅฏๅจredisๅmongodb)๏ผๅๅฏๅจไป่็น


ไธป่็นๅฏๅจ
```
go run main master
```

ไป่็นๅฏๅจ
```
go run main worker
```

## โ๏ธ Built Using <a name = "built_using"></a>

- [MongoDB](https://www.mongodb.com/)-4.4.1 -ๅญๅจๆฐๆฎๅบ
- [Redis](https://expressjs.com/)-6.2.1 - ่็นRPC้ไฟกไพ่ต
- [golang](https://vuejs.org/)-1.5 - ๅผๅ่ฏญ่จ


## โ๏ธ Authors <a name = "authors"></a>

- [@RockyHoo1209](https://github.com/RockyHoo1209) - Idea & Initial work
- ๆฌข่ฟ่ฎจ่ฎบๅไบคๆต,qq&email:1284219022@qq.com

## ๐ Acknowledgements <a name = "acknowledgement"></a>

- Hat tip to anyone whose code was used
- Inspiration
- References
  
