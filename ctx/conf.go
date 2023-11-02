package ctx

import (
	"fmt"

	conf "github.com/nixys/nxs-go-conf"
)

type confOpts struct {
	LogFile  string `conf:"logfile" conf_extraopts:"default=stdout"`
	LogLevel string `conf:"loglevel" conf_extraopts:"default=info"`

	CounterInterval string `conf:"counterInterval" conf_extraopts:"default=1s"`
}

func confRead(confPath string) (confOpts, error) {

	var c confOpts

	err := conf.Load(&c, conf.Settings{
		ConfPath:    confPath,
		ConfType:    conf.ConfigTypeYAML,
		UnknownDeny: true,
	})
	if err != nil {
		return c, fmt.Errorf("conf init: %w", err)
	}

	return c, nil
}
