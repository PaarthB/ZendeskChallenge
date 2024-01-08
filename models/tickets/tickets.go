// Package tickets -
//
// Defines the ticket model and its key fields. Implements DataStore and DataProcessor interfaces
package tickets

import (
	"ZendeskChallenge/internal"
	"encoding/json"
)

type Ticket []struct {
	Id               string   `json:"_id"`
	ExternalId       string   `json:"external_id"`
	Type             string   `json:"type"`
	Description      string   `json:"description,omitempty"`
	Priority         string   `json:"priority"`
	Status           string   `json:"status"`
	Subject          string   `json:"subject"`
	OrganizationId   int      `json:"organization_id"`
	SubmitterId      int      `json:"submitter_id"`
	AssigneeId       int      `json:"assignee_id,omitempty"`
	CreatedAt        string   `json:"created_at"`
	HasIncidents     bool     `json:"has_incidents"`
	DueAt            string   `json:"due_at"`
	Via              string   `json:"via"`
	Tags             []string `json:"tags"`
	SubmitterName    string   `json:",omitempty"`
	AssigneeName     string   `json:",omitempty"`
	OrganizationName string   `json:",omitempty"`
	Url              string   `json:"url"`
}

type TicketData struct {
	Raw       []byte
	Processed Ticket
	Filtered  Ticket
}

var KeyMappings = map[string]string{
	"_id":             "Id",
	"external_id":     "ExternalId",
	"type":            "Type",
	"description":     "Description",
	"priority":        "Priority",
	"status":          "Status",
	"subject":         "Subject",
	"organization_id": "OrganizationId",
	"submitter_id":    "SubmitterId",
	"assignee_id":     "AssigneeId",
	"created_at":      "CreatedAt",
	"has_incidents":   "HasIncidents",
	"due_at":          "DueAt",
	"via":             "Via",
	"tags":            "Tags",
	"url":             "Url",
}

type TicketSearchFlags struct {
	Value string
	Name  string `validate:"required,oneof=_id external_id type description priority status subject created_at due_at submitter_id assignee_id has_incidents via tags organization_id url"`
}

func (t TicketSearchFlags) FetchName() string {
	return t.Name
}

func (t TicketSearchFlags) FetchValue() string {
	return t.Value
}

// SetFiltered - Set the filtered list of Ticket, and return the parent struct (TicketData) they belong to
func (t *TicketData) SetFiltered(values any) (internal.DataProcessor, error) {
	var result Ticket
	jsonString, _ := json.Marshal(values)
	err := json.Unmarshal(jsonString, &result)
	if err != nil {
		return nil, err
	}
	t.Filtered = result
	return t, nil
}

// FetchFiltered - Get a filtered and processed list of Ticket (not raw bytes)
func (t *TicketData) FetchFiltered() internal.DataStore {
	return t.Filtered
}

// FetchProcessed - Get a processed list of Ticket (not raw bytes)
func (t *TicketData) FetchProcessed() []interface{} {
	var ticketData []interface{}
	for _, user := range t.Processed {
		ticketData = append(ticketData, user)
	}
	return ticketData
}

// FetchRaw - Get the raw Ticket data
func (t *TicketData) FetchRaw() []byte {
	return t.Raw
}

// Fetch - Get list of Ticket's (usually after by *TicketData#FetchFiltered method) is called
func (t Ticket) Fetch() []interface{} {
	var allTickets []interface{}
	for _, ticket := range t {
		allTickets = append(allTickets, ticket)
	}
	return allTickets
}
