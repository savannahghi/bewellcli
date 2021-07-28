package dto

// SchemaStatus holds Status and message(if any) for a schema request
type SchemaStatus struct {
	Valid   bool
	Message string
}
