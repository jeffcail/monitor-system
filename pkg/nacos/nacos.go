package nacos

import (
	"strings"

	"bz.service.cloud.monitoring/utils"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
)

var Nac *Nacos

type Nacos struct{}

// ParseConfig
func (v *Nacos) ParseConfig(ip string, port int, cfg string, config interface{}) {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: ip,
			Port:   uint64(port),
		},
	}
	client, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig: &constant.ClientConfig{
			TimeoutMs: 5000,
		},
		ServerConfigs: serverConfigs,
	})
	utils.CheckErr(err)
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: cfg,
		Group:  "monitor",
	})
	utils.CheckErr(err)
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(strings.NewReader(content))
	utils.CheckErr(err)
	err = viper.Unmarshal(&config)
	utils.CheckErr(err)
}
