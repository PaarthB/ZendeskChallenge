// Package list -
//
// Defines all the commands for list, and which entry-points to invoke for them
//

package list

import (
	"fmt"
	"github.com/spf13/cobra"
	_ "time"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List searchable fields for Zendesk CLI",
		Annotations: map[string]string{
			"args":   "",
			"output": "Zendesk CLI list command",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), cmd.Annotations["output"])
			return fieldList(cmd, args)
		},
	}
	return cmd
}
