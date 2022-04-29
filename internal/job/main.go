package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"go-chat/internal/pkg/logger"
	"go-chat/internal/provider"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	providers := Initialize(ctx, provider.ReadConfig(parseConfigArg()))

	// 设置日志输出
	logger.SetOutput(providers.Config.GetLogPath(), "logger-cli")

	newApp(ctx, providers.Commands.SubCommands())
}

func newApp(ctx context.Context, commands []*cli.Command) {
	cmd := cli.NewApp()

	cmd.Name = "LumenIM 在线聊天"
	cmd.Usage = "命令行管理工具"

	cmd.Flags = []cli.Flag{
		&cli.StringFlag{Name: "config", Aliases: []string{"c"}, Value: "./config.yaml", Usage: "配置文件路径", DefaultText: "./config.yaml"},
	}

	cmd.Commands = commands

	if err := cmd.RunContext(ctx, os.Args); err != nil {
		fmt.Printf("Command Error : %s", err.Error())
	}
}

func parseConfigArg() string {
	var conf string
	flag.StringVar(&conf, "config", "./config.yaml", "配置文件路径")
	flag.StringVar(&conf, "c", "./config.yaml", "配置文件路径")
	flag.Parse()
	return conf
}
