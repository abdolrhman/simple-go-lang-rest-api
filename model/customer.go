package model

// Data is mainle generated for filtering and pagination
type Data struct {
	TotalData    int64
	FilteredData int64
	Data         map[string][]Customer
}

type Args struct {
	Sort   string
	Order  string
	Offset string
	Limit  string
	Search string
}

type Customer struct {
	ID    int
	Name  string
	Phone string
}
