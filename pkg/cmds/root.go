package cmds

import (
	"github.com/Superm4n97/account-server/server"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func NewRootCmd() *cobra.Command {
	var (
		port   string
		dbUri  string
		dbName string
	)
	cmd := &cobra.Command{
		Use:   "account-server",
		Short: "Launch Account Server",
		Run: func(cmd *cobra.Command, args []string) {
			s := server.Server{
				Router:       server.NewRouter(),
				Port:         port,
				DatabaseURI:  dbUri,
				DatabaseName: dbName,
			}
			if err := s.Start(); err != nil {
				klog.Error(err.Error())
			}
		},
	}
	cmd.Flags().StringVar(&port, "port", ":8080", "The port server exposed to.")
	cmd.Flags().StringVar(&dbUri, "db-uri", "mongodb://localhost:27017", "Database connection string.")
	cmd.Flags().StringVar(&dbName, "db-name", "account-server", "Mongodb database name.")
	return cmd
}
