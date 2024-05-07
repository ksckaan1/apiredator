package domain

type Data struct {
	Request Request `json:"request"`
	Options Options `json:"options"`
}

type Request struct {
	Method string         `json:"method"`
	URL    string         `json:"url"`
	Header []KeyValueData `json:"header"`
	Body   Body           `json:"body"`
}

type KeyValueData struct {
	IsActive bool   `json:"is_active"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

type Body struct {
	Type         string         `json:"type"`
	Language     string         `json:"language"`
	RawValue     string         `json:"raw_value"`
	Binary       []string       `json:"binary"`
	FormData     []FormData     `json:"formdata"`
	XWWWFormData []KeyValueData `json:"xwwwformdata"`
}

type FormData struct {
	IsActive  bool     `json:"is_active"`
	Key       string   `json:"key"`
	TextValue string   `json:"text_value"`
	FileValue []string `json:"file_value"`
	RowType   string   `json:"row_type"`
}

type Options struct {
	TestType         TestType `json:"test_type"`
	TestDuration     string   `json:"test_duration"`
	RequestTimeout   string   `json:"request_timeout"`
	NumberOfClients  uint64   `json:"number_of_clients"`
	NumberOfRequests uint64   `json:"number_of_requests"`
}

type TestType string

const (
	TTCount    TestType = "count"
	TTDuration TestType = "duration"
)
