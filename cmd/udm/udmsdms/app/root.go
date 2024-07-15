package app

import (
	"github.com/spf13/cobra"
)

func NewUdmsdmsRootCommand() *cobra.Command {

	var cfgFile, etcdServer, etcdConfigKey string
	var reset bool

	rootCmd := &cobra.Command{
		Use:   "udmsdms",
		Short: "udmsdms is a SMF microservice in 5GC to handle pdu sms signalling",
		Long: `udmsdms provides the session management services in the smf
                        Complete documentation is available at http://wipro.com/5gc`,
		Run: func(cmd *cobra.Command, args []string) {
			RunUdmsdms(cfgFile, etcdServer, etcdConfigKey, reset)
		},
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is w5gc.io/wipro5gcore/configs/udmsdms)")
	rootCmd.PersistentFlags().StringVar(&etcdServer, "etcd-server", "", "etcd server to read config file (default is /w5gc/config/udmsdms.json)")
	rootCmd.PersistentFlags().StringVar(&etcdConfigKey, "etcd-config-key", "", "etcd server key for config file (default is /w5gc/config/udmsdms.json)")
	rootCmd.PersistentFlags().BoolVarP(&reset, "reset", "r", false, "reset flag")
	rootCmd.PersistentFlags().StringP("author", "a", "Wipro", "author name for copyright attribution")

	versionCmd := NewUdmsdmsVersionCommand()

	rootCmd.AddCommand(versionCmd)

	return rootCmd
}
