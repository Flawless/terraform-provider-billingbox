package client

import (
	"os"
	"testing"
)

func TestClientConnectivity(t *testing.T) {
	// Skip if not running integration tests
	if os.Getenv("INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test. Set INTEGRATION_TESTS=true to run.")
	}

	// Get Aidbox configuration from environment
	url := os.Getenv("AIDBOX_URL")
	if url == "" {
		t.Fatal("AIDBOX_URL environment variable is not set")
	}

	clientID := os.Getenv("AIDBOX_CLIENT_ID")
	if clientID == "" {
		t.Fatal("AIDBOX_CLIENT_ID environment variable is not set")
	}

	clientSecret := os.Getenv("AIDBOX_CLIENT_SECRET")
	if clientSecret == "" {
		t.Fatal("AIDBOX_CLIENT_SECRET environment variable is not set")
	}

	// Create client configuration
	config := &ClientConfig{
		URL:          url,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	// Create client
	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Try to get a resource to verify connectivity
	// Using a non-existent resource type to avoid side effects
	result, err := client.GetResource("TestResource", "test-id")
	if err != nil {
		// We expect a 404 error, but not a connection error
		if err.Error() == "error getting resource: status 404" {
			t.Log("Successfully connected to Aidbox (got expected 404)")
		} else {
			t.Fatalf("Unexpected error: %v", err)
		}
	} else if result != nil {
		t.Fatalf("Expected nil result for non-existent resource")
	}
}

func TestClientAuthentication(t *testing.T) {
	// Skip if not running integration tests
	if os.Getenv("INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test. Set INTEGRATION_TESTS=true to run.")
	}

	// Get Aidbox configuration from environment
	url := os.Getenv("AIDBOX_URL")
	if url == "" {
		t.Fatal("AIDBOX_URL environment variable is not set")
	}

	clientID := os.Getenv("AIDBOX_CLIENT_ID")
	if clientID == "" {
		t.Fatal("AIDBOX_CLIENT_ID environment variable is not set")
	}

	clientSecret := os.Getenv("AIDBOX_CLIENT_SECRET")
	if clientSecret == "" {
		t.Fatal("AIDBOX_CLIENT_SECRET environment variable is not set")
	}

	testCases := []struct {
		name        string
		config      *ClientConfig
		expectError bool
	}{
		{
			name: "Valid client credentials",
			config: &ClientConfig{
				URL:          url,
				ClientID:     clientID,
				ClientSecret: clientSecret,
			},
			expectError: false,
		},
		{
			name: "Invalid client credentials",
			config: &ClientConfig{
				URL:          url,
				ClientID:     "invalid",
				ClientSecret: "invalid",
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := NewClient(tc.config)
			if tc.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if client == nil {
					t.Error("Client is nil but no error was returned")
				}
			}
		})
	}
}
