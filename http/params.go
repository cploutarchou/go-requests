package http

type QueryParams interface {
	// Add adds the key, value pair to the query params.
	// If the key already exists, the value will be appended to the existing values.
	Add(key, value string)

	// Set sets the key, value pair to the query params.
	// If the key already exists, the value will be replaced with the new value.
	Set(key, value string)

	// Get gets the first value associated with the given key.
	// If there are no values associated with the key, Get returns "".
	// To access multiple values, use the map directly.
	Get(key string) string

	// Del deletes the values associated with key.
	Del(key string)

	// Values returns the values map.
	Values() map[string]string

	// Clone returns a copy of the QueryParams.
	Clone()

	// Reset resets the QueryParams to the initial state.
	Reset()
}

// implementation of QueryParams
type queryParams struct {
	values map[string]string
	// contains filtered or unexported fields

}

func (q queryParams) Add(key, value string) {
	// add the key, value pair to the query params.
	// If the key already exists, the value will be appended to the existing values. if not, create a new one
	q.values[key] = value
}

func (q queryParams) Set(key, value string) {
	// Set sets the key, value pair to the query params.
	// If the key already exists, the value will be replaced with the new value.
	q.values[key] = value

}

func (q queryParams) Get(key string) string {
	// Get gets the first value associated with the given key.
	// If there are no values associated with the key, Get returns "".
	// To access multiple values, use the map directly.
	if len(q.values[key]) > 0 {
		return q.values[key]
	}
	return ""
}

func (q queryParams) Del(key string) {
	// Del deletes the values associated with key.
	delete(q.values, key)

}

func (q queryParams) Values() map[string]string {

	return q.values
}

func (q queryParams) Clone() {
	//clone the queryParams
	clone := NewQueryParams()
	for key, values := range q.values {
		for _, value := range values {
			clone.Add(key, string(value))
		}
	}
}

func (q queryParams) Reset() {
	q.values = make(map[string]string)
}

func NewQueryParams() QueryParams {
	return &queryParams{
		values: make(map[string]string),
	}
}
