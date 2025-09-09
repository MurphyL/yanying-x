package main

import (
	"fmt"
	"os"
	"strings"
	"yanying-x/biz/projects"
	"yanying-x/misc/cli"
)

func main() {
	//fmt.Printf("工作目录【%s】\n", app.WorkDir())
	projects.DisplayProjectInfo()
	for {
		fmt.Print("请输入指令：")
		line := cli.ReadLine()
		if len(line) > 0 {
			switch strings.TrimSpace(line) {
			case "exit":
				fmt.Println("退出")
				os.Exit(0)
			case "list":
				fmt.Println("列出所有之类：", line)
			default:
				fmt.Println("执行命令：", line)
			}
		}
	}
}
