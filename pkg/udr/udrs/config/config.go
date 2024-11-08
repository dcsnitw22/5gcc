package config

import (
	"bytes"
	"fmt"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"

	"k8s.io/klog"
)

const (
	// DefaultConfigFilePath = "/go/src/w5gc.io/wipro5gcore/configs/smf"
	DefaultConfigFilePath = "/5gcc/configs/udr"
	DefaultConfigFileName = "udrs"
	DefaultEtcdConfigKey  = "/w5gc/config/udr/udrs.json"
	DefaultEtcdServer     = "http://localhost:2379"
	DefaultEtcdConfigType = "json"
	DefaultConfigType     = "json"
)

type UdrsConfig struct {
	Version     string
	NodeInfo    UdrNodeInfo
	N11AmfNodes []N11AmfNodeInfo
}

type UdrNodeInfo struct {
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

var defaultUdrsConfig = []byte(`
{
    "Version": "1.0",
    "NodeInfo": {
        "NodeId": "127.0.0.1",
        "ApiPort": ":8082"
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

var ConfigChannel = make(chan UdrsConfig)

var UdrsCfg, UdrsRuntimeCfg UdrsConfig

func InitConfig(cfgFile string, etcdServer string, etcdConfigKey string, resetFlag bool) (*UdrsConfig, error) {
	runtime_viper := viper.New()
	runtime_viper.Set("Verbose", true)
	var etcdConfig bool

	// Initialize ConfigChannel
	ConfigChannel = make(chan UdrsConfig)

	// Try to configure from etcd if resetFlag is false
	if !resetFlag {
		if etcdServer == "" {
			klog.Errorf("Udrs configuration using default etcd server %s", DefaultEtcdServer)
			etcdServer = DefaultEtcdServer
		}
		if etcdConfigKey == "" {
			klog.Errorf("Udrs configuration using default etcd server key %s", DefaultEtcdConfigKey)
			etcdConfigKey = DefaultEtcdConfigKey
		}

		runtime_viper.AddRemoteProvider("etcd", etcdServer, etcdConfigKey)
		runtime_viper.SetConfigType(DefaultEtcdConfigType)
		err := runtime_viper.ReadRemoteConfig()

		if err == nil {
			klog.Info("Udrs configured using etcd")
			etcdConfig = true
		} else {
			klog.Errorf("Udrs configuration using etcd failed: %v", err)
		}
	}

	if !etcdConfig {
		klog.Info("Udrs configuration using config file")

		if cfgFile != "" {
			runtime_viper.SetConfigFile(cfgFile)
		} else {
			home, err := homedir.Dir()
			if err != nil {
				klog.Errorf("Unable to get home directory: %s", err.Error())
				return nil, err
			}

			runtime_viper.AddConfigPath(home + DefaultConfigFilePath)
			runtime_viper.SetConfigName(DefaultConfigFileName)
		}

		runtime_viper.AutomaticEnv()
		klog.Info("Reading config file")
		if err := runtime_viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				klog.Error("Config file not found in viper")
			} else {
				klog.Error("Unable to read config file in viper")
			}
			klog.Errorf("Unable to read config file, setting default values: %s", err.Error())
			runtime_viper.SetConfigType(DefaultConfigType)
			if err := runtime_viper.ReadConfig(bytes.NewBuffer(defaultUdrsConfig)); err != nil {
				klog.Errorf("Unable to read config file with default values: %s", err.Error())
				return nil, err
			}
		}

		klog.Info("Unmarshalling config file")
	}

	// Unmarshal config into UdrsCfg
	err := runtime_viper.Unmarshal(&UdrsCfg)
	if err != nil {
		klog.Fatalf("Unable to decode Udrs config into struct: %v", err)
		return nil, err
	}

	// Ensure UdrsCfg is not nil before proceeding
	if UdrsCfg.Version == "" {
		klog.Error("UdrsCfg is not properly initialized")
		return nil, fmt.Errorf("UdrsCfg is not properly initialized")
	}
	klog.Info("UdrsCfg initialized successfully")
	// Start a goroutine to watch remote config changes if etcdConfig is true
	if etcdConfig {
		go func() {
			for {
				time.Sleep(time.Second * 5)
				err := runtime_viper.WatchRemoteConfig()
				if err != nil {
					klog.Errorf("Unable to read remote config: %v", err)
					continue
				}
				err = runtime_viper.Unmarshal(&UdrsRuntimeCfg)
				if err != nil {
					klog.Errorf("Unable to decode remote config into struct: %v", err)
					continue
				}
				ConfigChannel <- UdrsRuntimeCfg
			}
		}()
	}
	klog.Info("passed etcconfig not found test")
	// Send initial configuration to ConfigChannel
	go func() {
		time.Sleep(time.Second * 5)
		klog.Info("Sending initial configuration to ConfigChannel")
		ConfigChannel <- UdrsCfg
	}()
	return &UdrsCfg, nil
}
