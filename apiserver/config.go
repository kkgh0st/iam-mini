package apiserver

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"iam-mini/apiserver/server"
	"os"
)

func readConfigFile() *server.Opts {

	viper.SetConfigFile(ConfigFilePath)
	viper.SetConfigType("yaml")

	opts := server.NewOpts()

	fs := &pflag.FlagSet{}
	opts.AddFlags(fs)
	//fss.SetNormalizeFunc(func(f *pflag.FlagSet, name string) pflag.NormalizedName {
	//	return name
	//})

	err := viper.BindPFlags(fs)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(2)
	}

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(2)
	}

	//
	err = viper.Unmarshal(&opts)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(2)
	}

	return opts
}

func DoConfig() *server.Opts {
	return readConfigFile()
}
