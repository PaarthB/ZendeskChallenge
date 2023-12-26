package list

import (
	"github.com/spf13/cobra"
	_ "time"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List searchable fields for Zendesk CLI",
		Run:   fieldList,
	}
	return cmd
}
