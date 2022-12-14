package config

import (
	"flag"
	"github.com/flametest/go-gin-server-example/pkg/mysql"
	"github.com/spf13/viper"
	"log"
	"path"
	"runtime"
)

var cfgPath = flag.String("config-path", "/opt/conf/", "configuration file path")
var cfgName = flag.String("config-name", "config", "configuration file name")

type Config struct {
	ProjectName string        `json:"projectName" yaml:"projectName"`
	Env         string        `json:"env" yaml:"env"`
	HttpHost    string        `json:"httpHost" yaml:"httpHost"`
	GrpcHost    string        `json:"grpcHost" yaml:"grpcHost"`
	Mysql       *mysql.Config `json:"mysql" yaml:"mysql"`
}

var Cfg = &Config{}

func InitConfig() {
	// server config file
	viper.SetConfigName(*cfgName)
	viper.SetConfigType("yml")
	viper.AddConfigPath(*cfgPath)

	// local config file
	// get caller file path
	_, fp, _, _ := runtime.Caller(0)
	curDir := path.Dir(fp)
	localCfgPath := path.Join(curDir, "../../deployment")
	viper.AddConfigPath(localCfgPath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("err when reading in config: %s", err)
	}

	err = viper.Unmarshal(Cfg)
	if err != nil {
		return
	}

}
