package main

import (
	"flag"
	"path"
	"strings"
)

func main() {
	GoFileName := ""
	OutDir := ""
	OutPackageName := ""
	ImportCode := ""
	InPackageName := ""
	AutoInit := false
	flag.StringVar(&GoFileName, "GoFileName", "", "输入需要生成绑定的Go文件名")
	flag.StringVar(&OutDir, "OutDir", "", "生成绑定文件的文件夹")
	flag.StringVar(&OutPackageName, "OutPackageName", "", "生成绑定文件包名")
	flag.StringVar(&ImportCode, "ImportCode", "", "需要import的部分")
	flag.StringVar(&InPackageName, "InPackageName", "", "目标文件的包名")
	flag.BoolVar(&AutoInit, "AutoInit", false, "是否添加自动初始化代码")
	flag.Parse()
	{
		//GoFileName = `ttlua.go`
		//OutDir = ""luaUserData
		//OutPackageName = "main"
		//AutoInit = true
	}

	GoFileName = strings.Replace(GoFileName, `\`, `/`, -1)
	GoFileName = strings.Replace(GoFileName, `\`, `/`, -1)
	if OutDir == "" {
		OutDir = path.Dir(GoFileName)
	}
	Cfg.OutPackageName = OutPackageName
	Cfg.AutoInit = AutoInit
	Cfg.ImportString = strings.Split(ImportCode, ";")
	Cfg.InPackageName = InPackageName
	MakeLuaBindFile(GoFileName, OutDir)
	return
}
