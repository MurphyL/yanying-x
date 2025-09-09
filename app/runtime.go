package app

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var profile =

// Vars - 运行时参数
var Vars struct {
	WorkDir          string `json:"work_dir"`
	CurrentProjectId uint64 // 当前正在编辑的项目ID
	CountOfProjects  uint   // 当前的项目数量
}

func init() {
	ReloadUserProfile()
}

// ReloadUserProfile - 重载用户配置文件
func ReloadUserProfile() {
	filepath := filepath.Join(os.Getenv("USERPROFILE"), "yanying-x.json")
	bytes, _ := os.ReadFile(filepath)
	json.Unmarshal(bytes, &Vars)
}

func DumpUserProfile() {
	bytes, _ := json.Marshal(Vars)
	os.WriteFile()
}
