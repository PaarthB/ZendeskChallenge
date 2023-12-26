package users

import (
	"ZendeskChallenge/internal"
	"encoding/json"
)

type User []struct {
	Id               int      `json:"_id"`
	Alias            string   `json:"alias"`
	ExternalId       string   `json:"external_id"`
	Name             string   `json:"name"`
	Signature        string   `json:"signature"`
	Email            string   `json:"email"`
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
	OrganizationId   int      `json:"organization_id"`
	Tags             []string `json:"tags"`
	OrganizationName string   `json:",omitempty"`
	Tickets          []string `json:",omitempty"`
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
}

type UserFlags struct {
	Value string
	Name  string `validate:"required,oneof=_id alias external_id name signature email phone role locale created_at last_login_at timezone details shared suspended active verified organization_id tags"`
}

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

func (u *UserData) FetchFiltered() internal.DataStore {
	return u.Filtered
}

func (u *UserData) FetchProcessed() []interface{} {
	var userData []interface{}
	for _, user := range u.Processed {
		userData = append(userData, user)
	}
	return userData
}

func (u *UserData) FetchRaw() []byte {
	return u.Raw
}

func (u User) Fetch() []interface{} {
	var allUsers []interface{}
	for _, ticket := range u {
		allUsers = append(allUsers, ticket)
	}
	return allUsers
}
