# 创建坚果云忽略文件、文件夹规则文件

## 坚果云选择性忽略规则说明 & 问题说明

[坚果云如何选择性忽略部分文件/文件夹/文件类型？](http://help.jianguoyun.com/?p=1825)  

可以发现，只能忽略特定的文件夹，必须写特定的全路况，不能使用通配符，只能忽略特定后缀的文件，不能忽略某一规则的文件、文件夹  
比如
- 对于使用 IDEA 来说，要忽略所有的 .idea 文件夹
- 对于前端来说，要忽略所有的 node_modules 文件夹
- 对于 Go 来说，要忽略所有的 vendor 文件夹
- ...

所以为了解决这个问题，决定动手解决写这个项目

## 规则文件说明
目前只针对，Linux、MacOS，对于 Win 未知。。

默认规则文件路径是 `~/.nutstore/db/customExtRules.conf` 


## 使用
```text

$ go get -u -v github.com/dryyun/gather/go-ignore-jianguoyun   // 生成命令 $GOPATH/bin/go-ignore-jianguoyun

$ go-ignore-jianguoyun -h // 查看帮助

$ go-ignore-jianguoyun -depth=3 -dirs=".idea,node_modules,vendor" -root="~/Codes" -readonly=true

```

