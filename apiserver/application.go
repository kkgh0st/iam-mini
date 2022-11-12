package apiserver

import (
	"github.com/spf13/cobra"
)

type app struct {
	*cobra.Command
}

type cmdRunFunc func(cmd *cobra.Command, args []string)

func apiServerApp() *app {
	return &app{&cobra.Command{
		Use:   "Use",
		Short: "Short",
		// err := _app.Cmd.Execute() 其调用 Execute()时自动到这个函数
		Run: cmdRun(),
	}}
}
