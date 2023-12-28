package users

import (
	"ZendeskChallenge/internal"
	"encoding/json"
	"fmt"
	"os"

	//"go/models"
	//"strings"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
	userData UserData
	raw      []byte
	user     User
}

func TestProcessSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) SetupSuite() {
	data, _ := os.ReadFile("testdata/users.json")
	// Fill the instance from the JSON file content
	var testUsers User
	_ = json.Unmarshal(data, &testUsers)
	suite.user = testUsers
	suite.raw = data
	suite.userData = UserData{
		Raw:       data,
		Processed: testUsers,
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

func (suite *TestSuite) TestUser_Fetch() {
	suite.Run("Test Fetch works as expected", func() {
		suite.Equal(len(suite.user.Fetch()), 4)
		for i := range suite.user.Fetch() {
			suite.Equal(suite.user.Fetch()[i], suite.user[i])
		}
	})
}

func (suite *TestSuite) TestUser_ImplementsInterface() {
	suite.Run("Test UserData and User implement correct interface", func() {
		suite.Implements((*internal.DataStore)(nil), suite.user)
		suite.Implements((*internal.DataProcessor)(nil), &suite.userData)
	})
}

func (suite *TestSuite) TestUserData_FetchFiltered() {
	suite.Run("Test FetchFiltered works as expected", func() {
		suite.Equal(suite.userData.FetchFiltered(), User(nil))
	})
}

func (suite *TestSuite) TestUserData_FetchRaw() {
	suite.Run("Test FetchRaw works as expected", func() {
		suite.Equal(suite.userData.FetchRaw(), suite.raw)
	})
}

func (suite *TestSuite) TestUserData_SetFiltered() {
	suite.Run("Test SetFiltered works as expected", func() {
		suite.Equal(suite.userData.FetchFiltered(), User(nil))
		val, err := suite.userData.SetFiltered(suite.user)
		suite.NotNil(val)
		suite.Nil(err)
		suite.Equal(suite.userData.FetchFiltered(), suite.user)
	})
}

func (suite *TestSuite) TestUserData_FetchProcessed() {
	suite.Run("Test FetchProcessed works as expected", func() {
		suite.Equal(len(suite.userData.FetchProcessed()), 4)
		for i := range suite.userData.FetchProcessed() {
			suite.Equal(suite.userData.FetchProcessed()[i], suite.user[i])
		}
	})
}
