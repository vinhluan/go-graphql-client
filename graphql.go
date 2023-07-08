package graphql

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
	"golang.org/x/net/context/ctxhttp"
)

//go:generate mockgen -destination=./mock/graphql.go -package=mock . GraphQL
type GraphQL interface {
	QueryString(ctx context.Context, q string, variables map[string]interface{}, v interface{}) (*Result, error)
	Query(ctx context.Context, q interface{}, variables map[string]interface{}) (*Result, error)

	Mutate(ctx context.Context, m interface{}, variables map[string]interface{}) (*Result, error)
	MutateString(ctx context.Context, m string, variables map[string]interface{}, v interface{}) (*Result, error)
}

type Result struct {
	Data       *json.RawMessage
	Errors     OpErrors
	Extensions map[string]interface{}
}

// Client is a GraphQL client.
type Client struct {
	url        string // GraphQL server URL.
	httpClient *http.Client
}

// NewClient creates a GraphQL client targeting the specified GraphQL server URL.
// If httpClient is nil, then http.DefaultClient is used.
func NewClient(url string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		url:        url,
		httpClient: httpClient,
	}
}

// QueryString executes a single GraphQL query request,
// using the given raw query `q` and populating the response into the `v`.
// `q` should be a correct GraphQL query request string that corresponds to the GraphQL schema.
func (c *Client) QueryString(ctx context.Context, q string, variables map[string]interface{}, v interface{}) (*Result, error) {
	return c.do(ctx, q, variables, v)
}

// Query executes a single GraphQL query request,
// with a query derived from q, populating the response into it.
// q should be a pointer to struct that corresponds to the GraphQL schema.
func (c *Client) Query(ctx context.Context, q interface{}, variables map[string]interface{}) (*Result, error) {
	query := constructQuery(q, variables)
	return c.do(ctx, query, variables, q)
}

// Mutate executes a single GraphQL mutation request,
// with a mutation derived from m, populating the response into it.
// m should be a pointer to struct that corresponds to the GraphQL schema.
func (c *Client) Mutate(ctx context.Context, m interface{}, variables map[string]interface{}) (*Result, error) {
	query := constructMutation(m, variables)
	return c.do(ctx, query, variables, m)
}

// MutateString executes a single GraphQL mutation request,
// using the given raw query `m` and populating the response into it.
// `m` should be a correct GraphQL mutation request string that corresponds to the GraphQL schema.
func (c *Client) MutateString(ctx context.Context, m string, variables map[string]interface{}, v interface{}) (*Result, error) {
	return c.do(ctx, m, variables, v)
}

// do executes a single GraphQL operation.
func (c *Client) do(ctx context.Context, query string, variables map[string]interface{}, v interface{}) (result *Result, err error) {
	in := struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables,omitempty"`
	}{
		Query:     query,
		Variables: variables,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	resp, err := ctxhttp.Post(ctx, c.httpClient, c.url, "application/json", &buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 OK status code: %v", resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		// TODO: Consider including response body in returned error, if deemed helpful.
		return nil, err
	}
	if result.Data != nil {
		err = json.Unmarshal(*result.Data, v)
		if err != nil {
			// TODO: Consider including response body in returned error, if deemed helpful.
			return nil, err
		}
	}
	if len(result.Errors) > 0 {
		err = result.Errors
	}
	// Returns Result struct which contains parsed response body to expose raw data, errors and extensions fields
	return
}

// OpErrors represents the "errors" array in a response from a GraphQL server.
// If returned via error interface, the slice is expected to contain at least 1 element.
//
// Specification: https://facebook.github.io/graphql/#sec-Errors.
type OpErrors []struct {
	Message   string
	Locations []struct {
		Line   int
		Column int
	}
}

// Error implements error interface.
func (e OpErrors) Error() string {
	return e[0].Message
}

type operationType uint8

const (
	queryOperation operationType = iota
	mutationOperation
	// subscriptionOperation // Unused.
)
