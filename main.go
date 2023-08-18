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

func withNodeModules(dirs []string) []string {
	for _, v := range os.Args {
		if v == "-n" || v == "--node_modules" {
			return append(dirs, "node_modules")
		}
	}
	return dirs
}

func main() {
	if isHelp() {
		fmt.Print("\n")
		color.Cyan("clean-cache\n\n")
		fmt.Print("Description: go 写的清理 node 项目缓存，超级无敌快\n\n")
		fmt.Print("Usage: -n --node_modules 则会删除 node_modules\n\n")
		os.Exit(0)
	}

	cache_dirs := []string{".nuxt", "cache", ".cache", "@cache", "temp", ".temp", "@temp", "dist", ".output"}

	cache_dirs = withNodeModules(cache_dirs)

	for _, v := range cache_dirs {
		if !IsExist(v) || !IsDir(v) {
			continue
		}
		err := RemoveAll(v)
		if err != nil {
			color.Red("remove fail: %v \nroot: %v", err, v)
		} else {
			color.Green("remove success: %v", v)
		}
	}
}
