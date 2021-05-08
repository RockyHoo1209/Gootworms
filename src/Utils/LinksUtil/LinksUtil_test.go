/*
 * @Description:
 * @Author: Rocky Hoo
 * @Date: 2021-04-26 18:53:05
 * @LastEditTime: 2021-04-26 20:35:49
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
//  package LinksUtil

//  import (
// 	 "encoding/json"
// 	 "fmt"
// 	 "log"

// 	 // "reflect"
// 	 "regexp"
// 	 "strings"
// 	 _ "unsafe"
//  )

//  func main() {
//  s := `<!DOCTYPE html>
//   <html lang="zh-CN" dir="ltr">
// 	<head>
// 	  <title>跟随总书记的脚步 忆浴血湘江悟苦难辉煌_腾讯新闻</title>
// 	  <meta name="keywords" content="跟随总书记的脚步 忆浴血湘江悟苦难辉煌,湘江,习近平总书记,总书记,长征,中央红军,红军烈士">
// 	  <meta name="description" content="跟随总书记的脚步 忆浴血湘江悟苦难辉煌">
// 	  <meta name="apub:time" content="4/26/2021, 10:52:40 AM">
// 	  <meta name="apub:from" content="default">
// 	  <meta http-equiv="X-UA-Compatible" content="IE=Edge" />
//   <link rel="stylesheet" href="//mat1.gtimg.com/pingjs/ext2020/dcom-static/build/static/css/static.css" />
//   <!--[if lte IE 8]><meta http-equiv="refresh" content="0; url=/upgrade.htm"><![endif]-->
//   <!-- <meta name="sogou_site_verification" content="SYWy6ahy7s"/> -->
//   <meta name="baidu-site-verification" content="jJeIJ5X7pP" />
//   <link rel="shortcut icon" href="//mat1.gtimg.com/www/icon/favicon2.ico" />
//   <link rel="stylesheet" href="//vm.gtimg.cn/tencentvideo/txp/style/txp_desktop.css" />
//   <script src="//js.aq.qq.com/js/aq_common.js"></script>
//   <script>
// 	  // 判断如果是动态底层不加载此JS逻辑 2020/1/19 -1
// 	  if(location.href.indexOf('rain') === -1){
// 		  (function(){
// 			  var bp = document.createElement('script');
// 			  var curProtocol = window.location.protocol.split(':')[0];
// 			  if (curProtocol === 'https') {
// 				  bp.src = 'https://zz.bdstatic.com/linksubmit/push.js';
// 			  }
// 			  else {
// 				  bp.src = 'http://push.zhanzhang.baidu.com/push.js';
// 			  }
// 			  var s = document.getElementsByTagName("script")[0];
// 			  s.parentNode.insertBefore(bp, s);
// 		  })();
// 	  }
//   </script>
//   <script src="//mat1.gtimg.com/pingjs/ext2020/configF2017/5df6e3b3.js" charset="utf-8"></script>
//   <script src="//mat1.gtimg.com/pingjs/ext2020/configF2017/5a978a31.js" charset="utf-8"></script>
//   <script>window.conf_dcom = apub_5a978a31 </script><!--[if !IE]>|xGv00|61438c491c69d576aec9846de884f28b<![endif]-->
//   <!--[if !IE]>|xGv00|038d6e161753081e56c192d04873c65c<![endif]-->
// 	  <script>window.DATA = {
// 		  "article_id": "20210426003469",
// 		  "article_type": "0",
// 		  "title": "跟随总书记的脚步 忆浴血湘江悟苦难辉煌",
// 		  "abstract": null,
// 		  "catalog1": "politics",
// 		  "catalog2": "politics_zhongda",
// 		  "media": "央视新闻客户端",
// 		  "media_id": "央视新闻客户端",
// 		  "pubtime": "2021-04-26 10:36:29",
// 		  "comment_id": "6781298085",
// 		  "tags": "湘江,习近平总书记,总书记,长征,中央红军,红军烈士",
// 		  "political": 1,
// 		  "artTemplate": null,
// 		  "FztCompetition": null,
// 		  "FCompetitionName": null,
// 		  "is_deleted": 0,
// 		  "cms_id": "TWF2021042600346900",
// 		  "apubTpl": null,
// 		  "videoArr": []
//   }

// 	  </script>
// 	</head>
// 	<body>
// 	  <div id="TopNav"></div>
// 	  <div id="topAd"></div>
// 	  <div class="qq_conent clearfix">
// 		<div class="LEFT">
// 		  <h1>跟随总书记的脚步 忆浴血湘江悟苦难辉煌</h1>
// 		  <div class="content clearfix">
// 			<div class="LeftTool" id="LeftTool"></div>
// 			<!--内容-->
// 			<div class="content-article">
// 			  <!--导语-->
// 			  <div class="videoPlayerWrap"></div>
// 			  <script>window.DATA.videoArr.push({"title":"时政马上评丨跟随总书记的脚步 忆浴血湘江 悟苦难辉煌","vid":"b32420bvx0n","img":"http://inews.gtimg.com/newsapp_ls/0/13457020995_640480/0","desc":"视频：时政马上评丨跟随总书记的脚步 忆浴血湘江 悟苦难辉煌，时长约4分57秒"})</script>
// 			  <p class="one-p">2018年11月</p>
// 			  <p class="one-p">习近平总书记曾专门作出重要批示</p>
// 			  <p class="one-p">明确要求做好</p>
// 			  <p class="one-p">湘江战役烈士遗骸收殓保护工作</p>
// 			  <p class="one-p">规划建设好纪念设施</p>
// 			  <p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/13455074465/1000">
// 			  </p>
// 			  <p class="one-p">2019年9月</p>
// 			  <p class="one-p">红军长征湘江战役纪念园建成</p>
// 			  <p class="one-p">多年来发掘收殓的</p>
// 			  <p class="one-p">湘江战役红军烈士遗骸</p>
// 			  <p class="one-p">安葬于此</p>
// 			  <p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/13455074818/1000">
// 			  </p>
// 			  <p class="one-p">2021年4月25日</p>
// 			  <p class="one-p">习近平总书记来到</p>
// 			  <p class="one-p">广西桂林全州县才湾镇</p>
// 			  <p class="one-p">红军长征湘江战役纪念园</p>
// 			  <p class="one-p">向湘江战役红军烈士敬献花篮</p>
// 			  <p class="one-p">参观红军长征湘江战役纪念馆</p>
// 			  <p class="one-p">缅怀革命先烈</p>
// 			  <p class="one-p">赓续共产党人精神血脉</p>
// 			  <p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/13455074647/1000">
// 			  </p>
// 			  <p class="one-p">跟随总书记的脚步</p>
// 			  <p class="one-p">体会足迹中的深意</p>
// 			  <p class="one-p">湘江战役发生在</p>
// 			  <p class="one-p">1934年11月27日至12月1日</p>
// 			  <p class="one-p">中央红军与国民党军苦战五昼夜抢渡湘江</p>
// 			  <p class="one-p">突破了国民党军第四道封锁线</p>
// 			  <p class="one-p">粉碎了蒋介石围歼中央红军于湘江以东的企图</p>
// 			  <p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/13455075190/1000">
// 			  </p>
// 			  <p class="one-p">经过此战</p>
// 			  <p class="one-p">中央红军从长征出发时的8.6万余人</p>
// 			  <p class="one-p">锐减至3万余人</p>
// 			  <p class="one-p">三年不饮湘江水</p>
// 			  <p class="one-p">十年不食湘江鱼</p>
// 			  <p class="one-p">湘江之战苦难沉痛</p>
// 			  <p class="one-p">红军浴火重生 走向辉煌</p>
// 			  <p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/13455075334/1000">
// 			  </p>
// 			  <p class="one-p">长征是一次理想信念的伟大远征</p>
// 			  <p class="one-p">长征是一次检验真理的伟大远征</p>
// 			  <p class="one-p">长征是一次唤醒民众的伟大远征</p>
// 			  <p class="one-p">长征是一次开创新局的伟大远征</p>
// 			  <p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/13455074969/1000">
// 			  </p>
// 			  <p class="one-p">多年来 习近平总书记</p>
// 			  <p class="one-p">每次到革命老区考察调研</p>
// 			  <p class="one-p">都会瞻仰革命历史纪念场所</p>
// 			  <p class="one-p">了解历史才能看得远</p>
// 			  <p class="one-p">永葆初心才能走得远</p>
// 			  <p class="one-p">走好新时代的长征路</p>
// 			  <p class="one-p">就是对英烈最好的纪念</p>
// 			  <div id="Status"></div>
// 			</div>
// 		  </div>
// 		  <div id="Comment"></div>
// 		  <div id="Recomend"></div>
// 		</div>
// 		<div class="RIGHT" id="RIGHT"></div>
// 	  </div>
// 	  <div id="bottomAd"></div>
// 	  <div class="qq_footer" id="Foot"></div>
// 	  <div id="GoTop"></div>
// 	  <script src="//mat1.gtimg.com/libs/jquery/1.12.0/jquery.min.js"></script>
//   <script type="text/javascript" src="//h5.ssp.qq.com/static/web/websites/pcnewsplugin/sspad_20200821.js" charset="utf-8"></script>
//   <script src="//mat1.gtimg.com/pingjs/ext2020/dc2017/publicjs/m/ping.js" charset="gbk"></script>
//   <script src="//mat1.gtimg.com/pingjs/ext2020/2018/js/check-https-content.js"></script>
//   <script>
//   if(typeof(pgvMain) == 'function'){pgvMain();}
//   </script>
//   <script type="text/javascript" src="//imgcache.qq.com/qzone/biz/comm/js/qbs.js"></script>
//   <script type="text/javascript">
//   var TIME_BEFORE_LOAD_CRYSTAL = (new Date).getTime();
//   </script>
//   <script src="//ra.gtimg.com/web/crystal/v4.7Beta05Build050/crystal-min.js" id="l_qq_com" arguments="{'extension_js_src':'//ra.gtimg.com/web/crystal/v4.7Beta05Build050/crystal_ext-min.js', 'jsProfileOpen':'false', 'mo_page_ratio':'0.01', 'mo_ping_ratio':'0.01', 'mo_ping_script':'//ra.gtimg.com/sc/mo_ping-min.js'}"></script>
//   <script type="text/javascript">
//   if(typeof crystal === 'undefined' && Math.random() <= 1) {
// 	(function() {
// 	  var TIME_AFTER_LOAD_CRYSTAL = (new Date).getTime();
// 	  var img = new Image(1,1);
// 	  img.src = "//dp3.qq.com/qqcom/?adb=1&dm=new&err=1002&blockjs="+(TIME_AFTER_LOAD_CRYSTAL-TIME_BEFORE_LOAD_CRYSTAL);
// 	})();
//   }
//   </script>
//   <style>.absolute{position:absolute;}</style>
//   <!--[if !IE]>|xGv00|bfa6be71716f6329ed6738978b6c1e2d<![endif]-->

//   <script>
//   var _mtac = {};
//   (function() {
// 	  var mta = document.createElement("script");
// 	  mta.src = "//pingjs.qq.com/h5/stats.js?v2.0.2";
// 	  mta.setAttribute("name", "MTAH5");
// 	  mta.setAttribute("sid", "500651042");
// 	  var s = document.getElementsByTagName("script")[0];
// 	  s.parentNode.insertBefore(mta, s);
//   })();
//   </script><!--[if !IE]>|xGv00|4ade69361c277979a5d5757b4776e923<![endif]-->
//   <script src="//mat1.gtimg.com/pingjs/ext2020/dcom-static/build/static/js/static.js"></script>
//   <!--[if !IE]>|xGv00|78fe8a44ba68d8b81e1f6f713a13b0c5<![endif]-->
// 	</body>
//   </html>`
// 	 regex := `window.DATA.*?({.*?})`
// 	 //  regex := `({.*?})`
// 	 s = strings.Replace(s, "\n", "", -1)
// 	 s = strings.Replace(s, " ", "", -1)
// 	 s = strings.Replace(s, "\t", "", -1)
// 	 //  fmt.Printf(s)
// 	 reg := regexp.MustCompile(regex)
// 	 items := reg.FindStringSubmatch(s)
// 	 var mapResult map[string]interface{}
// 	 if err := json.Unmarshal([]byte(items[1]), &mapResult); err != nil {
// 		 log.Fatal(err)
// 	 }
// 	 mapResult["man"] = "rockyhoo"
// 	 fmt.Println(mapResult)
// 	 b,err:=json.Marshal(mapResult)
// 	 if err != nil {
// 		 fmt.Println("json.Marshal failed:", err)
// 		 return
// 	 }
// 	 fmt.Println(string(b))
//  }
package LinksUtil

import (
	"fmt"
	"log"
	"main/src/Utils/ConfigUtil"
	"testing"
)



func TestLinkUtil(*testing.T){
	s:=`<!DOCTYPE html>
	<html lang="zh-CN" dir="ltr">
	  <head>
		<title>跟随总书记的脚步 忆浴血湘江悟苦难辉煌_腾讯新闻</title>
		<meta name="keywords" content="跟随总书记的脚步 忆浴血湘江悟苦难辉煌,湘江,习近平总书记,总书记,长征,中央红军,红军烈士">
		<meta name="description" content="跟随总书记的脚步 忆浴血湘江悟苦难辉煌">
		<meta name="apub:time" content="4/26/2021, 10:52:40 AM">
		<meta name="apub:from" content="default">
		<meta http-equiv="X-UA-Compatible" content="IE=Edge" />
	<link rel="stylesheet" href="//mat1.gtimg.com/pingjs/ext2020/dcom-static/build/static/css/static.css" />
	<!--[if lte IE 8]><meta http-equiv="refresh" content="0; url=/upgrade.htm"><![endif]-->
	<!-- <meta name="sogou_site_verification" content="SYWy6ahy7s"/> -->
	<meta name="baidu-site-verification" content="jJeIJ5X7pP" />
	<link rel="shortcut icon" href="//mat1.gtimg.com/www/icon/favicon2.ico" />
	<link rel="stylesheet" href="//vm.gtimg.cn/tencentvideo/txp/style/txp_desktop.css" />
	<script src="//js.aq.qq.com/js/aq_common.js"></script>
	<script>
		// 判断如果是动态底层不加载此JS逻辑 2020/1/19 -1
		if(location.href.indexOf('rain') === -1){
			(function(){
				var bp = document.createElement('script');
				var curProtocol = window.location.protocol.split(':')[0];
				if (curProtocol === 'https') {
					bp.src = 'https://zz.bdstatic.com/linksubmit/push.js';
				}
				else {
					bp.src = 'http://push.zhanzhang.baidu.com/push.js';
				}
				var s = document.getElementsByTagName("script")[0];
				s.parentNode.insertBefore(bp, s);
			})();
		}
	</script>
	<script src="//mat1.gtimg.com/pingjs/ext2020/configF2017/5df6e3b3.js" charset="utf-8"></script>
	<script src="//mat1.gtimg.com/pingjs/ext2020/configF2017/5a978a31.js" charset="utf-8"></script>
	<script>window.conf_dcom = apub_5a978a31 </script><!--[if !IE]>|xGv00|61438c491c69d576aec9846de884f28b<![endif]-->
	<!--[if !IE]>|xGv00|038d6e161753081e56c192d04873c65c<![endif]-->
		<script>window.DATA = {
			"article_id": "20210426003469",
			"article_type": "0",
			"title": "跟随总书记的脚步 忆浴血湘江悟苦难辉煌",
			"abstract": null,
			"catalog1": "politics",
			"catalog2": "politics_zhongda",
			"media": "央视新闻客户端",
			"media_id": "央视新闻客户端",
			"pubtime": "2021-04-26 10:36:29",
			"comment_id": "6781298085",
			"tags": "湘江,习近平总书记,总书记,长征,中央红军,红军烈士",
			"political": 1,
			"artTemplate": null,
			"FztCompetition": null,
			"FCompetitionName": null,
			"is_deleted": 0,
			"cms_id": "TWF2021042600346900",
			"apubTpl": null,
			"videoArr": []
	}

		</script>
	  </head>
	  <body>
		<div id="TopNav"></div>
		<div id="topAd"></div>
		<div class="qq_conent clearfix">
		  <div class="LEFT">
			<h1>跟随总书记的脚步 忆浴血湘江悟苦难辉煌</h1>
			<div class="content clearfix">
			  <div class="LeftTool" id="LeftTool"></div>
			  <!--内容-->
			  <div class="content-article">
				<!--导语-->
				<div class="videoPlayerWrap"></div>
				<script>window.DATA.videoArr.push({"title":"时政马上评丨跟随总书记的脚步 忆浴血湘江 悟苦难辉煌","vid":"b32420bvx0n","img":"http://inews.gtimg.com/newsapp_ls/0/13457020995_640480/0","desc":"视频：时政马上评丨跟随总书记的脚步 忆浴血湘江 悟苦难辉煌，时长约4分57秒"})</script>
				<p class="one-p">2018年11月</p>
				<p class="one-p">习近平总书记曾专门作出重要批示</p>
				<p class="one-p">明确要求做好</p>
				<p class="one-p">湘江战役烈士遗骸收殓保护工作</p>
				<p class="one-p">规划建设好纪念设施</p>
				<p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/13455074465/1000">
				</p>
				<p class="one-p">2019年9月</p>
				<p class="one-p">红军长征湘江战役纪念园建成</p>
				<p class="one-p">多年来发掘收殓的</p>
				<p class="one-p">湘江战役红军烈士遗骸</p>
				<p class="one-p">安葬于此</p>
				<p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/13455074818/1000">
				</p>
				<p class="one-p">2021年4月25日</p>
				<p class="one-p">习近平总书记来到</p>
				<p class="one-p">广西桂林全州县才湾镇</p>
				<p class="one-p">红军长征湘江战役纪念园</p>
				<p class="one-p">向湘江战役红军烈士敬献花篮</p>
				<p class="one-p">参观红军长征湘江战役纪念馆</p>
				<p class="one-p">缅怀革命先烈</p>
				<p class="one-p">赓续共产党人精神血脉</p>
				<p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/13455074647/1000">
				</p>
				<p class="one-p">跟随总书记的脚步</p>
				<p class="one-p">体会足迹中的深意</p>
				<p class="one-p">湘江战役发生在</p>
				<p class="one-p">1934年11月27日至12月1日</p>
				<p class="one-p">中央红军与国民党军苦战五昼夜抢渡湘江</p>
				<p class="one-p">突破了国民党军第四道封锁线</p>
				<p class="one-p">粉碎了蒋介石围歼中央红军于湘江以东的企图</p>
				<p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/13455075190/1000">
				</p>
				<p class="one-p">经过此战</p>
				<p class="one-p">中央红军从长征出发时的8.6万余人</p>
				<p class="one-p">锐减至3万余人</p>
				<p class="one-p">三年不饮湘江水</p>
				<p class="one-p">十年不食湘江鱼</p>
				<p class="one-p">湘江之战苦难沉痛</p>
				<p class="one-p">红军浴火重生 走向辉煌</p>
				<p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/13455075334/1000">
				</p>
				<p class="one-p">长征是一次理想信念的伟大远征</p>
				<p class="one-p">长征是一次检验真理的伟大远征</p>
				<p class="one-p">长征是一次唤醒民众的伟大远征</p>
				<p class="one-p">长征是一次开创新局的伟大远征</p>
				<p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/13455074969/1000">
				</p>
				<p class="one-p">多年来 习近平总书记</p>
				<p class="one-p">每次到革命老区考察调研</p>
				<p class="one-p">都会瞻仰革命历史纪念场所</p>
				<p class="one-p">了解历史才能看得远</p>
				<p class="one-p">永葆初心才能走得远</p>
				<p class="one-p">走好新时代的长征路</p>
				<p class="one-p">就是对英烈最好的纪念</p>
				<div id="Status"></div>
			  </div>
			</div>
			<div id="Comment"></div>
			<div id="Recomend"></div>
		  </div>
		  <div class="RIGHT" id="RIGHT"></div>
		</div>
		<div id="bottomAd"></div>
		<div class="qq_footer" id="Foot"></div>
		<div id="GoTop"></div>
		<script src="//mat1.gtimg.com/libs/jquery/1.12.0/jquery.min.js"></script>
	<script type="text/javascript" src="//h5.ssp.qq.com/static/web/websites/pcnewsplugin/sspad_20200821.js" charset="utf-8"></script>
	<script src="//mat1.gtimg.com/pingjs/ext2020/dc2017/publicjs/m/ping.js" charset="gbk"></script>
	<script src="//mat1.gtimg.com/pingjs/ext2020/2018/js/check-https-content.js"></script>
	<script>
	if(typeof(pgvMain) == 'function'){pgvMain();}
	</script>
	<script type="text/javascript" src="//imgcache.qq.com/qzone/biz/comm/js/qbs.js"></script>
	<script type="text/javascript">
	var TIME_BEFORE_LOAD_CRYSTAL = (new Date).getTime();
	</script>
	<script src="//ra.gtimg.com/web/crystal/v4.7Beta05Build050/crystal-min.js" id="l_qq_com" arguments="{'extension_js_src':'//ra.gtimg.com/web/crystal/v4.7Beta05Build050/crystal_ext-min.js', 'jsProfileOpen':'false', 'mo_page_ratio':'0.01', 'mo_ping_ratio':'0.01', 'mo_ping_script':'//ra.gtimg.com/sc/mo_ping-min.js'}"></script>
	<script type="text/javascript">
	if(typeof crystal === 'undefined' && Math.random() <= 1) {
	  (function() {
		var TIME_AFTER_LOAD_CRYSTAL = (new Date).getTime();
		var img = new Image(1,1);
		img.src = "//dp3.qq.com/qqcom/?adb=1&dm=new&err=1002&blockjs="+(TIME_AFTER_LOAD_CRYSTAL-TIME_BEFORE_LOAD_CRYSTAL);
	  })();
	}
	</script>
	<style>.absolute{position:absolute;}</style>
	<!--[if !IE]>|xGv00|bfa6be71716f6329ed6738978b6c1e2d<![endif]-->

	<script>
	var _mtac = {};
	(function() {
		var mta = document.createElement("script");
		mta.src = "//pingjs.qq.com/h5/stats.js?v2.0.2";
		mta.setAttribute("name", "MTAH5");
		mta.setAttribute("sid", "500651042");
		var s = document.getElementsByTagName("script")[0];
		s.parentNode.insertBefore(mta, s);
	})();
	href="https://new.qq.com/ch/ent/"
	</script><!--[if !IE]>|xGv00|4ade69361c277979a5d5757b4776e923<![endif]-->
	<script src="//mat1.gtimg.com/pingjs/ext2020/dcom-static/build/static/js/static.js"></script>
	<!--[if !IE]>|xGv00|78fe8a44ba68d8b81e1f6f713a13b0c5<![endif]-->
	  </body>
	</html>`
	ConfigUtil.InitConfig("config")
	items,err:=ExtractItems2Map("",s,ConfigUtil.GetStringMap("client.rule.item_rule"))
	if err!=nil{
		log.Println(err)
	}
	fmt.Println(items)
	reg_url:="((http[s]{0,1})://)new((\\.[a-zA-Z]{2,6})|([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}))(:[0-9]{1,4})*(/[a-zA-Z0-9\\&%_\\./-~-]*)?"
	urls,_:=ExtractUrls(s,reg_url)
	fmt.Println(urls)
}
