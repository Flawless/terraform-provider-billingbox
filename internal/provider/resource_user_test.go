package provider

import (
	"fmt"
	"os"
	"terraform-provider-billingbox/internal/client"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccUserResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing with auto-generated ID
			{
				Config: testAccUserResourceConfig("John", "Doe", "", "john.doe@example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_user.test", "name.given_name", "John"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.family_name", "Doe"),
					resource.TestCheckResourceAttr("billingbox_user.test", "email", "john.doe@example.com"),
					resource.TestCheckResourceAttr("billingbox_user.test", "resource_type", "User"),
					resource.TestCheckResourceAttrSet("billingbox_user.test", "id"),
				),
			},
			// Create and Read testing with custom ID
			{
				Config: testAccUserResourceConfig("Jane", "Smith", "custom-id-123", "jane.smith@example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_user.test", "id", "custom-id-123"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.given_name", "Jane"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.family_name", "Smith"),
					resource.TestCheckResourceAttr("billingbox_user.test", "email", "jane.smith@example.com"),
					resource.TestCheckResourceAttr("billingbox_user.test", "resource_type", "User"),
				),
			},
			// ImportState testing
			{
				ResourceName:            "billingbox_user.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
			// Update and Read testing
			{
				Config: testAccUserResourceConfig("Jane", "Johnson", "custom-id-123", "jane.johnson@example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_user.test", "id", "custom-id-123"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.given_name", "Jane"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.family_name", "Johnson"),
					resource.TestCheckResourceAttr("billingbox_user.test", "email", "jane.johnson@example.com"),
					resource.TestCheckResourceAttr("billingbox_user.test", "resource_type", "User"),
				),
			},
			// Test not-found error handling
			{
				PreConfig: func() {
					// Create a client to delete the user directly
					client, err := client.NewClient(&client.ClientConfig{
						URL:          os.Getenv("AIDBOX_URL"),
						ClientID:     os.Getenv("AIDBOX_CLIENT_ID"),
						ClientSecret: os.Getenv("AIDBOX_CLIENT_SECRET"),
					})
					if err != nil {
						t.Fatalf("Failed to create client: %v", err)
					}
					// Delete the user directly through the API
					err = client.DeleteResource("User", "custom-id-123")
					if err != nil {
						t.Fatalf("Failed to delete user: %v", err)
					}
				},
				Config: testAccUserResourceConfig("Jane", "Johnson", "custom-id-123", "jane.johnson@example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_user.test", "id", "custom-id-123"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.given_name", "Jane"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.family_name", "Johnson"),
					resource.TestCheckResourceAttr("billingbox_user.test", "email", "jane.johnson@example.com"),
					resource.TestCheckResourceAttr("billingbox_user.test", "resource_type", "User"),
				),
			},
			// Destroy testing
			{
				Config:  testAccProviderConfig(),
				Destroy: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccUserResourceConfig(givenName, familyName, customID, email string) string {
	idConfig := ""
	if customID != "" {
		idConfig = fmt.Sprintf(`  id   = %q`, customID)
	}

	return fmt.Sprintf(`
provider "billingbox" {
  url           = %[3]q
  client_id     = %[4]q
  client_secret = %[5]q
}

resource "billingbox_user" "test" {
%[6]s
  name = {
    given_name  = %[1]q
    family_name = %[2]q
  }
  email    = %[7]q
  password = "test-password"
}
`, givenName, familyName, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"), idConfig, email)
}
