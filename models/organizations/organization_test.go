package organizations

import (
	"ZendeskChallenge/internal"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type TestSuite struct {
	suite.Suite
	orgData OrgData
	raw     []byte
	org     Organization
}

func TestProcessSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) SetupSuite() {
	data, _ := os.ReadFile("testdata/organizations.json")
	// Fill the instance from the JSON file content
	var testOrgs Organization
	_ = json.Unmarshal(data, &testOrgs)
	suite.org = testOrgs
	suite.raw = data
	suite.orgData = OrgData{
		Raw:       data,
		Processed: testOrgs,
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

func (suite *TestSuite) TestOrg_Fetch() {
	suite.Run("Test Fetch works as expected", func() {
		suite.Equal(len(suite.org.Fetch()), 3)
		for i, _ := range suite.org.Fetch() {
			suite.Equal(suite.org.Fetch()[i], suite.org[i])
		}
	})
}

func (suite *TestSuite) TestOrg_ImplementsInterface() {
	suite.Run("Test OrgData and Organization implement correct interface", func() {
		suite.Implements((*internal.DataStore)(nil), suite.org)
		suite.Implements((*internal.DataProcessor)(nil), &suite.orgData)
	})
}

func (suite *TestSuite) TestOrgData_FetchFiltered() {
	suite.Run("Test FetchFiltered works as expected", func() {
		suite.Equal(suite.orgData.FetchFiltered(), Organization(nil))
	})
}

func (suite *TestSuite) TestOrgData_FetchRaw() {
	suite.Run("Test FetchRaw works as expected", func() {
		suite.Equal(suite.orgData.FetchRaw(), suite.raw)
	})
}

func (suite *TestSuite) TestOrgData_SetFiltered() {
	suite.Run("Test SetFiltered works as expected", func() {
		suite.Equal(suite.orgData.FetchFiltered(), Organization(nil))
		val, err := suite.orgData.SetFiltered(suite.org)
		suite.NotNil(val)
		suite.Nil(err)
		suite.Equal(suite.orgData.FetchFiltered(), suite.org)
	})
}

func (suite *TestSuite) TestOrgData_FetchProcessed() {
	suite.Run("Test FetchProcessed works as expected", func() {
		suite.Equal(len(suite.orgData.FetchProcessed()), 3)
		for i, _ := range suite.orgData.FetchProcessed() {
			suite.Equal(suite.orgData.FetchProcessed()[i], suite.org[i])
		}
	})
}
