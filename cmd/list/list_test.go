package list

import (
	"ZendeskChallenge/models/organizations"
	"ZendeskChallenge/models/tickets"
	"ZendeskChallenge/models/users"
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_ExecuteListCommand(t *testing.T) {
	t.Run("Execute CLI list command and assert output", func(t *testing.T) {
		buffer := new(bytes.Buffer) // Redirecting output to custom buffer for testability.
		cmd := NewListCmd()
		cmd.SetOut(buffer)
		cmd.SetErr(buffer)
		err := cmd.Execute()
		assert.Nil(t, err)

		expected := "Zendesk CLI list command"
		assert.True(t, strings.HasPrefix(buffer.String(), expected), "actual is not expected")
		for key, _ := range users.KeyMappings {
			assert.True(t, strings.Contains(buffer.String(), key), "user field names are contained in output, for field "+key+"'")
		}
		for key, _ := range organizations.KeyMappings {
			assert.True(t, strings.Contains(buffer.String(), key), "organization field names are contained in output, for field "+key+"'")
		}
		for key, _ := range tickets.KeyMappings {
			assert.True(t, strings.Contains(buffer.String(), key), "ticket field names are contained in output, for field '"+key+"'")
		}
	})
}
