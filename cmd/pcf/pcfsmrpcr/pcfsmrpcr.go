package main

import (
	"os"

	"k8s.io/klog"

	"w5gc.io/wipro5gcore/cmd/pcf/pcfsmrpcr/app"
)

func main() {
	rootCmd := app.NewPcfsmrpcrRootCommand()
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
	klog.Info("UDM SDMS Stopped")

}
