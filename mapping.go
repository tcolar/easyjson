package easyjson

// JsonMapping allows hiding/renaming/adding Json fields during runtime encoding
type JsonMapping struct {
	mapped map[string]string
	added  map[string]string
}

// Mapping creates a new empty mapping
func Mapping() *JsonMapping {
	return &JsonMapping{
		mapped: map[string]string{},
		added:  map[string]string{},
	}
}

// Drop tells the encoder to omit the given field (by name)
func (j *JsonMapping) Omit(field string) *JsonMapping {
	j.mapped[field] = ""
	return j
}

// Rename tells the encoder to use a different name for a given field.
func (j *JsonMapping) Rename(field, to string) *JsonMapping {
	j.mapped[field] = to
	return j
}

// Add tells the encoder to add some extra items when encoding
// jsonVal is assumed already JSON encoded, ie: "true", "5", "\"foo\""
func (j *JsonMapping) Add(name string, jsonVal string) *JsonMapping {
	j.added[name] = jsonVal
	return j
}

// Name returns the mapped name for a field
// a returned value of empty string meand Omit the field
func (j *JsonMapping) Name(field string) string {
	if j == nil {
		return field
	}
	newName, found := j.mapped[field]
	if !found {
		return field
	}
	return newName
}

// Added returns the added items
func (j *JsonMapping) Added() map[string]string {
	return j.added
}
