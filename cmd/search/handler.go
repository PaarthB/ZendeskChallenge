// Package search -
//
// Defines all the commands for searching, and which entrypoints to invoke for them
//

package search

import (
	"github.com/spf13/cobra"
	_ "time"
)

// NewSearchCmd - Parent command setup for all search commands (user, ticket, organization) /*
func NewSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "Zendesk Search",
	}
	cmd.AddCommand(NewUserSearchCmd())
	cmd.AddCommand(NewOrgSearchCmd())
	cmd.AddCommand(NewTicketSearchCmd())

	return cmd
}

// NewUserSearchCmd - Define user search command /*
func NewUserSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "trigger user search",
		RunE:  triggerUserSearch,
	}

	cmd.PersistentFlags().String("name", "", "The name of the field to search for.")
	cmd.PersistentFlags().String("value", "", "Name of the field to search for")
	_ = cmd.MarkPersistentFlagRequired("name")
	//_ = cmd.MarkPersistentFlagRequired("value")
	return cmd
}

// NewTicketSearchCmd - Define ticket search command /*
func NewTicketSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ticket",
		Short: "trigger ticket search",
		RunE:  triggerTicketSearch, // method to run when ticket search is triggered by user
	}

	cmd.PersistentFlags().String("name", "", "The name of the field to search for.")
	cmd.PersistentFlags().String("value", "", "Name of the field to search for")
	_ = cmd.MarkPersistentFlagRequired("name")
	//_ = cmd.MarkPersistentFlagRequired("value")
	return cmd
}

// NewOrgSearchCmd - Define organization search command /*
func NewOrgSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "organization",
		Short: "trigger organization search",
		RunE:  triggerOrgSearch, // method to run when organization search is triggered by user
	}

	cmd.PersistentFlags().String("name", "", "The name of the field to search for.")
	cmd.PersistentFlags().String("value", "", "Name of the field to search for")
	_ = cmd.MarkPersistentFlagRequired("name")
	//_ = cmd.MarkPersistentFlagRequired("value")
	return cmd
}
