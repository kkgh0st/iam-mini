package opt

import "github.com/spf13/pflag"

// 目前就要来实现这两个接口
type option interface {
	AddFlags() *pflag.FlagSet
	Validate() bool
}
