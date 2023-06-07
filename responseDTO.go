package main

// swagger:response ResponseConfig
type ResponseConfig struct {

	// Id of the config
	// in: string
	Id string `json:"id"`

	// Version of the config
	// in: string
	Version string `json:"version"`

	// List of labels of the config
	// in: string
	Labels string `json:"labels"`

	// nemam pojma
	// in: map[string]string
	Entries map[string]string `json:"entries"`

	// nemam pojma
	// in: string
	Group_Id string `json:"group_id"`

	// nemam pojma
	// in: string
	Group_Version string `json:"group_version"`
}

// swagger:response ResponseGroup
type ResponseGroup struct {

	// Id of the config
	// in: string
	Id string `json:"id"`

	// Version of the config
	// in: string
	Version string `json:"version"`

	// List of labels of the config
	// in: string
	Labels string `json:"labels"`

	// nemam pojma
	// in: map[string]string
	Entries map[string]string `json:"entries"`

	// nemam pojma
	// in: string
	Group_Id string `json:"group_id"`

	// nemam pojma
	// in: string
	Group_Version string `json:"group_version"`
}

// swagger:response ErrorResponse
type ErrorResponse struct {

	// Error status code
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
}

// swagger:response NoContentResponse
type NoContentResponse struct{}
