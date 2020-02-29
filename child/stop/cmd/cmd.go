package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"iboxctl/pkg/root"
	"iboxctl/pkg/tools"
	"iboxctl/pkg/udpsend"
)
var(
	stop        = make(chan struct{})
	ch string
	ip string
	ipCmd      = &cobra.Command{
		Use:   "stop",
		Short: "stop server",
		RunE: func(c *cobra.Command, args []string) error {
			tools.PrintFlags(c.Flags())
			fmt.Printf("stop server %s\n",ch)
			udpsend.Exec("pwr_ch"+ch+"_off")

			return nil
		},
	}
)


func init(){
	ipCmd.PersistentFlags().StringVar(&ch, "ch", "","server's id")
	//ipCmd.PersistentFlags().StringVar(&ip, "ip", "","server's id")
	root.AddCmd(ipCmd)
}