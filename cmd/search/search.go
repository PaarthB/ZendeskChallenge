// Package search -
//
// This is the entry point of various search queries - user / ticket / organization, each of which invoke different
// entry-points. The source of data is read from real JSON files for users and from testdata files for tests
//

package search

import (
	"ZendeskChallenge/internal"
	"ZendeskChallenge/models/organizations"
	"ZendeskChallenge/models/tickets"
	"ZendeskChallenge/models/users"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

const (
	UsersFile         = "users.json"
	TicketsFile       = "tickets.json"
	OrganizationsFile = "organizations.json"
)

func getFileData(fileName string) ([]byte, error) {
	isTest, err := strconv.ParseBool(os.Getenv("TEST_ENV"))
	var prefixPath string
	if err == nil && isTest {
		prefixPath = "testdata/"
	}
	return os.ReadFile(prefixPath + fileName)
}

func triggerUserSearch(cmd *cobra.Command, args []string) error {
	var data, _ = getFileData(UsersFile)
	var orgData, _ = getFileData(OrganizationsFile)
	var ticketData, _ = getFileData(TicketsFile)
	// Declaration of the instance of the struct that we want to fill
	var allUsers users.User
	// Fill the instance from the JSON file content
	err := json.Unmarshal(data, &allUsers)
	if err != nil {
		log.Errorf("error %v", err)
		cmd.PrintErrf("error %v", err)
		return err
	}
	userData := users.UserData{
		Raw:       data,
		Processed: allUsers,
	}

	value, _ := cmd.Flags().GetString("value")
	name, _ := cmd.Flags().GetString("name") // This is already validated by Cobra framework before reaching here

	flags := users.UserFlags{
		Name:  name,
		Value: value,
	}
	result, err := evaluateSearch(flags, &userData, tickets.KeyMappings)
	if err != nil {
		cmd.PrintErr(err)
		log.Errorf(err.Error())
		return err
	}
	filteredUsers := result.FetchFiltered().(users.User)
	addRelatedUserEntities(filteredUsers, orgData, ticketData)
	internal.DisplayResults(cmd, filteredUsers, users.KeyMappings)
	log.Debugf("All results displayed")
	return nil
}

func triggerTicketSearch(cmd *cobra.Command, args []string) error {
	var data, _ = getFileData(TicketsFile)
	var orgData, _ = getFileData(OrganizationsFile)
	var userData, _ = getFileData(UsersFile)
	// Declaration of the instance of the struct that we want to fill
	var allTickets tickets.Ticket
	// Fill the instance from the JSON file content
	err := json.Unmarshal(data, &allTickets)
	if err != nil {
		log.Errorf("here!! error %v", err)
		cmd.PrintErrf("error %v", err)
		return err
	}
	ticketData := tickets.TicketData{
		Raw:       data,
		Processed: allTickets,
	}
	value, _ := cmd.Flags().GetString("value")
	name, _ := cmd.Flags().GetString("name") // This is already validated by Cobra framework before reaching here
	flags := tickets.TicketFlags{
		Name:  name,
		Value: value,
	}
	result, err := evaluateSearch(flags, &ticketData, tickets.KeyMappings)
	if err != nil {
		cmd.PrintErr(err)
		log.Errorf(err.Error())
		return err
	}
	filteredTickets := result.FetchFiltered().(tickets.Ticket)
	addRelatedTicketEntities(filteredTickets, orgData, userData)
	internal.DisplayResults(cmd, filteredTickets, tickets.KeyMappings)
	log.Info("All results displayed")
	return nil
}

func triggerOrgSearch(cmd *cobra.Command, args []string) error {
	var data, _ = getFileData(OrganizationsFile)
	// Declaration of the instance of the struct that we want to fill
	allOrgs := organizations.Organization{}
	// Fill the instance from the JSON file content
	err := json.Unmarshal(data, &allOrgs)
	if err != nil {
		log.Errorf("error %v", err)
		cmd.PrintErrf("error %v", err)
		return err
	}
	orgData := organizations.OrgData{
		Raw:       data,
		Processed: allOrgs,
	}
	value, _ := cmd.Flags().GetString("value")
	name, _ := cmd.Flags().GetString("name") // This is already validated by Cobra framework before reaching here
	flags := organizations.OrganizationFlags{
		Name:  name,
		Value: value,
	}
	result, err := evaluateSearch(flags, &orgData, organizations.KeyMappings)
	if err != nil {
		cmd.PrintErr(err)
		log.Errorf(err.Error())
		return err
	}
	internal.DisplayResults(cmd, result.FetchFiltered(), organizations.KeyMappings)
	log.Info("All results displayed")
	return nil
}
