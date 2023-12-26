package search

import (
	"github.com/spf13/cobra"
	_ "time"
)

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

func NewUserSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "trigger user search",
		Run:   triggerUserSearch,
	}

	cmd.PersistentFlags().String("name", "", "The name of the field to search for.")
	cmd.PersistentFlags().String("value", "", "Name of the field to search for")
	_ = cmd.MarkPersistentFlagRequired("name")
	//_ = cmd.MarkPersistentFlagRequired("value")
	return cmd
}

func NewTicketSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ticket",
		Short: "Trigger user search",
		Run:   triggerTicketSearch,
	}

	cmd.PersistentFlags().String("name", "", "The name of the field to search for.")
	cmd.PersistentFlags().String("value", "", "Name of the field to search for")
	_ = cmd.MarkPersistentFlagRequired("name")
	//_ = cmd.MarkPersistentFlagRequired("value")
	return cmd
}

func NewOrgSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "organization",
		Short: "Trigger organization search",
		Run:   triggerOrgSearch,
	}

	cmd.PersistentFlags().String("name", "", "The name of the field to search for.")
	cmd.PersistentFlags().String("value", "", "Name of the field to search for")
	_ = cmd.MarkPersistentFlagRequired("name")
	//_ = cmd.MarkPersistentFlagRequired("value")
	return cmd
}
