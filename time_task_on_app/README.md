# 实现 Mac 上定时开启、关闭某个 App 的小功能

## 知识点汇总

初略的看了 AppleScript 脚本语法，osascript 命令使用，launchctl 定时任务配置

## 文件介绍
- killapp.scpt 是练手 AppleScript 脚本文件
    - 可以判断某个 app 是否在运行，在运行就使其退出
    - 获取 app 的名称，查看 `/Applications` 目录下的名称即可
- handleApp.sh 是具体执行的 shell 文件
    - 需要 chmod +x ，加上可执行权限
    - 判断时间，上午开启应用，傍晚关闭应用
    - 主要依赖 `osascript` 命令
- com.handleApp.plist 是 launchctl 的配置文件
    - 设置某些时间点执行，类似 crontab 吧
    - 需要 `mv ~/Library/LaunchAgents` 目录下，用于用户执行
    - 加载 `launchctl load -w  ~/Library/LaunchAgents/com.handleApp.plist` 
    - 查看 `launchctl list |grep handle`
    - 卸载 `launchctl unload -w  ~/Library/LaunchAgents/com.handleApp.plist` 
    - 每次修改了 plist 文件，都需要 unload - load 操作一遍，并没有找到 reload 


