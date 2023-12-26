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
