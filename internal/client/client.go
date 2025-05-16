package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Client represents an API client.
type Client struct {
	URL          string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
	HTTPClient   *http.Client
	accessToken  string
	authMethod   string
}

// ClientConfig holds the configuration for creating a new client.
type ClientConfig struct {
	URL          string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

// NewClient creates a new API client.
func NewClient(config *ClientConfig) (*Client, error) {
	client := &Client{
		URL:          strings.TrimSuffix(config.URL, "/"),
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		HTTPClient:   &http.Client{},
	}

	// Try client credentials
	err := client.authenticateClientCredentials()
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %w", err)
	}
	client.authMethod = "client_credentials"

	return client, nil
}

func (c *Client) authenticateClientCredentials() error {
	ctx := context.Background()
	tflog.Debug(ctx, "Authenticating with client credentials", map[string]interface{}{
		"url": fmt.Sprintf("%s/auth/token", c.URL),
	})

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", c.ClientID)
	data.Set("client_secret", c.ClientSecret)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/token", c.URL), strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("error creating auth request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making auth request: %w", err)
	}
	defer resp.Body.Close()

	// Read and log the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading auth response body: %w", err)
	}

	tflog.Debug(ctx, "Received auth response", map[string]interface{}{
		"status_code": resp.StatusCode,
		"body":        string(body),
	})

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("auth request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("error parsing auth response: %w", err)
	}

	accessToken, ok := result["access_token"].(string)
	if !ok {
		return fmt.Errorf("access_token not found in response")
	}

	c.accessToken = accessToken
	return nil
}

// setAuthHeader sets the appropriate authentication header based on the successful auth method.
func (c *Client) setAuthHeader(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.accessToken)
}

// CreateResource creates a new resource in the API.
func (c *Client) CreateResource(resourceType string, data interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data: %w", err)
	}
	tflog.Info(context.Background(), "Creating resource", map[string]interface{}{
		"resourceType": resourceType,
		"data":         string(jsonData),
	})

	// Log the request payload for debugging
	ctx := context.Background()
	tflog.Debug(ctx, "Making request to Aidbox", map[string]interface{}{
		"method":  "PUT",
		"url":     fmt.Sprintf("%s/%s", c.URL, resourceType),
		"payload": string(jsonData),
	})

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s", c.URL, resourceType), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.accessToken)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Read and log the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	tflog.Debug(ctx, "Received response from Aidbox", map[string]interface{}{
		"status_code": resp.StatusCode,
		"body":        string(body),
	})

	// Parse the response
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	tflog.Info(context.Background(), "[BILLINGBOX] Created resource", map[string]interface{}{
		"resourceType": resourceType,
		"id":           result["id"],
		"user":         result["user"],
		"name":         result["name"],
		"meta":         result["meta"],
	})

	return result, nil
}

// GetResource retrieves a resource from the API.
func (c *Client) GetResource(resourceType string, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/%s", c.URL, resourceType, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	tflog.Info(context.Background(), "Getting resource", map[string]interface{}{
		"resourceType": resourceType,
		"id":           id,
	})

	// Log the request details for debugging
	tflog.Debug(context.Background(), "GetResource Request", map[string]interface{}{
		"url":    url,
		"method": "GET",
		"headers": map[string]string{
			"Authorization": "Bearer " + c.accessToken,
		},
	})

	c.setAuthHeader(req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error getting resource: status %d, body: %s", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	// Log the response for debugging
	jsonResponse, _ := json.MarshalIndent(result, "", "  ")
	tflog.Debug(context.Background(), "GetResource Response", map[string]interface{}{
		"response": string(jsonResponse),
		"status":   resp.StatusCode,
		"headers":  resp.Header,
	})

	tflog.Info(context.Background(), "Got resource", map[string]interface{}{
		"resourceType": resourceType,
		"id":           result["id"],
		"user":         result["user"],
		"name":         result["name"],
		"meta":         result["meta"],
	})

	return result, nil
}

// UpdateResource updates an existing resource in the API.
func (c *Client) UpdateResource(resourceType string, id string, data interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data: %w", err)
	}

	url := fmt.Sprintf("%s/%s/%s", c.URL, resourceType, id)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	c.setAuthHeader(req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error updating resource: status %d, body: %s", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}

// DeleteResource deletes a resource from the API.
func (c *Client) DeleteResource(resourceType string, id string) error {
	url := fmt.Sprintf("%s/%s/%s", c.URL, resourceType, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	tflog.Info(context.Background(), "Deleting resource", map[string]interface{}{
		"resourceType": resourceType,
		"id":           id,
	})

	c.setAuthHeader(req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error deleting resource: status %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

// IsNotFoundError checks if the error indicates that a resource was not found.
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found")
}
