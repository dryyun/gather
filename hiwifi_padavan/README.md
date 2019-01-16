# 极路由刷固件 Padavan 固件

## 刷机变砖恢复教程

[极路由Rom下载 ](https://app.hiwifi.com/dstore.php?m=download&a=info)

[变砖恢复教程](https://s.histatic.com/ued/rom/romdoc.html)

```txt
简述：
断电 - 按 Reset - 加电 - 4-6 秒后松开 Reset - 本机自动获取，网关为 192.168.2.1 - 上传极路由 ROM 恢复
```

## 刷固件步骤

### source 文件夹内容说明

- HC5962-sysupgrade-20170810-3a807c77.bin  极路由 B70 恢复固件
- pb-boot-hc5962.bin  输入的 pb-boot
- B70_3.4.3.9-099_20170910-2224.trx  Padavan 固件



### 开通 SSH 权限

我是直接【开发者模式】开启了 Root 权限，可以使用1022端口登录到极路由，这样就不能保修了，不过也不是很在意。

网上有绕过这个限制的方法，可以找找。

### 备份

#### 备份分区

```shell
// 极路由操作，登录之后
$ cd /tmp
$ cat /proc/mtd  // => 查看固件信息

dd if=/dev/mtd0 of=/tmp/u-boot.bin   // => 备份打包 mtd0 为 u-boot.bin 文件到 tmp 目录下
dd if=/dev/mtd1 of=/tmp/debug.bin
dd if=/dev/mtd2 of=/tmp/Factory.bin
dd if=/dev/mtd3 of=/tmp/firmware.bin
dd if=/dev/mtd4 of=/tmp/kernel.bin
dd if=/dev/mtd5 of=/tmp/rootfs.bin
dd if=/dev/mtd6 of=/tmp/hw_panic.bin
dd if=/dev/mtd7 of=/tmp/bdinfo.bin
dd if=/dev/mtd8 of=/tmp/backup.bin
dd if=/dev/mtd9 of=/tmp/overlay.bin
dd if=/dev/mtd10 of=/tmp/firmware_backup.bin
dd if=/dev/mtd11 of=/tmp/oem.bin
dd if=/dev/mtd12 of=/tmp/opt.bin

```

```shell
// 本机操作
// 应该是把上面的 *.bin 文件备份到本机
// 但是，我懒，我直接备份了整个 /tmp 目录到 Downloads 目录

$ scp -r -P 1022 root@192.168.199.1:/tmp ~/Downloads  

```

#### 备份 MAC 地址

通过 `ifconfig` 命令就能全部获取，不过保险起见可以这样

> LAN MAC，在机器盒子背面就有
>
> WAN MAC，基本上是在 LAN MAC 最后加一位
>
> 2.4G MAC，同 LAN MAC
>
> 5G MAC，与 2.4G MAC 区别是第二位不同
>
> 获得的内容，到 ifconfig 去验证
> 



### 刷入固件

> 说明，网上一种做法是刷入 pb-boot ，然后路由器 Reset 之后，进入 192.168.1.1 ，出现一个界面，选择新固件上传等等，但是我试了两次都没成功，我选择直接刷入新固件

```shell
// 本机操作
scp -P 1022 pb-boot-hc5962.bin root@192.168.199.1:/tmp
scp -P 1022 B70_3.4.3.9-099_20170910-2224.trx root@192.168.199.1:/tmp

// 极路由操作

mtd write /tmp/B70_3.4.3.9-099_20170910-2224.trx firmware
mtd erase firmware_backup
mtd write /tmp/pb-boot-hc5962.bin u-boot

// 然后重启路由器，进入 192.168.5.1 ，admin - admin 就成功了
```



### MAC 修复

通过 `ssh admin@192.168.5.1 ` 登录 

```shell
lan_eeprom_mac xxx          //命令说明：写入 LAN MAC。
wan_eeprom_mac xxx          //命令说明：写入 WAN MAC。
radio2_eeprom_mac  xxx      //命令说明：写入 2.4G MAC。
radio5_eeprom_mac  xxx      //命令说明：写入 5G MAC。
sync                        //命令说明：保存设置
reboot                      //命令说明：重启路由器
```

### 设置中文 && 双清 

在 `Administration > System > SelectWebUI Language 路径下选择[简体中文]就行` 

在`系统管理 >恢复/导出/上传设置路径 `下在`路由器参数`和`路由器内部存储`下面的`恢复出厂设置`，点击“重置”按钮



### 开启 Google-fu  插件 

到 `路由器的管理页面--系统管理-- 控制台`

输入以下命令，一行一个，分次输入

```
nvram set google_fu_mode=0xDEADBEEF
nvram set ext_show_lse=1
nvram commit
```



### 参考

[极路由B70刷机教程 解锁breed刷入老毛子padavan固件](https://pannixilin.com/archives/119.html)

[极路由B70刷固件详细步骤说明](https://www.right.com.cn/forum/thread-250789-1-1.html)


## Google-fu 设置 

### chnroute mode 规则更新 

```shell
// 生成最新的 chnroute.txt  
curl 'http://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest' | grep ipv4 | grep CN | awk -F\| '{ printf("%s/%d\n", $4, 32-log($5)/log(2)) }' > chnroute.txt  

// 生成 md5sum 文件  
md5sum chnroute.txt > chnroute.txt.md5sum.txt  

// ps，如果是 mac 需要先 `brew install md5sha1sum`  
```

使用 `https://raw.githubusercontent.com/dryyun/gather/master/hiwifi_padavan/chnroute.txt`

### gfwlist mode 规则更新

```shell

// 生成最新的 plain   glwlist.txt
curl https://raw.githubusercontent.com/gfwlist/gfwlist/master/gfwlist.txt  | base64 --decode  > ss.gfwlist.lite.txt.plain

curl  https://raw.githubusercontent.com/gfwlist/gfwlist/master/gfwlist.txt | base64 --decode  > ss.gfwlist.lite.txt

// 生成 plain  md5sum 文件  

md5sum ss.gfwlist.lite.txt > ss.gfwlist.lite.txt.plain.md5sum.txt

// 下载最新的 decode 文件 
curl https://raw.githubusercontent.com/gfwlist/gfwlist/master/gfwlist.txt  > ss.gfwlist.lite.txt

// 生成 decode 文件对应的   md5sum 文件  
md5sum ss.gfwlist.lite.txt > ss.gfwlist.lite.txt.md5sum.txt

```

使用 `https://raw.githubusercontent.com/dryyun/gather/master/hiwifi_padavan/ss.gfwlist.lite.txt`  



