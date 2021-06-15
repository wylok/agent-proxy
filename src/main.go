package main

import (
	_ "bytes"
	_ "github.com/CodyGuo/godaemon"
	"github.com/mitchellh/go-ps"
	"io/ioutil"
	"os"
	"proxy/config"
	"proxy/kits"
	"proxy/modules"
)

func main() {
	// 检查proxy进程是否重复启动
	if kits.CheckFile(config.PidFile) {
		f, err := ioutil.ReadFile(config.PidFile)
		if err != nil {
			kits.Log(err.Error(), "error", "main")
		}
		p, err := ps.FindProcess(int(kits.BytesToInt64(f)))
		if p != nil {
			println("agent-proxy进程已运行!")
			os.Exit(1)
		}
	}
	pid := os.Getpid()
	err := ioutil.WriteFile(config.PidFile, kits.Int64ToBytes(int64(pid)), 0666)
	if err != nil {
		kits.Log(err.Error(), "error", "main")
	}
	modules.ProxyServer()
}
