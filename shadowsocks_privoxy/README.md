# shadowsocks + privoxy 实现代理

## Imges
[shadowsocks](https://hub.docker.com/r/mritd/shadowsocks/)  
[privoxy](https://hub.docker.com/r/splazit/privoxy-alpine/)  

## 说明

### shadowsocks 配置


cp ./shadowsocks/sslocal.example.json ./shadowsocks/sslocal.json   

配置文件 ./shadowsocks/sslocal.json  

`local` 部分保持不动，主要是 `local_port` 配置了 1080 ， 牵一发而动全身  
`server` 部分配置 ss-server 的内容  

### privoxy 配置

配置文件 ./privoxy/config  

`gfwlist.action` 生成使用了 [gfwlist2privoxy](https://github.com/zfl9/gfwlist2privoxy)， 稍微有所改动  

>  
> linux 使用  bash gfwlist2privoxy  'shadowsocks:1080'  && mv -f gfwlist.action ./privoxy/  
> mac 使用 bash gfwlist2privoxy-mac-zsh  'shadowsocks:1080'  && mv -f gfwlist.action ./privoxy/  
> 懒得生成的话，使用自带的，勉强也能用  

### docker compose 设置

查看文件 docker-compose.yml   
暴露端口   
1080 是 shadowsocks 设置的 sock5 代理端口  
8118 是 privoxy 转发 1080 的 http 代理端口  

## 运行 

docker-compose up -d privoxy 

## 使用

在需要配置代理的地方，设置成 127.0.0.1:8118 就行  
比如我的终端设置  

vim .bashrc  增加了 function  
```
function proxy() {
    export http_proxy=http://127.0.0.1:8118 && export https_proxy=http://127.0.0.1:8118
    echo "已开启代理";
}

function unproxy() {
    unset http_proxy && unset http_proxy
    echo "已关闭代理"
}
```


## 参考

[ss-local 全局代理](https://www.zfl9.com/ss-local.html)









