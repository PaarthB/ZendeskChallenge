package search

import (
	"ZendeskChallenge/internal"
	"ZendeskChallenge/models/organizations"
	"ZendeskChallenge/models/tickets"
	"ZendeskChallenge/models/users"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/ohler55/ojg/oj"
	"github.com/stretchr/testify/suite"
	_ "github.com/stretchr/testify/suite"
	"os"
	"reflect"
	"strconv"
	"testing"
)

// Test processing of user invoked search queries, and evaluating the output

type TestSuite struct {
	suite.Suite
	orgRaw     []byte
	userRaw    []byte
	ticketRaw  []byte
	orgData    organizations.OrgData
	userData   users.UserData
	ticketData tickets.TicketData
}

func TestProcessSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func GetSampleOrgData() organizations.OrgData {
	data, _ := os.ReadFile("testdata/organizations.json")
	var allOrgs organizations.Organization
	// Fill the instance from the JSON file content
	_ = json.Unmarshal(data, &allOrgs)

	return organizations.OrgData{
		Raw:       data,
		Processed: allOrgs,
	}
}

func GetSampleTicketData() tickets.TicketData {
	data, _ := os.ReadFile("testdata/tickets.json")
	var allTickets tickets.Ticket
	// Fill the instance from the JSON file content
	_ = json.Unmarshal(data, &allTickets)

	return tickets.TicketData{
		Raw:       data,
		Processed: allTickets,
	}
}

func GetSampleUserData() users.UserData {
	data, _ := os.ReadFile("testdata/users.json")
	var allUsers users.User
	// Fill the instance from the JSON file content
	_ = json.Unmarshal(data, &allUsers)

	return users.UserData{
		Raw:       data,
		Processed: allUsers,
	}
}

// this function executes after all tests executed
func (suite *TestSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}

func (suite *TestSuite) TestAddRelatedUserEntities() {
	testsSuccess := []struct {
		title     string
		users     users.User
		rawTicket []byte
		rawOrg    []byte
		expected  map[int]map[string]interface{}
	}{
		{
			title: "Add related ticket entities - entities non-existent",
			// !! IMPORTANT this test case needs to run before to avoid using the newly set values from next test case
			users:     suite.userData.Processed,
			rawTicket: nil,
			rawOrg:    nil,
			expected: map[int]map[string]interface{}{
				707070707: {
					"OrganizationName": "",
					"Tickets":          []string(nil),
				},
				70: {
					"OrganizationName": "",
					"Tickets":          []string(nil),
				},
				22: {
					"OrganizationName": "",
					"Tickets":          []string(nil),
				},
				74: {
					"OrganizationName": "",
					"Tickets":          []string(nil),
				},
				43: {
					"OrganizationName": "",
					"Tickets":          []string(nil),
				},
			},
		},
		{
			title:     "Add related user entities - success",
			users:     suite.userData.Processed,
			rawTicket: suite.ticketData.Raw,
			rawOrg:    suite.orgData.Raw,
			expected: map[int]map[string]interface{}{
				707070707: {
					"OrganizationName": "Isotronic",
					"Tickets":          []string(nil),
				},
				70: {
					"OrganizationName": "Isotronic",
					"Tickets":          []string{"test description 2"},
				},
				22: {
					"OrganizationName": "Terrasys",
					"Tickets":          []string{"test description 1"},
				},
				74: {
					"OrganizationName": "Hotcâkes",
					"Tickets":          []string(nil),
				},
				43: {
					"OrganizationName": "",
					"Tickets":          []string(nil),
				},
			},
		},
	}

	for _, tt := range testsSuccess {
		suite.Run(tt.title, func() {
			for i := range tt.users { // Related entities all null before search
				suite.Empty(tt.users[i].Tickets)
				suite.Empty(tt.users[i].OrganizationName)
			}
			addRelatedUserEntities(tt.users, tt.rawOrg, tt.rawTicket) // Add entities to all tickets
			for _, user := range tt.users {
				value, _ := tt.expected[user.Id]["Tickets"]
				suite.Equal(user.Tickets, value)
				value, _ = tt.expected[user.Id]["OrganizationName"]
				suite.Equal(user.OrganizationName, value)
			}
		})
	}
}

func (suite *TestSuite) TestAddRelatedTicketEntities() {
	testsSuccess := []struct {
		title    string
		tickets  tickets.Ticket
		rawUser  []byte
		rawOrg   []byte
		expected map[string]map[string]string
	}{
		{
			title: "Add related ticket entities - nothing to change",
			// !! IMPORTANT this test case needs to run before to avoid using the newly set values from next test case
			tickets: suite.ticketData.Processed,
			rawUser: nil,
			rawOrg:  nil,
			expected: map[string]map[string]string{
				"20615fe1-765b-4ff5-b4f6-ea42dcc8cac3": {
					"SubmitterName":    "",
					"AssigneeName":     "",
					"OrganizationName": "",
				},
				"3ff0599a-fe0f-4f8f-ac31-e2636843bcea": {
					"SubmitterName":    "",
					"AssigneeName":     "",
					"OrganizationName": "",
				},
				"7c67b6ed-6776-4065-bd4a-f2d9d12c33b7": {
					"AssigneeName":     "",
					"OrganizationName": "",
					"SubmitterName":    "",
				},
			},
		},
		{
			title:   "Add related ticket entities - success",
			tickets: suite.ticketData.Processed,
			rawUser: suite.userData.Raw,
			rawOrg:  suite.orgData.Raw,
			expected: map[string]map[string]string{
				"20615fe1-765b-4ff5-b4f6-ea42dcc8cac3": {
					"SubmitterName":    "Moran Daniels",
					"AssigneeName":     "Catalina Simpson",
					"OrganizationName": "Geekfarm",
				},
				"3ff0599a-fe0f-4f8f-ac31-e2636843bcea": {
					"SubmitterName":    "Valentine Ashley",
					"AssigneeName":     "",
					"OrganizationName": "Geekfarm",
				},
				"7c67b6ed-6776-4065-bd4a-f2d9d12c33b7": {
					"AssigneeName":     "Melissa Bishop",
					"OrganizationName": "Terrasys",
					"SubmitterName":    "",
				},
			},
		},
	}

	for _, tt := range testsSuccess {
		suite.Run(tt.title, func() {
			//suite.Nil(tt.data.FetchFiltered()) // No filtered results prior to search
			for i := range tt.tickets { // Related entities all null before search
				suite.Empty(tt.tickets[i].SubmitterName)
				suite.Empty(tt.tickets[i].AssigneeName)
				suite.Empty(tt.tickets[i].OrganizationName)
			}
			addRelatedTicketEntities(tt.tickets, tt.rawOrg, tt.rawUser) // Add entities to all tickets
			for _, ticket := range tt.tickets {
				value, _ := tt.expected[ticket.Id]["SubmitterName"]
				suite.Equal(ticket.SubmitterName, value)
				value, _ = tt.expected[ticket.Id]["AssigneeName"]
				suite.Equal(ticket.AssigneeName, value)
				value, _ = tt.expected[ticket.Id]["OrganizationName"]
				suite.Equal(ticket.OrganizationName, value)
			}
		})
	}
}

func (suite *TestSuite) TestEvaluateSearch_Error() {
	testsError := []struct {
		title        string
		data         internal.DataProcessor
		mappings     map[string]string
		flags        interface{}
		errorMessage string
	}{
		{
			title:    "evaluate org search - pass invalid field to search for",
			data:     &suite.orgData,
			mappings: organizations.KeyMappings,
			flags:    organizations.OrganizationFlags{Name: "invalid_field", Value: "121"},
		},
		{
			title:    "evaluate user search - pass invalid field to search for",
			mappings: users.KeyMappings,
			data:     &suite.userData,
			flags:    users.UserFlags{Name: "random field", Value: "74"},
		},
		{
			title:    "evaluate ticket search - pass invalid field to search for",
			mappings: tickets.KeyMappings,
			data:     &suite.ticketData,
			flags:    tickets.TicketFlags{Name: "useless field", Value: "7c67b6ed-6776-4065-bd4a-f2d9d12c33b7"},
		},
	}

	for _, tt := range testsError {
		suite.Run(tt.title, func() {
			suite.Nil(tt.data.FetchFiltered()) // No filtered results prior to search
			val, err := evaluateSearch(tt.flags, tt.data, tt.mappings)
			suite.NotNil(err)
			suite.Nil(val)
			suite.IsType((validator.ValidationErrors)(nil), err) // Assert correct error type thrown
			suite.Nil(tt.data.FetchFiltered())                   // Filtered results remain nil as error occurred
		})
	}

}
func (suite *TestSuite) TestEvaluateSearch_Success() {
	testsSuccess := []struct {
		title    string
		data     internal.DataProcessor
		mappings map[string]string
		count    int
		flags    interface{}
		dataType string
	}{
		{
			title:    "evaluate org search",
			data:     &suite.orgData,
			mappings: organizations.KeyMappings,
			count:    1,
			dataType: "*organizations.OrgData",
			flags:    organizations.OrganizationFlags{Name: "_id", Value: "121"},
		},
		{
			title:    "evaluate user search",
			mappings: users.KeyMappings,
			data:     &suite.userData,
			count:    1,
			dataType: "*users.UserData",
			flags:    users.UserFlags{Name: "_id", Value: "74"},
		},
		{
			title:    "evaluate ticket search",
			mappings: tickets.KeyMappings,
			data:     &suite.ticketData,
			count:    1,
			dataType: "*tickets.TicketData",
			flags:    tickets.TicketFlags{Name: "_id", Value: "7c67b6ed-6776-4065-bd4a-f2d9d12c33b7"},
		},
	}

	for _, tt := range testsSuccess {
		suite.Run(tt.title, func() {
			suite.Nil(tt.data.FetchFiltered()) // No filtered results prior to search
			val, err := evaluateSearch(tt.flags, tt.data, tt.mappings)
			suite.Nil(err)
			suite.NotNil(val)
			suite.Equal(reflect.TypeOf(val).String(), tt.dataType)
			suite.NotNil(tt.data.FetchFiltered()) // Filtered is set after successful search
			suite.Equal(len(val.FetchFiltered().Fetch()), tt.count)
		})
	}
}

func (suite *TestSuite) TestEvaluateSearchResultByDataType_Error() {
	testsError := []struct {
		fieldKind    reflect.Kind
		title        string
		name         string
		value        string
		data         internal.DataProcessor
		count        int
		errorMessage string
	}{
		{
			title:        "Search by integer field but specifying invalid integer",
			fieldKind:    reflect.Int,
			value:        "invalid integer",
			name:         "_id",
			data:         &suite.orgData,
			errorMessage: "Please specify int type of --value associated with --name of _id\n",
		},
		{
			title:        "Search by bool field but specifying invalid bool",
			fieldKind:    reflect.Bool,
			value:        "invalid bool",
			name:         "suspended",
			data:         &suite.userData,
			errorMessage: "Please specify bool type of --value associated with --name of suspended\n",
		},
		{
			title:        "Invalid data type not supported by CLI",
			fieldKind:    reflect.Pointer,
			value:        "invalid bool",
			name:         "suspended",
			data:         &suite.userData,
			errorMessage: "invalid data type not supported",
		},
	}

	for _, tt := range testsError {
		suite.Run(tt.title, func() {
			suite.Nil(tt.data.FetchFiltered()) // No filtered results prior to search
			val, err := evaluateSearchResultByDataType(tt.fieldKind, tt.value, tt.name, tt.data)
			suite.NotNil(err) // Error has occurred
			suite.Equal(tt.errorMessage, err.Error())
			suite.Nil(val)                     // Error causes nil value to be returned
			suite.Nil(tt.data.FetchFiltered()) // Results remain nil as error occurred
		})
	}
}
func (suite *TestSuite) TestEvaluateSearchResultByDataType_Success() {
	testsSuccess := []struct {
		fieldKind reflect.Kind
		title     string
		name      string
		value     string
		data      internal.DataProcessor
		count     int
	}{
		{
			title:     "Search by integer field",
			fieldKind: reflect.Int,
			value:     strconv.Itoa(121),
			name:      "_id",
			data:      &suite.orgData,
			count:     1,
		},
		{
			title:     "Search by string field",
			fieldKind: reflect.String,
			value:     "rosannasimpson@flotonic.com",
			name:      "email",
			data:      &suite.userData,
			count:     1,
		},
		{
			title:     "Search by array-type field",
			fieldKind: reflect.Slice,
			value:     "Massachusetts",
			name:      "tags",
			data:      &suite.ticketData,
			count:     1,
		},
		{
			title:     "Search by bool field",
			fieldKind: reflect.Bool,
			value:     strconv.FormatBool(true),
			name:      "verified",
			data:      &suite.userData,
			count:     1,
		},
		{
			title:     "Multiple result count - when searching for users who are suspended",
			fieldKind: reflect.Bool,
			value:     strconv.FormatBool(true),
			name:      "suspended",
			data:      &suite.userData,
			count:     3,
		},
	}

	for _, tt := range testsSuccess {
		suite.Run(tt.title, func() {
			suite.Nil(tt.data.FetchFiltered()) // No filtered results prior to search
			val, err := evaluateSearchResultByDataType(tt.fieldKind, tt.value, tt.name, tt.data)
			suite.Nil(err)
			suite.NotNil(val)
			suite.NotNil(tt.data.FetchFiltered()) // Filtered is set after successful search
			suite.Equal(len(val.FetchFiltered().Fetch()), tt.count)
			for i := 0; i < len(val.FetchFiltered().Fetch()); i++ {
				// The filtered value returned is same struct as the filtered value in original struct
				suite.Equal(tt.data.FetchFiltered().Fetch()[i], val.FetchFiltered().Fetch()[i])
			}
			_, _ = tt.data.SetFiltered(nil) // Resetting filtered for next run of test cases
		})
	}
}

func (suite *TestSuite) TestParseDataType() {
	testsSuccess := []struct {
		name  string
		data  interface{}
		value []byte
	}{
		{
			name:  "Should return the right raw data if OrgData struct passed in",
			value: []byte(`{"id":"someID","data":"str1str2"}`),
			data:  &organizations.OrgData{Raw: []byte(`{"id":"someID","data":"str1str2"}`)},
		},
		{
			name:  "Should return the right raw data if UserData struct passed in",
			value: []byte(`{"id":"someID","data":"str1str2"}`),
			data:  &users.UserData{Raw: []byte(`{"id":"someID","data":"str1str2"}`)},
		},
		{
			name:  "Should return the right raw data if TicketData struct passed in",
			value: []byte(`{"id":"someID","data":"str1str2"}`),
			data:  &tickets.TicketData{Raw: []byte(`{"id":"someID","data":"str1str2"}`)},
		},
	}
	testsFail := []struct {
		name  string
		data  interface{}
		value []byte
	}{
		{
			name:  "Parsing should throw error if random struct type is passed in",
			data:  "random non-struct data",
			value: []byte(`Non JSON value`),
		},
	}

	for _, tt := range testsSuccess {
		suite.Run(tt.name, func() {
			val1, err1 := parseByDataType(tt.data)
			val2, err2 := oj.Parse(tt.value)
			suite.Nil(err1)
			suite.Nil(err2)
			suite.Equal(val1, val2)
		})

		for _, tt := range testsFail {
			suite.Run(tt.name, func() {
				val1, err1 := parseByDataType(tt.data)
				val2, err2 := oj.Parse(tt.value)
				suite.NotNil(err1)
				suite.NotNil(err2)
				suite.Nil(val1)
				suite.Nil(val2)
				suite.Equal(err1.Error(), "Invalid data type not supported: string")
				suite.IsType(err2, &oj.ParseError{})
			})
		}
	}
}
