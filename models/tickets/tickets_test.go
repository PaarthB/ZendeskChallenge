package tickets

import (
	"ZendeskChallenge/internal"
	"encoding/json"
	"fmt"

	//"go/models"
	//"strings"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type TestSuite struct {
	suite.Suite
	ticketData TicketData
	raw        []byte
	ticket     Ticket
}

func TestProcessSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) SetupSuite() {
	data, _ := os.ReadFile("testdata/tickets.json")
	// Fill the instance from the JSON file content
	var testTickets Ticket
	_ = json.Unmarshal(data, &testTickets)
	suite.ticket = testTickets
	suite.raw = data
	suite.ticketData = TicketData{
		Raw:       data,
		Processed: testTickets,
	}
}

// this function executes after all tests executed
func (suite *TestSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}

func (suite *TestSuite) SetupTest() {
	// reset StartingNumber to one
	fmt.Println("-- From SetupTest")
}

// this function executes after each test case
func (suite *TestSuite) TearDownTest() {
	fmt.Println("-- From TearDownTest")
}

func (suite *TestSuite) TestTicket_Fetch() {
	suite.Run("Test Fetch works as expected", func() {
		suite.Equal(len(suite.ticket.Fetch()), 3)
		for i, _ := range suite.ticket.Fetch() {
			suite.Equal(suite.ticket.Fetch()[i], suite.ticket[i])
		}
	})
}

func (suite *TestSuite) TestTicket_ImplementsInterface() {
	suite.Run("Test TicketData and Ticket implement correct interface", func() {
		suite.Implements((*internal.DataStore)(nil), suite.ticket)
		suite.Implements((*internal.DataProcessor)(nil), &suite.ticketData)
	})
}

func (suite *TestSuite) TestTicketData_FetchFiltered() {
	suite.Run("Test FetchFiltered works as expected", func() {
		suite.Equal(suite.ticketData.FetchFiltered(), Ticket(nil))
	})
}

func (suite *TestSuite) TestTicketData_FetchRaw() {
	suite.Run("Test FetchRaw works as expected", func() {
		suite.Equal(suite.ticketData.FetchRaw(), suite.raw)
	})
}

func (suite *TestSuite) TestTicketData_SetFiltered() {
	suite.Run("Test SetFiltered works as expected", func() {
		suite.Equal(suite.ticketData.FetchFiltered(), Ticket(nil))
		val, err := suite.ticketData.SetFiltered(suite.ticket)
		suite.NotNil(val)
		suite.Nil(err)
		suite.Equal(suite.ticketData.FetchFiltered(), suite.ticket)
	})
}

func (suite *TestSuite) TestTicketData_FetchProcessed() {
	suite.Run("Test FetchProcessed works as expected", func() {
		suite.Equal(len(suite.ticketData.FetchProcessed()), 3)
		for i, _ := range suite.ticketData.FetchProcessed() {
			suite.Equal(suite.ticketData.FetchProcessed()[i], suite.ticket[i])
		}
	})
}
