package search

import (
	"ZendeskChallenge/internal"
	"ZendeskChallenge/types/organizations"
	"ZendeskChallenge/types/tickets"
	"ZendeskChallenge/types/users"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

const (
	UsersFile         = "users.json"
	TicketsFile       = "tickets.json"
	OrganizationsFile = "organizations.json"
)

func triggerUserSearch(cmd *cobra.Command, args []string) {
	value, err := cmd.Flags().GetString("value")
	var data, _ = os.ReadFile(UsersFile)
	var orgData, _ = os.ReadFile(OrganizationsFile)
	var ticketData, _ = os.ReadFile(TicketsFile)
	// Declaration of the instance of the struct that we want to fill
	var allUsers users.User
	// Fill the instance from the JSON file content
	err = json.Unmarshal(data, &allUsers)
	if err != nil {
		log.Errorf("error %v", err)
		cmd.PrintErrf("error %v", err)
		return
	}
	userData := users.UserData{
		Raw:       data,
		Processed: allUsers,
	}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		log.Errorf("error %v", err)
		cmd.PrintErrf("error %v", err)
		return
	}
	flags := users.UserFlags{
		Name:  name,
		Value: value,
	}
	result, err := evaluateSearch(flags, &userData, tickets.KeyMappings)
	if err != nil {
		cmd.PrintErr(err)
		log.Errorf(err.Error())
		return
	}
	filteredUsers := result.FetchFiltered().(users.User)
	addRelatedUserEntities(filteredUsers, orgData, ticketData)
	internal.DisplayResults(cmd, result.FetchFiltered(), users.KeyMappings)
	log.Debugf("All results displayed")
}

func triggerTicketSearch(cmd *cobra.Command, args []string) {
	value, err := cmd.Flags().GetString("value")
	var data, _ = os.ReadFile(TicketsFile)
	var orgData, _ = os.ReadFile(OrganizationsFile)
	var userData, _ = os.ReadFile(UsersFile)
	// Declaration of the instance of the struct that we want to fill
	var allTickets tickets.Ticket
	// Fill the instance from the JSON file content
	err = json.Unmarshal(data, &allTickets)
	if err != nil {
		log.Errorf("here!! error %v", err)
		cmd.PrintErrf("error %v", err)
		return
	}
	ticketData := tickets.TicketData{
		Raw:       data,
		Processed: allTickets,
	}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		log.Errorf("error %v", err)
		cmd.PrintErrf("error %v", err)
		return
	}
	flags := tickets.TicketFlags{
		Name:  name,
		Value: value,
	}
	result, err := evaluateSearch(flags, &ticketData, tickets.KeyMappings)
	if err != nil {
		cmd.PrintErr(err)
		log.Errorf(err.Error())
		return
	}
	filteredTickets := result.FetchFiltered().(tickets.Ticket)
	addRelatedTicketEntities(filteredTickets, orgData, userData)
	internal.DisplayResults(cmd, result.FetchFiltered(), tickets.KeyMappings)
	log.Debugf("All results displayed")
}

func triggerOrgSearch(cmd *cobra.Command, args []string) {
	value, err := cmd.Flags().GetString("value")
	var data, _ = os.ReadFile(OrganizationsFile)
	// Declaration of the instance of the struct that we want to fill
	allOrgs := organizations.Organization{}
	// Fill the instance from the JSON file content
	err = json.Unmarshal(data, &allOrgs)
	if err != nil {
		log.Errorf("error %v", err)
		cmd.PrintErrf("error %v", err)
		return
	}
	orgData := organizations.OrgData{
		Raw:       data,
		Processed: allOrgs,
	}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		log.Errorf("error %v", err)
		cmd.PrintErrf("error %v", err)
		return
	}
	flags := organizations.OrganizationFlags{
		Name:  name,
		Value: value,
	}
	result, err := evaluateSearch(flags, &orgData, organizations.KeyMappings)
	if err != nil {
		cmd.PrintErr(err)
		log.Errorf(err.Error())
		return
	}
	internal.DisplayResults(cmd, result.FetchFiltered(), organizations.KeyMappings)
	log.Debugf("All results displayed")
}
