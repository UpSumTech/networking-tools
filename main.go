package main

import (
	"github.com/spf13/cobra"
	"github.com/sumanmukherjee03/networking-tools/cmd/iptables"
)

var (
	rootShortDesc = "networking-tools is a simple utilty to print use cases of iptables and iproute2 commands"
	rootLongDesc  = `networking tools is meant to configure boxes for different purposes based on networking requirements.
	It prints out iptables and iproute2 commands to help configure and debug network issues in a linux VM.`
)

func main() {
	var rootCmd = &cobra.Command{
		Use:              "networking-tools [sub]",
		Short:            rootShortDesc,
		Long:             rootLongDesc,
		TraverseChildren: true,
	}

	rootCmd.AddCommand(iptables.InitMain())
	rootCmd.Execute()
}
