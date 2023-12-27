package main

import (
	"ZendeskChallenge/cmd/list"
	"ZendeskChallenge/cmd/search"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "cli",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				err := cmd.Help()
				if err != nil {
					return err
				}
				os.Exit(0)
			}
			return nil
		},
	}
	cmd.AddCommand(search.NewSearchCmd())
	cmd.AddCommand(list.NewListCmd())
	return cmd
}

func main() {
	ll := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	if ll == "ERROR" {
		log.SetLevel(log.ErrorLevel)
	} else if ll == "DEBUG" {
		log.SetLevel(log.DebugLevel)
	} else if ll == "TRACE" {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if err := NewRootCmd().Execute(); err != nil {
		log.Error(err)
		os.Exit(101)
	}
}
