package gql_payload

import (
	"bytes"
	"encoding/json"
)

type GraphQL struct {
	q    string
	vars map[string]interface{}
}

// NewField makes a new gql object.
func NewField(q string) *GraphQL {
	g := &GraphQL{
		q: q,
	}
	return g
}

// Var sets a variable.
func (g *GraphQL) Var(key string, value interface{}) *GraphQL {
	if g.vars == nil {
		g.vars = make(map[string]interface{})
	}
	g.vars[key] = value

	return g
}

//GenerateRequestBody generates the JSON encoded request body with:
// Query (Required)
// Variables (Optional)
// ref: https://graphql.org/learn/queries/#variables
func (g *GraphQL) GenerateRequestBody() (*bytes.Buffer, error) {
	var requestBody bytes.Buffer
	requestBodyObj := struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables"`
	}{
		Query:     g.q,
		Variables: g.vars,
	}
	if err := json.NewEncoder(&requestBody).Encode(requestBodyObj); err != nil {
		return nil, err
	}
	return &requestBody, nil
}
