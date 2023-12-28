package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func isHelp() bool {
	for _, v := range os.Args {
		if v == "-h" || v == "--help" || v == "help" {
			return true
		}
	}
	return false
}

func withNodeModules(dirs *[]string) {
	for _, v := range os.Args {
		if v == "-n" || v == "--node_modules" {
			*dirs = append(*dirs, "node_modules")
		}
	}
}

func withDist(dirs *[]string) {
	for _, v := range os.Args {
		if v == "-d" || v == "--dist" {
			*dirs = append(*dirs, "dist", ".output")
			return
		}
	}
}

func withAppend(dirs *[]string)  {
	for _, v := range os.Args {
		if v == "-r" || v == "--root" {
			*dirs = append(*dirs, "./")
			return
		}
	}
}

func main() {
	if isHelp() {
		fmt.Print("\n")
		color.Cyan("clean-cache\n\n")
		fmt.Print("Description: go 写的清理 node 项目缓存，超级无敌快\n\n")
		fmt.Print("Usage: -r --root 删除根目录\n\n")
		fmt.Print("		  -d --dist 则会删除 dist 和 .output\n\n")
		fmt.Print("       -n --node_modules 则会删除 node_modules\n\n")
		os.Exit(0)
	}

	cacheDirs := []string{".nuxt", "cache", ".cache", "@cache", "temp", ".temp", "@temp"}

	withAppend(&cacheDirs)

	withDist(&cacheDirs)

	withNodeModules(&cacheDirs)

	for _, dir := range cacheDirs {
		if !IsExist(dir) || !IsDir(dir) {
			continue
		}
		err := RemoveAll(dir)
		if err != nil {
			color.Red("remove fail: %v \nroot: %v", err, dir)
		} else {
			color.Green("remove success: %v", dir)
		}
	}
}
