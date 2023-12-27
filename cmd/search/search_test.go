package search

import (
	"ZendeskChallenge/models/organizations"
	"ZendeskChallenge/models/tickets"
	"ZendeskChallenge/models/users"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func (suite *TestSuite) SetupSuite() {
	suite.orgRaw, _ = os.ReadFile("testdata/organizations.json")
	suite.ticketRaw, _ = os.ReadFile("testdata/tickets.json")
	suite.userRaw, _ = os.ReadFile("testdata/users.json")
	suite.orgData = GetSampleOrgData()
	suite.ticketData = GetSampleTicketData()
	suite.userData = GetSampleUserData()
}

func (suite *TestSuite) SetupTest() {
	// reset StartingNumber to one
	fmt.Println("-- From SetupTest")
	_ = os.Setenv("TEST_ENV", "true") // Set for using different file data source for tests
}

// this function executes after each test case
func (suite *TestSuite) TearDownTest() {
	fmt.Println("-- From TearDownTest")
}

func (suite *TestSuite) Test_ExecuteSearchCommand_Help() {
	suite.Run("Execute search command with help flag invoked and assert output", func() {
		buffer := new(bytes.Buffer) // Redirecting output to custom buffer for testability.
		cmd := NewSearchCmd()

		subCommands := cmd.Commands()
		for i := range subCommands {
			subCommands[i].SetOut(buffer)
			subCommands[i].SetErr(buffer)
		}

		cmd.SetOut(buffer)
		cmd.SetErr(buffer)
		cmd.SetArgs([]string{"--help"})

		err := cmd.Execute()
		suite.Nil(err)

		expected := "Zendesk Search"
		suite.True(strings.HasPrefix(buffer.String(), expected), "Message output starts as expected")

		// Contains direction to help messages
		suite.True(strings.Contains(buffer.String(), "trigger organization search"), "Shows organization search usage")
		suite.True(strings.Contains(buffer.String(), "trigger user search"), "Shows user search usage")
		suite.True(strings.Contains(buffer.String(), "trigger ticket search"), "Shows ticket search usage")
	})
}

func (suite *TestSuite) Test_ExecuteSearchCommand_UserInvalid() {
	suite.Run("Execute invalid ticket search", func() {
		buffer := new(bytes.Buffer) // Redirecting output to custom buffer for testability.
		cmd := NewUserSearchCmd()

		cmd.SetOut(buffer)
		cmd.SetErr(buffer)
		cmd.SetArgs([]string{"user"})

		err := cmd.Execute()
		expected := "required flag(s) \"name\" not set"
		suite.Equal(err.Error(), expected)
	})
}

func (suite *TestSuite) Test_ExecuteSearchCommand_UserValid() {
	suite.Run("Execute valid ticket search command and assert output", func() {
		buffer := new(bytes.Buffer) // Redirecting output to custom buffer for testability.
		cmd := NewUserSearchCmd()

		cmd.SetOut(buffer)
		cmd.SetErr(buffer)
		var name, value = "_id", "707070707" // Random ID being read from test file dynamically by use of environment variable
		cmd.SetArgs([]string{"user", "--name", name, "--value", value})
		buffer.Reset()
		err := cmd.Execute()
		suite.Nil(err)

		expected := "======== All results ========"
		suite.True(strings.HasPrefix(buffer.String(), expected), "Message output starts as expected")

		var user users.User
		_ = json.Unmarshal(suite.userRaw, &user)
		var userVal = reflect.ValueOf(user[0])
		for i, val := range user[0].Tags { // Test tags are set separately as they are list values
			suite.True(strings.Contains(buffer.String(), "tags_"+strconv.Itoa(i)+": "+val), "Tags are contained in output")
		}
		for key, val := range users.KeyMappings {
			var outputVal string
			switch userVal.FieldByName(val).Kind() {
			case reflect.Int:
				outputVal = strconv.Itoa(int(userVal.FieldByName(val).Int()))
			case reflect.String:
				outputVal = userVal.FieldByName(val).String()
			case reflect.Bool:
				outputVal = strconv.FormatBool(userVal.FieldByName(val).Bool())
			default:
				continue
			}
			// Contains correct data as expected
			suite.True(strings.Contains(buffer.String(), key+": "+outputVal), "Shows user search resulted in correct fields")
		}
	})
}

func (suite *TestSuite) Test_ExecuteSearchCommand_TicketInvalid() {
	suite.Run("Execute invalid ticket search", func() {
		buffer := new(bytes.Buffer) // Redirecting output to custom buffer for testability.
		cmd := NewTicketSearchCmd()

		cmd.SetOut(buffer)
		cmd.SetErr(buffer)
		cmd.SetArgs([]string{"ticket"})

		err := cmd.Execute()
		expected := "required flag(s) \"name\" not set"
		suite.Equal(err.Error(), expected)
	})
}

func (suite *TestSuite) Test_ExecuteSearchCommand_TicketValid() {
	suite.Run("Execute valid ticket search command and assert output ", func() {
		buffer := new(bytes.Buffer) // Redirecting output to custom buffer for testability.
		cmd := NewTicketSearchCmd()

		cmd.SetOut(buffer)
		cmd.SetErr(buffer)
		var name, value = "_id", "test_id" // Random ID being read from test file dynamically by use of environment variable
		cmd.SetArgs([]string{"ticket", "--name", name, "--value", value})
		buffer.Reset()
		err := cmd.Execute()
		suite.Nil(err)

		expected := "======== All results ========"
		suite.True(strings.HasPrefix(buffer.String(), expected), "Message output starts as expected")

		var ticket tickets.Ticket
		_ = json.Unmarshal(suite.ticketRaw, &ticket)
		var ticketVal = reflect.ValueOf(ticket[0])
		for i, val := range ticket[0].Tags { // Test tags are set separately as they are list values
			suite.True(strings.Contains(buffer.String(), "tags_"+strconv.Itoa(i)+": "+val), "Tags are contained in output")
		}
		for key, val := range tickets.KeyMappings {
			var outputVal string
			switch ticketVal.FieldByName(val).Kind() {
			case reflect.Int:
				outputVal = strconv.Itoa(int(ticketVal.FieldByName(val).Int()))
			case reflect.String:
				outputVal = ticketVal.FieldByName(val).String()
			case reflect.Bool:
				outputVal = strconv.FormatBool(ticketVal.FieldByName(val).Bool())
			default:
				continue
			}
			// Contains correct data as expected
			suite.True(strings.Contains(buffer.String(), key+": "+outputVal), "Shows ticket search resulted in correct fields")
		}
	})
}

func (suite *TestSuite) Test_ExecuteSearchCommand_OrgInvalid() {
	suite.Run("Execute invalid organization search", func() {
		buffer := new(bytes.Buffer) // Redirecting output to custom buffer for testability.
		cmd := NewOrgSearchCmd()

		cmd.SetOut(buffer)
		cmd.SetErr(buffer)
		cmd.SetArgs([]string{"organization"})

		err := cmd.Execute()
		expected := "required flag(s) \"name\" not set"
		suite.Equal(err.Error(), expected)
	})
}

func (suite *TestSuite) Test_ExecuteSearchCommand_OrgValid() {
	suite.Run("Execute valid organization search command and assert output", func() {
		buffer := new(bytes.Buffer) // Redirecting output to custom buffer for testability.
		cmd := NewOrgSearchCmd()

		cmd.SetOut(buffer)
		cmd.SetErr(buffer)
		var name, value = "_id", "919191919" // Random ID being read from test file dynamically by use of environment variable
		cmd.SetArgs([]string{"organization", "--name", name, "--value", value})
		buffer.Reset()
		err := cmd.Execute()
		suite.Nil(err)

		expected := "======== All results ========"
		suite.True(strings.HasPrefix(buffer.String(), expected), "Message output starts as expected")
		//suite.Equal(buffer.String(), "")
		var organization organizations.Organization
		_ = json.Unmarshal(suite.orgRaw, &organization)
		var orgVal = reflect.ValueOf(organization[0])
		for i, val := range organization[0].Tags { // Test tags are set separately as they are list values
			suite.True(strings.Contains(buffer.String(), "tags_"+strconv.Itoa(i)+": "+val), "Tags are contained in output")
		}
		for i, val := range organization[0].DomainNames { // Test domain names are set separately as they are list values
			suite.True(strings.Contains(buffer.String(), "domain_names_"+strconv.Itoa(i)+": "+val), "Tags are contained in output")
		}
		for key, val := range organizations.KeyMappings {
			var outputVal string
			switch orgVal.FieldByName(val).Kind() {
			case reflect.Int:
				outputVal = strconv.Itoa(int(orgVal.FieldByName(val).Int()))
			case reflect.String:
				outputVal = orgVal.FieldByName(val).String()
			case reflect.Bool:
				outputVal = strconv.FormatBool(orgVal.FieldByName(val).Bool())
			default:
				continue
			}
			// Contains correct data as expected
			suite.True(strings.Contains(buffer.String(), key+": "+outputVal), "Shows ticket search resulted in correct fields")
		}
	})
}
