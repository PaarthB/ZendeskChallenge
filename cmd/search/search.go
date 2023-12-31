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

// Flags - Interface implemented by all models to allow easy retrieval of search flags, and use common data type
type Flags interface {
	FetchName() string
	FetchValue() string
}

/*
*		Get file data of specific file being queried.
*		Method behaves differently in test environment to allow reading test files
*
*	    @return ([]byte, error): Data content of file and error if reading caused issue (such as fs.PathError)
 */
func getFileData(fileName string) ([]byte, error) {
	isTest, err := strconv.ParseBool(os.Getenv("TEST_ENV"))
	var prefixPath string
	if err == nil && isTest {
		prefixPath = "testdata/"
	}
	return os.ReadFile(prefixPath + fileName)
}

/*
*		Trigger user search. Extracts flag values and delegates processing of query and output evaluation to `process.go` methods
*
*	    @return (error): If any error occurs during validation of flags, or evaluation of search
*		Displays results if no errors
 */
func triggerUserSearch(cmd *cobra.Command, args []string) error {
	var data, _ = getFileData(UsersFile)
	var orgData, _ = getFileData(OrganizationsFile)
	var ticketData, _ = getFileData(TicketsFile)
	// Declaration of the instance of the struct that we want to fill
	var allUsers users.User
	// Fill the instance from the JSON file content
	err := json.Unmarshal(data, &allUsers)
	if err != nil || len(allUsers) == 0 {
		cmd.PrintErrf("error occurred during parsing %v: %v, length of data: %v",
			UsersFile, err, len(allUsers))
		return err
	}
	userData := users.UserData{
		Raw:       data,
		Processed: allUsers,
	}
	value, _ := cmd.Flags().GetString("value")
	name, _ := cmd.Flags().GetString("name") // This is already validated by Cobra framework before reaching here
	flags := users.UserSearchFlags{
		Name:  name,
		Value: value,
	}
	result, err := evaluateSearch(flags, &userData, users.KeyMappings)
	if err != nil {
		cmd.PrintErr(err)
		log.Errorf(err.Error())
		return err
	}
	filteredUsers := result.FetchFiltered().(users.User)
	addRelatedUserEntities(filteredUsers, orgData, ticketData)
	internal.DisplayResults(cmd, filteredUsers, users.KeyMappings)
	log.Infof("All results displayed")
	return nil
}

/*
*		Trigger ticket search. Extracts flag values and delegates processing of query and output evaluation to `process.go` methods
*
*	    @return (error): If any error occurs during validation of flags, or evaluation of search
*		Displays results if no errors
 */
func triggerTicketSearch(cmd *cobra.Command, args []string) error {
	var data, _ = getFileData(TicketsFile)
	var orgData, _ = getFileData(OrganizationsFile)
	var userData, _ = getFileData(UsersFile)
	// Declaration of the instance of the struct that we want to fill
	var allTickets tickets.Ticket
	// Fill the instance from the JSON file content
	err := json.Unmarshal(data, &allTickets)
	if err != nil || len(allTickets) == 0 {
		cmd.PrintErrf("error occurred during parsing %v: %v, length of data: %v",
			TicketsFile, err, len(allTickets))
		return err
	}
	ticketData := tickets.TicketData{
		Raw:       data,
		Processed: allTickets,
	}
	value, _ := cmd.Flags().GetString("value")
	name, _ := cmd.Flags().GetString("name") // This is already validated by Cobra framework before reaching here
	flags := tickets.TicketSearchFlags{
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

/*
*		Trigger organization search. Extracts flag values and delegates processing of query and output evaluation to `process.go` methods
*
*	    @return (error): If any error occurs during validation of flags, or evaluation of search
*		Displays results if no errors
 */
func triggerOrgSearch(cmd *cobra.Command, args []string) error {
	var data, _ = getFileData(OrganizationsFile)
	// Declaration of the instance of the struct that we want to fill
	allOrgs := organizations.Organization{}
	// Fill the instance from the JSON file content
	err := json.Unmarshal(data, &allOrgs)
	if err != nil || len(allOrgs) == 0 {
		cmd.PrintErrf("error occurred during parsing %v: %v, length of data: %v",
			OrganizationsFile, err, len(allOrgs))
		return err
	}
	orgData := organizations.OrgData{
		Raw:       data,
		Processed: allOrgs,
	}
	value, _ := cmd.Flags().GetString("value")
	name, _ := cmd.Flags().GetString("name") // This is already validated by Cobra framework before reaching here
	flags := organizations.OrganizationSearchFlags{
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
