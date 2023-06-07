package configstore

// // swagger:model ConfigGroup
// type ConfigGroup struct {

// 	// Version of the group
// 	// in: string
// 	ServiceVersion string `json:"version"`

// 	// Configs in the group
// 	// in: []*Config
// 	Data []*Config `json:"group"`
// }

// // swagger:model Config
// type Config struct {

// 	// Id of the config
// 	// in: string
// 	Id string `json:"id"`

// 	// Version of the config
// 	// in: string
// 	Version string `json:"version"`

// 	// List of labels of the config
// 	// in: string
// 	Labels string `json:"labels"`

// 	// nemam pojma
// 	// in: map[string]string
// 	Entries map[string]string `json:"entries"`
// }

type Config struct {

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


