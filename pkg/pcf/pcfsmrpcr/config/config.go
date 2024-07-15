package config

import (
	"bytes"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"

	"k8s.io/klog"
)

const (
	// DefaultConfigFilePath = "/go/src/w5gc.io/wipro5gcore/configs/smf"
	DefaultConfigFilePath = "/wipro5gc/configs/pcf"
	DefaultConfigFileName = "pcfsmrpcr.json"
	DefaultEtcdConfigKey  = "/w5gc/config/pcf/pcfsmrpcr.json"
	DefaultEtcdServer     = "http://localhost:2379"
	DefaultEtcdConfigType = "json"
	DefaultConfigType     = "json"
)

type PcfsmrpcrConfig struct {
	Version     string
	NodeInfo    PcfNodeInfo
	N11AmfNodes []N11AmfNodeInfo
}

type PcfNodeInfo struct {
	NodeId          string
	ApiPort         string
	LoadControl     bool
	OverloadControl bool
}

type N11AmfNodeInfo struct {
	NodeId string
	Port   string
}

type Update struct {
}

var defaultPcfsmrpcrConfig = []byte(`
{
    "Version": "1.0",
    "NodeInfo": {
        "NodeId": "127.0.0.1",
        "ApiPort": ":8084"
    },
    "N11AmfNodes": [
        {
            "NodeId": "127.0.0.1", 
            "Port": "8083"
        },
        {
            "NodeId": "10.250.108.35",
            "Port": "8080"
        }
    ]
}`)

var ConfigChannel chan PcfsmrpcrConfig
var PcfsmrpcrCfg, PcfsmrpcrRuntimeCfg PcfsmrpcrConfig

func InitConfig(cfgFile string, etcdServer string, etcdConfigKey string, resetFlag bool) (*PcfsmrpcrConfig, error) {

	runtime_viper := viper.New()
	runtime_viper.Set("Verbose", true)
	//viper.Set("LogFile", LogFile)
	var etcdConfig bool

	// If Pcfsmrpcr is not reset try to configure from etcd
	if resetFlag != true {
		if etcdServer == "" {
			klog.Errorf("Pcfsmrpcr configuration using default etcd server %s", DefaultEtcdServer)
			etcdServer = DefaultEtcdServer
		}
		if etcdConfigKey == "" {
			klog.Errorf("Pcfsmrpcr configuration using default etcd server key %s", DefaultEtcdConfigKey)
			etcdConfigKey = DefaultEtcdConfigKey
		}

		runtime_viper.AddRemoteProvider("etcd", etcdServer, etcdConfigKey)
		runtime_viper.SetConfigType(DefaultEtcdConfigType)
		err := runtime_viper.ReadRemoteConfig()

		if err == nil {
			klog.Info("Pcfsmrpcr configured using etcd")
			etcdConfig = true
		}

		klog.Error("Pcfsmrpcr configuration using etcd failed")
	}

	if !etcdConfig {
		klog.Info("Pcfsmrpcr configuration using config file")

		// Check for config file parameter
		if cfgFile != "" {
			runtime_viper.SetConfigFile(cfgFile)
		} else {
			// Find home directory.
			home, err := homedir.Dir()
			if err != nil {
				klog.Errorf("[Unable to get home directory] %s", err.Error())
			}

			// Search config in home directory with name "go/src/w5gc.io/wipro5gcore/configs/" .
			klog.Info(home)
			runtime_viper.AddConfigPath(home + DefaultConfigFilePath)
			runtime_viper.SetConfigName(DefaultConfigFileName)
		}

		runtime_viper.AutomaticEnv()

		if err := runtime_viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				klog.Error("[Config file not found in viper]")
			} else {
				klog.Error("[Unable to read config file in viper]")
			}
			klog.Errorf("[Unable to read config file, setting default values] %s", err.Error())
			runtime_viper.SetConfigType(DefaultConfigType)
			if err := runtime_viper.ReadConfig(bytes.NewBuffer(defaultPcfsmrpcrConfig)); err != nil {
				klog.Errorf("[Unable to read config file with default values] %s", err.Error())
				os.Exit(1)
			}
		}
	}

	// Need to set default vaues TODO GURU

	// unmarshal config in viper to  config
	err := runtime_viper.Unmarshal(&PcfsmrpcrCfg)
	if err != nil {
		klog.Fatalf("Unable to decode pdusmsp config into struct, %v", err)
		os.Exit(1)
	}

	// Write config to remote TODO GURU
	// klog.Info(etcdConfig)
	// etcdConfig = true
	if etcdConfig {
		ConfigChannel = make(chan PcfsmrpcrConfig)

		// Start a goroutine to watch remote config changes forever
		// Watch changes for config file? TODO GURU
		go func() {
			for {
				// delay after each request
				time.Sleep(time.Second * 5)

				// currently, only tested with etcd support
				err := runtime_viper.WatchRemoteConfig()
				if err != nil {
					klog.Errorf("unable to read remote pdusmsp config: %v", err)
					continue
				}

				// unmarshal new config into our runtime config struct
				runtime_viper.Unmarshal(&PcfsmrpcrRuntimeCfg)

				// Notify pdusmsp event handler of the change
				ConfigChannel <- PcfsmrpcrRuntimeCfg
			}
		}()
	}
	// temperory| jugad
	go func() {
		ConfigChannel = make(chan PcfsmrpcrConfig)
		klog.Info(PcfsmrpcrCfg)
		ConfigChannel <- PcfsmrpcrCfg
	}()
	return &PcfsmrpcrCfg, err
}
