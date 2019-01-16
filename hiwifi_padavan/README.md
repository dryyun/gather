# 极路由刷 Padavan 固件

## Google-fu 

### chnroute mode 

```shell
// 生成最新的 chnroute.txt  
curl 'http://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest' | grep ipv4 | grep CN | awk -F\| '{ printf("%s/%d\n", $4, 32-log($5)/log(2)) }' > chnroute.txt  

// 生成 md5sum 文件  
md5sum chnroute.txt > chnroute.txt.md5sum.txt  

// ps，如果是 mac 需要先 `brew install md5sha1sum`  
```

使用 `https://raw.githubusercontent.com/dryyun/gather/master/hiwifi_padavan/chnroute.txt`

### gfwlist mode

```shell

// 生成最新的  glwlist.txt

curl  https://raw.githubusercontent.com/gfwlist/gfwlist/master/gfwlist.txt > ss.gfwlist.lite.txt

// 生成 md5sum 文件  

md5sum ss.gfwlist.lite.txt > ss.gfwlist.lite.txt.md5sum.txt

//
curl https://raw.githubusercontent.com/gfwlist/gfwlist/master/gfwlist.txt  | base64 --decode  > ss.gfwlist.lite.txt.plain

// 
md5sum ss.gfwlist.lite.txt.plain > ss.gfwlist.lite.txt.plain.md5sum.txt

```
