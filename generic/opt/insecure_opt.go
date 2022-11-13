package opt

import (
	"github.com/spf13/pflag"
)

type InSecureServingInfo struct {
	BindAddress string
	BindPort    int
}

type InsecureOpt struct {
	BindAddress string `json:"bind"                     mapstructure:"bind"`
	BindPort    int    `json:"bind-port"    mapstructure:"bind-port"`
}

func (o *InsecureOpt) ApplyTo(info *InSecureServingInfo) {
	info.BindAddress = o.BindAddress
	info.BindPort = o.BindPort

}

func (o *InsecureOpt) Validate() bool {
	return true
}

func (o *InsecureOpt) AddFlags(fss *pflag.FlagSet) {
	fss.StringVar(&o.BindAddress, "insecure.bind-address", o.BindAddress, ""+
		"The IP address on which to serve the --insecure.bind-port "+
		"(set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")
	fss.IntVar(&o.BindPort, "insecure.bind-port", o.BindPort, ""+
		"The port on which to serve unsecured, unauthenticated access. It is assumed "+
		"that firewall rules are set up such that this port is not reachable from outside of "+
		"the deployed machine and that port 443 on the iam public address is proxied to this "+
		"port. This is performed by nginx in the default setup. Set to zero to disable.")
}

func NewInsecureOpt() *InsecureOpt {
	return &InsecureOpt{
		BindAddress: "127.0.0.1",
		BindPort:    8080,
	}
}
