package main

import (
	"os"

	"k8s.io/klog"

	"w5gc.io/wipro5gcore/cmd/udm/udmsdms/app"
)

func main() {
	rootCmd := app.NewUdmsdmsRootCommand()
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
	klog.Info("UDM SDMS Stopped")

}
