# 关闭坚果云
osascript -e 'quit app "Nutstore"'

go-ignore-jianguoyun -depth=3 -dirs=".idea,node_modules,vendor" -root="~/Codes"

# 打开坚果云
osascript -e 'open app "Nutstore"'
