package apiserver

import (
	"github.com/spf13/pflag"
)

var ConfigFilePath string = ""

func init() {
	pflag.StringVar(&ConfigFilePath, "config", "", "the absolute path for api-server's config file")
}

func Run() {
	pflag.Parse()

	_app := apiServerApp()
	err := _app.Execute()
	if err != nil {
		return
	}
}
