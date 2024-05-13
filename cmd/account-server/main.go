package main

import (
	"fmt"
	"github.com/Superm4n97/account-server/pkg/cmds"
	"gomodules.xyz/logs"
	"os"
)

func main() {
	if err := execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func execute() error {
	cmd := cmds.NewRootCmd()
	logs.Init(cmd, true)
	defer logs.FlushLogs()

	return cmd.Execute()
}
