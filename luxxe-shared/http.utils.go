package shared

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type HttpUtilsReq struct {
	BaseRoute string
	Ext       string
	Query     map[string]string
	Body      map[string]interface{}
	Token     string
}

// Function to handle HTTP requests for all methods (POST, GET, PUT, DELETE)
func sendRequest[T any](method string, baseAPIReq HttpUtilsReq) (T, error) {
	baseRoute, ext, query, body, token := baseAPIReq.BaseRoute, baseAPIReq.Ext, baseAPIReq.Query, baseAPIReq.Body, baseAPIReq.Token

	// Build the URL with query parameters if any
	baseURL, err := url.Parse(baseRoute + ext)
	if err != nil {
		var zeroValue T
		return zeroValue, err
	}

	// Add query parameters to the URL
	params := url.Values{}
	for key, value := range query {
		params.Add(key, value)
	}
	baseURL.RawQuery = params.Encode()

	// Only marshal the body if it's not empty and method is not GET or DELETE
	var jsonBody []byte
	if method != http.MethodGet && method != http.MethodDelete && body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			var zeroValue T
			return zeroValue, err
		}
	}

	// Create a new HTTP request
	req, err := http.NewRequest(method, baseURL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		var zeroValue T
		return zeroValue, err
	}
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		var zeroValue T
		return zeroValue, err
	}
	defer resp.Body.Close()

	// Read the raw response body
	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		var zeroValue T
		return zeroValue, err
	}

	// fmt.Println("Raw Response:", string(rawBody)) // Print raw response body for debugging

	// Unmarshal the response into the specified type
	var res T
	err = json.Unmarshal(rawBody, &res)
	if err != nil {
		var zeroValue T
		return zeroValue, err
	}

	return res, nil
}

// POST request
func POST[T any](baseAPIReq HttpUtilsReq) (T, error) {
	return sendRequest[T](http.MethodPost, baseAPIReq)
}

// GET request
func GET[T any](baseAPIReq HttpUtilsReq) (T, error) {
	return sendRequest[T](http.MethodGet, baseAPIReq)
}

// PUT request
func PUT[T any](baseAPIReq HttpUtilsReq) (T, error) {
	return sendRequest[T](http.MethodPut, baseAPIReq)
}

// DELETE request
func DELETE[T any](baseAPIReq HttpUtilsReq) (T, error) {
	return sendRequest[T](http.MethodDelete, baseAPIReq)
}
