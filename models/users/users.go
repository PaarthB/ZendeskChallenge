// Package users -
//
// Defines the users model and its key fields. Implements DataStore and DataProcessor interfaces
//

package users

import (
	"ZendeskChallenge/internal"
	"encoding/json"
)

type User []struct {
	Id               int      `json:"_id"`
	Alias            string   `json:"alias,omitempty"`
	ExternalId       string   `json:"external_id"`
	Name             string   `json:"name"`
	Signature        string   `json:"signature"`
	Email            string   `json:"email,omitempty"`
	Phone            string   `json:"phone"`
	Role             string   `json:"role"`
	Locale           string   `json:"locale"`
	CreatedAt        string   `json:"created_at"`
	LastLoginAt      string   `json:"last_login_at"`
	Timezone         string   `json:"timezone"`
	Shared           bool     `json:"shared"`
	Suspended        bool     `json:"suspended"`
	Active           bool     `json:"active"`
	Verified         bool     `json:"verified"`
	OrganizationId   int      `json:"organization_id,omitempty"`
	Tags             []string `json:"tags"`
	OrganizationName string   `json:",omitempty"`
	Tickets          []string `json:",omitempty"`
	Url              string   `json:"url"`
}

type UserData struct {
	Raw       []byte
	Processed User
	Filtered  User
}

var KeyMappings = map[string]string{
	"_id":               "Id",
	"alias":             "Alias",
	"external_id":       "ExternalId",
	"name":              "Name",
	"suspended":         "Suspended",
	"active":            "Active",
	"tags":              "Tags",
	"verified":          "Verified",
	"shared":            "Shared",
	"timezone":          "Timezone",
	"last_login_at":     "LastLoginAt",
	"organization_id":   "OrganizationId",
	"created_at":        "CreatedAt",
	"locale":            "Locale",
	"phone":             "Phone",
	"email":             "Email",
	"signature":         "Signature",
	"role":              "Role",
	"organization_name": "OrganizationName",
	"tickets":           "Tickets",
	"url":               "Url",
}

type UserSearchFlags struct {
	Value string
	Name  string `validate:"required,oneof=_id alias external_id name signature email phone role locale created_at last_login_at timezone details shared suspended active verified organization_id tags url"`
}

func (u UserSearchFlags) FetchName() string {
	return u.Name
}

func (u UserSearchFlags) FetchValue() string {
	return u.Value
}

// SetFiltered - Set the filtered list of User, and return the parent struct (UserData) they belong to
func (u *UserData) SetFiltered(values any) (internal.DataProcessor, error) {
	var result User
	jsonString, _ := json.Marshal(values)
	err := json.Unmarshal(jsonString, &result)
	if err != nil {
		return nil, err
	}
	u.Filtered = result
	return u, nil
}

// FetchFiltered - Get a filtered and processed list of User (not raw bytes)
func (u *UserData) FetchFiltered() internal.DataStore {
	return u.Filtered
}

// FetchProcessed - Get a processed list of User (not raw bytes)
func (u *UserData) FetchProcessed() []interface{} {
	var userData []interface{}
	for _, user := range u.Processed {
		userData = append(userData, user)
	}
	return userData
}

// FetchRaw - Get the raw User data
func (u *UserData) FetchRaw() []byte {
	return u.Raw
}

// Fetch - Get list of users (usually after by UserData#FetchFiltered method) is called
func (u User) Fetch() []interface{} {
	var allUsers []interface{}
	for _, ticket := range u {
		allUsers = append(allUsers, ticket)
	}
	return allUsers
}
