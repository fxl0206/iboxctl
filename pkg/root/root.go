package root

import (
	"github.com/spf13/cobra"
	"iboxctl/pkg/tools"
	"log"
	"os"
)

var(
	RootCmd = &cobra.Command{
		Use:          "ibox",
		Short:        "ibox",
		Long:         "ibox start/stop/status",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			tools.PrintFlags(c.Flags())
			return nil
		},
	}
)



func Run(){
	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}

func AddCmd(cmds ...*cobra.Command){
	for i,_:=range cmds {
		RootCmd.AddCommand(cmds[i])
	}
}

func SetArgs(args []string){
	RootCmd.SetArgs(args)
}