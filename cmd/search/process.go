// Package search -
//
// This file is meant for processing of all search queries, and is where most of the heavy-lifting of storage
// and data processing is done
package search

import (
	"ZendeskChallenge/internal"
	"ZendeskChallenge/models/organizations"
	"ZendeskChallenge/models/tickets"
	"ZendeskChallenge/models/users"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"reflect"
	"strconv"
)

/*
*	Add related user entities for each user (tickets, organizations) to each user, in the resulting filtered output
 */
func addRelatedUserEntities(result users.User, orgDataRaw, ticketDataRaw []byte) {
	var allOrgs organizations.Organization
	var allTickets tickets.Ticket

	// Fill the instance from the JSON file content
	_, _ = json.Unmarshal(orgDataRaw, &allOrgs), json.Unmarshal(ticketDataRaw, &allTickets)
	orgData := organizations.OrgData{
		Raw:       orgDataRaw,
		Processed: allOrgs,
	}
	ticketsData := tickets.TicketData{
		Raw:       ticketDataRaw,
		Processed: allTickets,
	}
	for i, u := range result {
		obj, _ := parseByDataType(&orgData)
		parseString, _ := jp.ParseString("$..[?(@._id==" + strconv.Itoa(u.OrganizationId) + ")].name")
		org := parseString.Get(obj)
		if len(org) > 0 {
			u.OrganizationName = org[0].(string)
		}
		obj, _ = parseByDataType(&ticketsData)
		parseString, _ = jp.ParseString("$..[?(@.submitter_id==" + strconv.Itoa(u.Id) + ")].description")
		relatedTickets := parseString.Get(obj)
		if len(relatedTickets) > 0 {
			var allTickets []string
			for _, ticket := range relatedTickets {
				allTickets = append(allTickets, ticket.(string))
			}
			u.Tickets = allTickets
		}
		result[i] = u
	}
}

/*
*	Add related ticket entities for each ticket (user, organizations) to each ticket, in the resulting filtered output
 */
func addRelatedTicketEntities(results tickets.Ticket, orgDataRaw, userDataRaw []byte) {
	var allOrgs organizations.Organization
	var allUsers users.User
	_, _ = json.Unmarshal(orgDataRaw, &allOrgs), json.Unmarshal(userDataRaw, &allUsers)
	orgData := organizations.OrgData{
		Raw:       orgDataRaw,
		Processed: allOrgs,
	}
	userData := users.UserData{
		Raw:       userDataRaw,
		Processed: allUsers,
	}
	for i, ticket := range results {
		obj, _ := parseByDataType(&orgData)
		parseString, _ := jp.ParseString("$..[?(@._id==" + strconv.Itoa(ticket.OrganizationId) + ")].name")
		org := parseString.Get(obj)
		if len(org) > 0 {
			ticket.OrganizationName = org[0].(string)
		}
		obj, _ = parseByDataType(&userData)
		parseString, _ = jp.ParseString("$..[?(@._id==" + strconv.Itoa(ticket.SubmitterId) + ")].name")
		submitter := parseString.Get(obj)
		if len(submitter) > 0 {
			ticket.SubmitterName = submitter[0].(string)
		}
		parseString, _ = jp.ParseString("$..[?(@._id==" + strconv.Itoa(ticket.AssigneeId) + ")].name")
		assignee := parseString.Get(obj)
		if len(assignee) > 0 {
			ticket.AssigneeName = assignee[0].(string)
		}
		results[i] = ticket
	}
}

/*
* Generic Search evaluator, used for all models of searching (user, ticket and organizations)
*
*  - Validates the input flags as part of the underlying flags structs complying to a standard interface,
*    making it easy to treat flags as a common type
*
*    @return error, DataProcessor: Error if any, and all consolidated search in DataProcessor object
 */
func evaluateSearch(flags interface{}, data internal.DataProcessor, mappings map[string]string) (internal.DataProcessor, error) {
	validate := validator.New()
	var name = reflect.ValueOf(flags).FieldByName("Name").String()
	var value = reflect.ValueOf(flags).FieldByName("Value").String()
	err := validate.Struct(flags)
	if err != nil {
		return nil, err
	}
	val := reflect.ValueOf(data.FetchProcessed()[0]) // Get property value from one object to find field type
	field := val.FieldByName(mappings[name])
	result, err := evaluateSearchResultByDataType(field.Kind(), value, name, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
*
*	Evaluate the result of each search depending on type of field (underlying data type) being queried
 */
func evaluateSearchResultByDataType(fieldKind reflect.Kind, value, name string, data internal.DataProcessor) (internal.DataProcessor, error) {
	switch fieldKind {
	case reflect.Int:
		//log.Printf("VAL %v", value)
		_, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Please specify int type of --value associated with --name of %v\n", name))
		} else {
			obj, err := parseByDataType(data)
			parseString, err := jp.ParseString("$..[?(@." + name + "==" + value + ")]")
			if err != nil {
				return nil, err
			}
			return data.SetFiltered(parseString.Get(obj))
		}
	case reflect.String:
		obj, err := parseByDataType(data)
		parseString, err := jp.ParseString("$..[?(@." + name + "==\"" + value + "\")]")
		if err != nil {
			return nil, err
		}
		return data.SetFiltered(parseString.Get(obj))
	case reflect.Bool:
		parsedBool, err := strconv.ParseBool(value)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Please specify bool type of --value associated with --name of %v\n", name))
		} else {
			obj, err := parseByDataType(data)
			parseString, err := jp.ParseString("$..[?(@." + name + "==" + strconv.FormatBool(parsedBool) + ")]")
			if err != nil {
				return nil, err
			}
			return data.SetFiltered(parseString.Get(obj))
		}
	case reflect.Slice:
		obj, err := parseByDataType(data)
		parseString, err := jp.ParseString("$[?('" + value + "' in @." + name + ")]")
		if err != nil {
			return nil, err
		}
		return data.SetFiltered(parseString.Get(obj))
	default:
		return nil, errors.New("invalid data type not supported")
	}
}

/*
*
*	Parse the raw []bytes based on underlying model being queried, to extract the raw bytes from the common interface
*	they implement.
*
*	This allows method to be used interchangeably, making code simpler, and single responsibility.
 */
func parseByDataType(data interface{}) (any, error) {
	switch reflect.TypeOf(data).String() {
	case "*users.UserData":
		return oj.Parse(data.(*users.UserData).Raw)
	case "*tickets.TicketData":
		return oj.Parse(data.(*tickets.TicketData).Raw)
	case "*organizations.OrgData":
		return oj.Parse(data.(*organizations.OrgData).Raw)
	default:
		return nil, errors.New(fmt.Sprintf("Invalid data type not supported: %v", reflect.TypeOf(data)))
	}
}
