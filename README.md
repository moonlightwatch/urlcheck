# 使用

urlcheck [url]


举例：
```
$ urlcheck https://bing.com
```

将会获得如下输出：

```

--- Checking HTTP delay ---
🕒 681.522233ms
🔥 200

--- Resolving DNS ---
🚩 204.79.197.200
🚩 13.107.21.200
🚩 2620:1ec:c11::200

--- Checking HTTP version support ---
🌞 HTTP/1.1
🌞 HTTP/2.0

--- Checking TLS support ---
✅ TLS 1.2

--- Checking certificate ---
🔰 CN=Microsoft Azure TLS Issuing CA 02,O=Microsoft Corporation,C=US
🌐 *.platform.bing.com|*.bing.com|bing.com|ieonline.microsoft.com|*.windowssearch.com|cn.ieonline.microsoft.com|*.origin.bing.com|*.mm.bing.net|*.api.bing.com|*.cn.bing.net|*.cn.bing.com|ssl-api.bing.com|ssl-api.bing.net|*.api.bing.net|*.bingapis.com|bingsandbox.com|feedback.microsoft.com|insertmedia.bing.office.net|r.bat.bing.com|*.r.bat.bing.com|*.dict.bing.com|*.ssl.bing.com|*.appex.bing.com|*.platform.cn.bing.com|wp.m.bing.com|*.m.bing.com|global.bing.com|windowssearch.com|search.msn.com|*.bingsandbox.com|*.api.tiles.ditu.live.com|*.t0.tiles.ditu.live.com|*.t1.tiles.ditu.live.com|*.t2.tiles.ditu.live.com|*.t3.tiles.ditu.live.com|3d.live.com|api.search.live.com|beta.search.live.com|cnweb.search.live.com|ditu.live.com|farecast.live.com|image.live.com|images.live.com|local.live.com.au|localsearch.live.com|ls4d.search.live.com|mail.live.com|mapindia.live.com|local.live.com|maps.live.com|maps.live.com.au|mindia.live.com|news.live.com|origin.cnweb.search.live.com|preview.local.live.com|search.live.com|test.maps.live.com|video.live.com|videos.live.com|virtualearth.live.com|wap.live.com|webmaster.live.com|www.local.live.com.au|www.maps.live.com.au|webmasters.live.com|ecn.dev.virtualearth.net|www.bing.com
🕒 2024-06-27 23:59:59

2024/03/12 08:17:31 Done
```