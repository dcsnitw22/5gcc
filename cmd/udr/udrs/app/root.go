package app

import (
	"github.com/spf13/cobra"
)

func NewUdrsRootCommand() *cobra.Command {

	var cfgFile, etcdServer, etcdConfigKey string
	var reset bool

	rootCmd := &cobra.Command{
		Use:   "udrs",
		Short: "udrs is a UDR microservice in 5GC ",
		Long: `udrs provides the session management services in the smf
                        Complete documentation is available at http://wipro.com/5gc`,
		Run: func(cmd *cobra.Command, args []string) {
			RunUdrs(cfgFile, etcdServer, etcdConfigKey, reset)
		},
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is w5gc.io/wipro5gcore/configs/udrs)")
	rootCmd.PersistentFlags().StringVar(&etcdServer, "etcd-server", "", "etcd server to read config file (default is /w5gc/config/udrs.json)")
	rootCmd.PersistentFlags().StringVar(&etcdConfigKey, "etcd-config-key", "", "etcd server key for config file (default is /w5gc/config/udrs.json)")
	rootCmd.PersistentFlags().BoolVarP(&reset, "reset", "r", false, "reset flag")
	rootCmd.PersistentFlags().StringP("author", "a", "Wipro", "author name for copyright attribution")

	versionCmd := NewUdrsVersionCommand()

	rootCmd.AddCommand(versionCmd)

	return rootCmd
}
