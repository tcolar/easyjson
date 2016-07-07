package easyjson

// JsonMapping allows hiding/renaming/adding Json fields during runtime encoding
type JsonMapping struct {
	mapped    map[string]string
	added     map[string]string
	inclusive bool
}

// Mapping creates a new empty mapping
func Mapping() *JsonMapping {
	return &JsonMapping{
		mapped: map[string]string{},
		added:  map[string]string{},
	}
}

// Only tells the encoder to only include the given fields (by name)
// incompatible with Omit()
func (j *JsonMapping) Only(fields ...string) *JsonMapping {
	j.inclusive = true
	for _, field := range fields {
		j.mapped[field] = ""
	}
	return j
}

// Drop tells the encoder to omit the given field (by name)
func (j *JsonMapping) Omit(fields ...string) *JsonMapping {
	j.inclusive = false
	for _, field := range fields {
		j.mapped[field] = ""
	}
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
// a returned value of empty string means Omit the field
func (j *JsonMapping) Name(field, jsonName string) string {
	if j == nil {
		return jsonName
	}
	if !j.inclusive { // exclusive "omit()"
		newName, found := j.mapped[field]
		if !found {
			return jsonName
		}
		return newName
	}
	// inclusive "only()"
	newName, found := j.mapped[field]
	if found {
		if len(newName) > 0 {
			return newName
		}
		return jsonName
	}
	return ""
}

// Added returns the added items
func (j *JsonMapping) Added() map[string]string {
	if j == nil {
		return map[string]string{}
	}
	return j.added
}
