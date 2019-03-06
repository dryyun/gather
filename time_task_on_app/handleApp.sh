#!/bin/bash
# path 需要替换为真实路径
# appName 需要替换为真实 App，只需要替换 set appName to "Safari" 这里的 appName 即可

hour=`date +%H`
if [[ "$hour" > 18 ]]; then
    osascript /path/killapp.scpt 
elif [[ "$hour" < 12 ]]; then 
    osascript -e '
    on is_running(appName)
        tell application "System Events" to (name of processes) contains appName
    end is_running
    set appName to "Safari"
    set isRunning to is_running(appName)
    if isRunning then
        log (appName & " is runnning  ")
    else
        log (appName & " is not running to activate")
        tell application appName to activate
    end if
'
else
    echo "nothing"
fi
