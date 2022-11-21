package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"math/rand"
	"os"
	"time"
)

var App *Conf

func init() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	App = loadConfig()
}

// Service 服务配置
type Service struct {
	Host string `yaml:"host"` // 服务监听地址
	Port string `yaml:"port"` // 服务监听端口
	Path string `yaml:"path"` // 文件目录
}

type Conf struct {
	Service `yaml:"service"`
}

// 加载 app.yaml 配置
func loadConfig() *Conf {
	path := flag.String("c", "./app.yaml", "输入 -c xxx.yaml 自定义配置文件")
	flag.Parse()
	file, e := os.ReadFile(*path)
	if e != nil {
		panic(e)
	}

	var app Conf
	e = yaml.Unmarshal(file, &app)
	if e != nil {
		panic(e)
	}
	fmt.Println("🔨 Config -> " + *path)
	return &app
}
