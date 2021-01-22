package datauri

import "fmt"

type SchemaNotReigsteredError struct {
	Schema string
}

func (e *SchemaNotReigsteredError) Error() string {
	return fmt.Sprintf("datauri: schema [%s] not registered", e.Schema)
}
func NewSchemaNotReigsteredError(schema string) *SchemaNotReigsteredError {
	return &SchemaNotReigsteredError{Schema: schema}
}
