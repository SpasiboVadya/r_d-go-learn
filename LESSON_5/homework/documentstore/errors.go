package documentstore

import "errors"

var (
	ErrDocumentNotFound         = errors.New("document not found")
	ErrCollectionAlreadyExists  = errors.New("collection already exists")
	ErrCollectionNotFound       = errors.New("collection not found")
	ErrUnsupportedDocumentField = errors.New("unsupported document field")
	ErrInvalidDocument          = errors.New("invalid document")
	ErrInvalidFieldType         = errors.New("invalid field type")
)
