package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var ignoreSettingDirs []string // 忽略的文件夹名称，精确匹配
//var ignoreSettingFiles []string                      // 忽略的文件名称，模糊匹配
var ignoreDirs []string // 全部忽略的文件夹
//var ignoreFiles []string                             // 全部忽略的文件
var searchDepth = 3                                  // 查找文件目录深度
var baseDir string                                   // 查找起始目录
var rulesConf = "~/.nutstore/db/customExtRules.conf" // 坚果云的忽略规则文件

var (
	setBaseDir  = flag.String("root", ".", "设置查找根目录，默认当前文件夹")
	setDepth    = flag.Int("depth", 3, "设置查找文件夹深度")
	setDirs     = flag.String("dirs", ".idea,node_modules,vendor", "设置忽略文件夹，使用[,] 间隔")
	setRuleFile = flag.String("rulefile", rulesConf, "规则文件路径")
	setReadonly = flag.Bool("readonly", true, "只查看结果，不做任何操作")
)

func main() {
	flag.Parse()

	searchDepth = *setDepth

	ignoreSettingDirs = strings.Split(*setDirs, ",")

	baseDir = *setBaseDir
	baseDir, _ = realPath(baseDir)

	rulesConf = *setRuleFile
	rulesConf, _ = realPath(rulesConf)

	filepath.Walk(baseDir, walkFn)

	for _, dir := range ignoreDirs {
		fmt.Println(string(dir))
	}

	if *setReadonly {
		fmt.Println("do nothing, read result only")
		return
	}

	// 没有匹配规则，保持原来的不变
	if len(ignoreDirs) == 0 {
		return
	}

	var err error
	if isFileExist(rulesConf) {
		err = os.Rename(rulesConf, rulesConf+".backup")
		if err != nil {
			fmt.Println("backup "+rulesConf+" failed", err)
			return
		}
	}

	f, err := os.Create(rulesConf)
	if err != nil {
		fmt.Println("create "+rulesConf+" failed", err)
	}
	for _, dir := range ignoreDirs {
		f.WriteString(string(dir + "\n"))
	}
	f.Close()
	fmt.Println("create " + rulesConf + " success")
}

/**
判断文件是否存在
*/
func isFileExist(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

/*
获取文件真实路径，支持 [ ~ ]
*/
func realPath(path string) (rp string, err error) {
	u, err := user.Current()
	if err != nil {
		return
	}
	dir := u.HomeDir
	path = strings.Trim(path, " ")

	if path == "~" || path == "~/" || path == "~/." {
		rp = dir
	} else if strings.HasPrefix(path, "~/") {
		rp = filepath.Join(dir, path[2:])
	}
	rp, err = filepath.Abs(rp)
	return
}

/**
判断 slice 中是否存在元素 e
*/
func sliceContains(a []string, e string) bool {
	for _, n := range a {
		if e == n {
			return true
		}
	}
	return false
}

func walkFn(path string, f os.FileInfo, err error) error {
	depth := strings.Count(path, "/") - strings.Count(baseDir, "/")
	if depth > searchDepth {
		return filepath.SkipDir
	}

	if f.IsDir() && sliceContains(ignoreSettingDirs, f.Name()) {
		ignoreDirs = append(ignoreDirs, path)
		return filepath.SkipDir
	}
	return nil
}
