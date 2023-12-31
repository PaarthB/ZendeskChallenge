// Package list -
//
// Defines the single entry-point for displaying all possible searchable fields under the 'list' command
//

package list

import (
	"ZendeskChallenge/models/organizations"
	"ZendeskChallenge/models/tickets"
	"ZendeskChallenge/models/users"
	"github.com/spf13/cobra"
)

func fieldList(cmd *cobra.Command, args []string) error {
	cmd.Print("Searchable user fields with 'search user' command")
	cmd.Print("\n--------------------------------------------\n")
	for field := range users.KeyMappings {
		cmd.Printf("%v\n", field)
	}
	cmd.Print("\n\nSearchable organization fields with 'search organization' command")
	cmd.Print("\n--------------------------------------------\n")
	for field := range organizations.KeyMappings {
		cmd.Printf("%v\n", field)
	}
	cmd.Print("\n\nSearchable ticket fields with 'search ticket' command")
	cmd.Print("\n--------------------------------------------\n")
	for field := range tickets.KeyMappings {
		cmd.Printf("%v\n", field)
	}
	return nil
}
