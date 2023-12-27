package internal

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"strings"

	"testing"
)

func TestDisplayResults(t *testing.T) {
	t.Run("test display results works as expected - without passing actual data", func(t *testing.T) {
		buffer := new(bytes.Buffer) // Redirecting output to custom buffer for testability.
		cmd := cobra.Command{}      // Dummy command
		cmd.SetOut(buffer)          // Set buffer for testability of output statements
		cmd.SetErr(buffer)
		DisplayResults(&cmd, nil, map[string]string{
			"test key ": "test key mapping",
		}) // Not passing actual user/ticket/org structs to bypass cyclic imports
		assert.True(t, strings.HasPrefix(buffer.String(), "======== All results ========"))
	})
}
