package iptables

import (
	"fmt"

	"github.com/spf13/cobra"
)

func InitMain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iptables [no options!]",
		Short: `iptables command examples`,
		Long:  `iptables command examples`,
		Example: `
### Example commands for iptables
networking-tools iptables
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
######### iptables commands #########
			`)
		},
	}

	return cmd
}
