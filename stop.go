package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

var cmdStop = &Command{
	Exec:        runStop,
	UsageLine:   "stop [OPTIONS] SERVER [SERVER...]",
	Description: "Stop a running server",
	Help:        "Stop a running server.",
}

func init() {
	cmdStop.Flag.BoolVar(&stopT, []string{"t", "-terminate"}, false, "Stop and trash a server with its volumes")
	cmdStop.Flag.BoolVar(&stopHelp, []string{"h", "-help"}, false, "Print usage")
}

// Flags
var stopT bool    // -t flag
var stopHelp bool // -h, --help flag

func runStop(cmd *Command, args []string) {
	if stopHelp {
		cmd.PrintUsage()
	}
	if len(args) < 1 {
		cmd.PrintShortUsage()
	}

	has_error := false
	for _, needle := range args {
		server := cmd.GetServer(needle)
		action := "poweroff"
		if stopT {
			action = "terminate"
		}
		err := cmd.API.PostServerAction(server, action)
		if err != nil {
			if err.Error() != "server should be running" && err.Error() != "server is being stopped or rebooted" {
				log.Warningf("failed to stop server %s: %s", server, err)
				has_error = true
			}
		} else {
			fmt.Println(needle)
		}
		if has_error {
			os.Exit(1)
		}
	}
}
