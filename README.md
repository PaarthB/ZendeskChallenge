# Zendesk Challenge - Search CLI

![Build Status](https://raw.githubusercontent.com/dwyl/repo-badges/main/svg/build-passing.svg)
![Go version](https://img.shields.io/badge/Go_version-1.21-blue)
![Code Coverage](https://img.shields.io/badge/Code_coverage-90_percent-b6d7a8)

### System Requirements
- Go `1.21`
- Environment variables for correct usage of go are set, such as `GOROOT`, `GOPATH` and `GO111MODULE=on` (default)
- Tested on `Mac Sonoma M1`. Should work fine with any `darwin-arm64` architecture.

### Usage Instructions
- Open command line in root directory of this repo.
- Run `make build`
- Then run `make setup`, this will move the `cli` executable to current director
- To make it easier, a copy of `cli` executable is left in the root directory, for being able to run directly.

#### Listing searchable fields 
- Run `./cli list` for getting all possible searchable fields

#### Searching
- Searching cane be done via `./cli search` command. Type `--help` to see usage
- To search for empty fields, don't specify `--value` flag, as the CLI then treats it as empty string (**NOTE**: be careful they are empty strings, for values that require integers, it will still error saying you need to specify int value, as empty value cannot be int)
- Eg: `cli search user --name _id --value 1` searches for user with `_id` attribute as `1`, shows output:
```
======== All results ========
------------------------------------------------
tags_0: Springville
tags_1: Sutton
tags_2: Hartsville/Hartley
tags_3: Diaperville
verified: true
created_at: 2016-04-15T05:19:46 -10:00
external_id: 74341f74-9c79-49d5-9611-87ef9b6eb75f
name: Francisca Rasmussen
last_login_at: 2013-08-04T01:03:27 -10:00
email: coffeyrasmussen@flotonic.com
_id: 1
suspended: true
shared: false
role: admin
organization_name: Multron
alias: Miss Coffey
active: true
timezone: Sri Lanka
organization_id: 119
locale: en-AU
phone: 8335-422-718
signature: Don't Worry Be Happy!
tickets_0: Ipsum reprehenderit non ea officia labore aute. Qui sit aliquip ipsum nostrud anim qui pariatur ut anim aliqua non aliqua.
tickets_1: Nostrud veniam eiusmod reprehenderit adipisicing proident aliquip. Deserunt irure deserunt ea nulla cillum ad.
```
### Testing Instructions
- Run `make test`
- All tests are defined within the individual packages themselves
- For test coverage, run `make coverage`, see output below:
```
go tool cover -func profile.cov               
ZendeskChallenge/cmd/list/handler.go:14:                        NewListCmd                      100.0%
ZendeskChallenge/cmd/list/list.go:15:                           fieldList                       100.0%
ZendeskChallenge/cmd/search/handler.go:13:                      NewSearchCmd                    100.0%
ZendeskChallenge/cmd/search/handler.go:25:                      NewUserSearchCmd                100.0%
ZendeskChallenge/cmd/search/handler.go:39:                      NewTicketSearchCmd              100.0%
ZendeskChallenge/cmd/search/handler.go:53:                      NewOrgSearchCmd                 100.0%
ZendeskChallenge/cmd/search/process.go:22:                      addRelatedUserEntities          100.0%
ZendeskChallenge/cmd/search/process.go:57:                      addRelatedTicketEntities        100.0%
ZendeskChallenge/cmd/search/process.go:97:                      evaluateSearch                  91.7%
ZendeskChallenge/cmd/search/process.go:115:                     evaluateSearchResultByDataType  85.7%
ZendeskChallenge/cmd/search/process.go:161:                     parseByDataType                 100.0%
ZendeskChallenge/cmd/search/search.go:27:                       getFileData                     100.0%
ZendeskChallenge/cmd/search/search.go:36:                       triggerUserSearch               73.9%
ZendeskChallenge/cmd/search/search.go:74:                       triggerTicketSearch             73.9%
ZendeskChallenge/cmd/search/search.go:110:                      triggerOrgSearch                68.4%
ZendeskChallenge/internal/output.go:14:                         DisplayResults                  100.0%
ZendeskChallenge/models/organizations/organizations.go:46:      FetchFiltered                   100.0%
ZendeskChallenge/models/organizations/organizations.go:50:      SetFiltered                     85.7%
ZendeskChallenge/models/organizations/organizations.go:61:      FetchProcessed                  100.0%
ZendeskChallenge/models/organizations/organizations.go:69:      FetchRaw                        100.0%
ZendeskChallenge/models/organizations/organizations.go:73:      Fetch                           100.0%
ZendeskChallenge/models/tickets/tickets.go:61:                  SetFiltered                     85.7%
ZendeskChallenge/models/tickets/tickets.go:72:                  FetchFiltered                   100.0%
ZendeskChallenge/models/tickets/tickets.go:76:                  FetchProcessed                  100.0%
ZendeskChallenge/models/tickets/tickets.go:84:                  FetchRaw                        100.0%
ZendeskChallenge/models/tickets/tickets.go:88:                  Fetch                           100.0%
ZendeskChallenge/models/users/users.go:70:                      SetFiltered                     85.7%
ZendeskChallenge/models/users/users.go:81:                      FetchFiltered                   100.0%
ZendeskChallenge/models/users/users.go:85:                      FetchProcessed                  100.0%
ZendeskChallenge/models/users/users.go:93:                      FetchRaw                        100.0%
ZendeskChallenge/models/users/users.go:97:                      Fetch                           100.0%
total:                                                          (statements)                    90.0%
```
### Design tradeoffs
#### Searching through JSON
- For searching through JSON efficiently, `JSONPath` query language (similar to `XPath` for XML) was used. It is quite efficient, and more about it can be read [here](https://goessner.net/articles/JsonPath/) 
- This query language is designed for keeping memory overhead small, and searches efficient, without linear increase in time as more data is added.
- Specific queries were formulated, which can be seen in `cmd/search/process.go` to find entries in JSON that match what we are looking for (item in list, or key equals value where value can be of various types like bool/string/integer etc.)

#### Adding related entities
1. When searching for users
   1. All ticket descriptions of tickets that this user has submitted, are shown
   2. Name of organization is shown

2. When searching for tickets
   1. Assignee name and submitter name is shown
   2. Organization name is shown


