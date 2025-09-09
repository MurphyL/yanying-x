package projects

import (
	"fmt"
	"strconv"
	"yanying-x/app"
	"yanying-x/misc/cli"
)

type ProjectInfo struct {
	Id string
}

func DisplayProjectInfo() {
	if app.Vars.CurrentProjectId == 0 {
		fmt.Printf("暂未指定任何项目，请选择：")
		line := cli.ReadLine()
		projectId, _ := strconv.ParseUint(line, 10, 32)
		app.Vars.CurrentProjectId = projectId
	} else {
		fmt.Sprintf("当前项目【%v】\n", app.Vars.CurrentProjectId)
	}
}
