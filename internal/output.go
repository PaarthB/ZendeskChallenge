// Package internal -
//
// Defines output related generic functions for the CLI application
package internal

import (
	"fmt"
	"github.com/spf13/cobra"
	"reflect"
	"strconv"
	"strings"
)

// DisplayResults - Displays final result to user, for the command that they ran
func DisplayResults(cmd *cobra.Command, results DataStore, keyMappings map[string]string) {
	cmd.Print("======== All results ========\n")
	if results == nil {
		cmd.Print("Nothing to display")
		return
	}
	var outputString = ""
	for _, entity := range results.Fetch() {
		cmd.Print("------------------------------------------------\n")
		r := reflect.ValueOf(entity)
		for key, val := range keyMappings {
			field := r.FieldByName(val)
			switch r.FieldByName(val).Kind() {
			case reflect.Slice:
				for i := 0; i < field.Len(); i++ {
					outputString += fmt.Sprintf("%v: %v\n", strings.Join([]string{key, strconv.Itoa(i)}, "_"), field.Index(i))
				}
				break
			default:
				outputString += fmt.Sprintf("%v: %v\n", key, r.FieldByName(val))
				break
			}
		}
		cmd.Print(outputString)
	}
}
