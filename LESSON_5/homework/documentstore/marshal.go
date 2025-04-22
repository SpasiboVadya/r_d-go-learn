package documentstore

import (
	"encoding/json"
	"fmt"
)

func MarshalDocument(input any) (*Document, error) {
	// Convert input to JSON bytes
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal input: %w", err)
	}

	// Unmarshal into a map
	var data map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal to map: %w", err)
	}

	// Create document with fields
	doc := &Document{
		Fields: make(map[string]DocumentField),
	}

	// Convert map values to DocumentField
	for key, value := range data {
		field, err := valueToDocumentField(value)
		if err != nil {
			return nil, fmt.Errorf("failed to convert field %s: %w", key, err)
		}
		doc.Fields[key] = field
	}

	return doc, nil
}

func UnmarshalDocument(doc *Document, output any) error {
	// Convert document to map
	data := make(map[string]interface{})
	for key, field := range doc.Fields {
		value, err := documentFieldToValue(field)
		if err != nil {
			return fmt.Errorf("failed to convert field %s: %w", key, err)
		}
		data[key] = value
	}

	// Convert map to JSON bytes
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal document: %w", err)
	}

	// Unmarshal into output
	if err := json.Unmarshal(jsonBytes, output); err != nil {
		return fmt.Errorf("failed to unmarshal to output: %w", err)
	}

	return nil
}

func valueToDocumentField(value interface{}) (DocumentField, error) {
	switch v := value.(type) {
	case string:
		return DocumentField{Type: DocumentFieldTypeString, Value: v}, nil
	case float64:
		return DocumentField{Type: DocumentFieldTypeNumber, Value: v}, nil
	case bool:
		return DocumentField{Type: DocumentFieldTypeBool, Value: v}, nil
	case []interface{}:
		return DocumentField{Type: DocumentFieldTypeArray, Value: v}, nil
	case map[string]interface{}:
		return DocumentField{Type: DocumentFieldTypeObject, Value: v}, nil
	default:
		return DocumentField{}, ErrUnsupportedDocumentField
	}
}

func documentFieldToValue(field DocumentField) (interface{}, error) {
	switch field.Type {
	case DocumentFieldTypeString, DocumentFieldTypeNumber, DocumentFieldTypeBool:
		return field.Value, nil
	case DocumentFieldTypeArray, DocumentFieldTypeObject:
		return field.Value, nil
	default:
		return nil, ErrInvalidFieldType
	}
}
