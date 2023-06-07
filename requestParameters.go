package main

import cs "github.com/cvule25/airs-projekat/configstore"

// swagger:parameters config createConfig
type RequestConfigBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/Config"
	//  required: true
	Body cs.Config `json:"body"`
}

// swagger:parameters getConfigById
type GetConfigById struct {
	// Config ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters getConfigByLabel
type GetConfigByLabel struct {

	// Config ID
	//
	// in: path
	Id string `json:"id"`

	//Config version
	//
	// in: path
	Version string `json:"version"`

	//Config labels
	//
	// in: path
	Labels string `json:"labels"`
}

// swagger:parameters deleteConfig
type DeleteConfig struct {
	// Config ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters group createGroup
type RequestGroupBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/[]Config"
	//  required: true
	Body []*cs.Config `json:"body"`
}

// swagger:parameters getGroup
type GetGroup struct {
	// Group ID
	// in: path
	Id string `json:"id"`

	//Config version
	//
	// in: path
	Version string `json:"version"`
}

// swagger:parameters deleteGroup
type DeleteGroup struct {
	// Group ID
	// in: path
	Id string `json:"id"`

	//Config version
	//
	// in: path
	Version string `json:"version"`
}
