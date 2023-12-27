// Package internal -
//
// Defines all essential interfaces that multiple underlying models implement, to make code simple and concise,
// and conform to standards.
package internal

type DataProcessor interface {
	FetchFiltered() DataStore
	SetFiltered(values any) (DataProcessor, error)
	FetchProcessed() []interface{}
	FetchRaw() []byte
}

type DataStore interface {
	Fetch() []interface{}
}
