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
		Use:   "start",
		Short: "start server",
		RunE: func(c *cobra.Command, args []string) error {
			tools.PrintFlags(c.Flags())
			fmt.Printf("start server %s\n",ch)
			//ping.StartPing(ip)
			udpsend.Exec("power_ch"+ch)
			return nil
		},
	}
)


func init(){
	ipCmd.PersistentFlags().StringVar(&ch, "ch", "","server's id")
	//ipCmd.PersistentFlags().StringVar(&ip, "ip", "","server's id")
	root.AddCmd(ipCmd)
}