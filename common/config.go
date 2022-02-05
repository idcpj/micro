package common

import (
	"github.com/asim/go-micro/plugins/config/source/consul/v4"
	"go-micro.dev/v4/config"
	"strconv"
)

func GetConsulConfig(host string,port int64,prefix string)(config.Config,error) {

	var source = consul.NewSource(
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		consul.WithPrefix(prefix),
		consul.StripPrefix(true),
	)
	newConfig, err := config.NewConfig()
	if err != nil {
		return nil,err
	}
	err = newConfig.Load(source)
	if err != nil {
		return nil,err
	}

	return newConfig,nil
}
