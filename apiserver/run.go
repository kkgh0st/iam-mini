package apiserver

import (
	"github.com/spf13/cobra"
	"iam-mini/apiserver/server"
)

func cmdRun() cmdRunFunc {
	// 这里需要我们来定义相关匿名函数
	return func(cmd *cobra.Command, args []string) {
		run() // 来配置相关参数
	}
}

func run() {

	// 下面来完成这些工作，这本身一点也不难
	opts := DoConfig()
	serverConfig := server.NewConfig()
	opts.ApplyTo(serverConfig)
	config, err := CreateConfigFromOptions(opts)
	if err != nil {
		return
	}
	apiServer, err := CreateAPIServer(config)
	if err != nil {
		return
	}

	err = apiServer.PrepareRun().Run()
	if err != nil {
		return
	}
}
