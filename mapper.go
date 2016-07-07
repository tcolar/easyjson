package easyjson

type JsonMapper struct {
	m map[string]string
}

func NewMapper() *JsonMapper {
	return &JsonMapper{
		m: map[string]string{},
	}
}

func (j *JsonMapper) Drop(field string) *JsonMapper {
	j.m[field] = ""
	return j
}

func (j *JsonMapper) Rename(field, to string) *JsonMapper {
	j.m[field] = to
	return j
}

func (j *JsonMapper) Add(name, value string) *JsonMapper {
	j.m[name] = value
	return j
}

func (j *JsonMapper) MappedName(field string) string {
	if j == nil || j.m == nil {
		return field
	}
	newName, found := j.m[field]
	if !found {
		return field
	}
	return newName
}
