package main

import (
	"encoding/json"
)

// Root of the JSON Schema
type Root struct {
	Schema
	SchemaType string `json:"$schema"`
}

// Schema JSON schema.
type Schema struct {
	Title       string   `json:"title"`
	ID          string   `json:"id"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Readonly    bool     `json:"readonly"`
	Enums       []string `json:"enum"`
	Definitions map[string]*Schema
	Properties  map[string]*Schema
	Reference   string `json:"$ref"`
}

func getPrimitiveTypeName(schemaType string) string {
	switch schemaType {
	case "array":
		return ""
	case "boolean":
		return "bool"
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "null":
		return ""
	case "object":
		return ""
	case "string":
		return "string"
	}

	return "undefined"
}

// Parse parses a JSON schema from a string.
func Parse(schema string) (*Root, error) {
	s := &Root{}
	err := json.Unmarshal([]byte(schema), s)

	return s, err
}
