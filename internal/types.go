// Package internal -
//
// Defines all essential interfaces that multiple underlying models implement, to make code simple and concise,
// and conform to standards.
package internal

// DataProcessor - Defines common methods that all model-helper structs conform to,
// to allow fetching filtered, processed, raw data at any time
type DataProcessor interface {
	// FetchFiltered - Get a filtered and processed list of underlying entity implementation (not raw bytes)
	FetchFiltered() DataStore

	// SetFiltered - Set the filtered list of underlying entity implementation, and return the parent struct they belong to
	SetFiltered(values any) (DataProcessor, error)

	// FetchProcessed - Get a processed list of underlying entity implementation (not raw bytes)
	FetchProcessed() []interface{}

	// FetchRaw - Get the raw byte data of underlying entity implementation
	FetchRaw() []byte
}

// DataStore - For ensuring underlying model can fetch the final list of resulting entities, before displaying
type DataStore interface {
	// Fetch - Get list of underlying entity (usually after by FetchFiltered method) is called
	Fetch() []interface{}
}
