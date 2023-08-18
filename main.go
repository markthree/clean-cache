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

func main() {
	if isHelp() {
		fmt.Print("\n")
		color.Cyan("clean-cache")
		fmt.Print("Description: go 写的清理 node 项目缓存，超级无敌快\n\n")
		os.Exit(0)
	}

	cache_dirs := []string{".nuxt", "cache", ".cache", "@cache", "temp", ".temp", "@temp"}

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
