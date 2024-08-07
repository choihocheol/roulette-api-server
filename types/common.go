package types

type QueryFilter struct {
	Key   string
	Value string
}

type QueryFilterMap map[string]string

type Counter struct {
	Cnt uint64 `json:"cnt" db:"cnt"`
}
