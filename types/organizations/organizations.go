package organizations

import (
	"ZendeskChallenge/internal"
	"encoding/json"
)

type Organization []struct {
	Id            int      `json:"_id"`
	Url           string   `json:"url"`
	ExternalId    string   `json:"external_id"`
	Name          string   `json:"name"`
	DomainNames   []string `json:"domain_names"`
	CreatedAt     string   `json:"created_at"`
	Details       string   `json:"details"`
	SharedTickets bool     `json:"shared_tickets"`
	Tags          []string `json:"tags"`
}

type OrgData struct {
	Raw       []byte
	Processed Organization
	Filtered  Organization
}

var KeyMappings = map[string]string{
	"_id":            "Id",
	"url":            "Url",
	"external_id":    "ExternalId",
	"name":           "Name",
	"domain_names":   "DomainNames",
	"created_at":     "CreatedAt",
	"details":        "Details",
	"shared_tickets": "SharedTickets",
	"tags":           "Tags",
}

type OrganizationFlags struct {
	Value string
	Name  string `validate:"required,oneof=_id url external_id name domain_names details shared_tickets tickets"`
}

func (o *OrgData) FetchFiltered() internal.DataStore {
	return o.Filtered
}

func (o *OrgData) SetFiltered(values any) (internal.DataProcessor, error) {
	var result Organization
	jsonString, _ := json.Marshal(values)
	err := json.Unmarshal(jsonString, &result)
	if err != nil {
		return nil, err
	}
	o.Filtered = result
	return o, nil
}

func (o *OrgData) FetchProcessed() []interface{} {
	var orgData []interface{}
	for _, user := range o.Processed {
		orgData = append(orgData, user)
	}
	return orgData
}

func (o *OrgData) FetchRaw() []byte {
	return o.Raw
}

func (o Organization) Fetch() []interface{} {
	var allOrgs []interface{}
	for _, ticket := range o {
		allOrgs = append(allOrgs, ticket)
	}
	return allOrgs
}
