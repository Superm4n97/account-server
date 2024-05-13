package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	var (
		port string
	)
	cmd := &cobra.Command{
		Use:   "account-server",
		Short: "Launch Account Server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello world")
			fmt.Println(port)
		},
	}
	cmd.Flags().StringVar(&port, "port", ":8080", "The port server exposed to.")
	return cmd
}
