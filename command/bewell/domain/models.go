package domain

//RegistryResponse is the response got when making schema registry requests
type RegistryResponse struct {
	Success bool              `json:"success,omitempty"`
	Message string            `json:"message,omitempty"`
	Details []ResponseDetails `json:"details,omitempty"`
}

// ResponseDetails is the response got when making schema registry requests
// Usually contains a reason why the operation/request failed
type ResponseDetails struct {
	Message string `json:"message,omitempty"`
}

// GraphqlSchemaPayload is the payload made when making schema registry requests
type GraphqlSchemaPayload struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Version  string `json:"version"`
	TypeDefs string `json:"type_defs"`
}
