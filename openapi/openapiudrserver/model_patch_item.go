/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PatchItem - it contains information on data to be changed.
type PatchItem struct {
	Op PatchOperation `json:"op"`

	// contains a JSON pointer value (as defined in IETF RFC 6901) that references a location of a resource on which the patch operation shall be performed.
	Path string `json:"path"`

	// indicates the path of the source JSON element (according to JSON Pointer syntax) being moved or copied to the location indicated by the \"path\" attribute.
	From string `json:"from,omitempty"`

	Value *interface{} `json:"value,omitempty"`
}

// AssertPatchItemRequired checks if the required fields are not zero-ed
func AssertPatchItemRequired(obj PatchItem) error {
	elements := map[string]interface{}{
		"op":   obj.Op,
		"path": obj.Path,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPatchOperationRequired(obj.Op); err != nil {
		return err
	}
	return nil
}

// AssertPatchItemConstraints checks if the values respects the defined constraints
func AssertPatchItemConstraints(obj PatchItem) error {
	return nil
}
