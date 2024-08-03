package testcase

import (
	"bytes"
	"strings"
	"text/template"
)

type RequestType string

const (
	HTTP      RequestType = "HTTP"
	HTTPS     RequestType = "HTTPS"
	WebSocket RequestType = "WebSocket"
	FTP       RequestType = "FTP"
)

type Assertion struct {
	Type   string
	Target string
	Value  interface{}
}

type TestCase struct {
	Name        string
	Description string
	RequestType RequestType
	URL         string
	Method      string
	Headers     map[string]string
	Body        string
	Params      map[string]string
	Assertions  []Assertion
}

func (tc *TestCase) Render(params map[string]string) (string, error) {
	t, err := template.New("body").Parse(tc.Body)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, params)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (tc *TestCase) ValidateResponse(response string, params map[string]string) bool {
	for _, assertion := range tc.Assertions {
		switch assertion.Type {
		case "contains":
			if !strings.Contains(response, assertion.Value.(string)) {
				return false
			}
		case "equals":
			if response != assertion.Value.(string) {
				return false
			}
			// More assertion types...
		}
	}
	return true
}

type TestCaseStore struct {
	testCases map[string]*TestCase
}

func NewTestCaseStore() *TestCaseStore {
	return &TestCaseStore{
		testCases: make(map[string]*TestCase),
	}
}

func (store *TestCaseStore) AddTestCase(tc *TestCase) {
	store.testCases[tc.Name] = tc
}

func (store *TestCaseStore) GetTestCase(name string) (*TestCase, bool) {
	tc, exists := store.testCases[name]
	return tc, exists
}

func (store *TestCaseStore) DeleteTestCase(name string) {
	delete(store.testCases, name)
}

func (store *TestCaseStore) ListTestCases() []*TestCase {
	tcs := make([]*TestCase, 0, len(store.testCases))
	for _, tc := range store.testCases {
		tcs = append(tcs, tc)
	}
	return tcs
}

func NewTestCase(name string, request RequestType) *TestCase {
	return &TestCase{
		Name:        name,
		RequestType: request,
	}
}
