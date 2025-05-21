package provider

import (
	"fmt"
	"os"
	"terraform-provider-billingbox/internal/client"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRoleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing with auto-generated ID
			{
				Config: testAccRoleResourceConfig("test-role", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_role.test", "name", "test-role"),
					resource.TestCheckResourceAttr("billingbox_role.test", "resource_type", "Role"),
					resource.TestCheckResourceAttrSet("billingbox_role.test", "id"),
					resource.TestCheckResourceAttr("billingbox_role.test", "user.id", "test-user"),
				),
			},
			// Create and Read testing with custom ID
			{
				Config: testAccRoleResourceConfig("test-role-custom", "custom-id-123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_role.test", "id", "custom-id-123"),
					resource.TestCheckResourceAttr("billingbox_role.test", "name", "test-role-custom"),
					resource.TestCheckResourceAttr("billingbox_role.test", "resource_type", "Role"),
					resource.TestCheckResourceAttr("billingbox_role.test", "user.id", "test-user"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "billingbox_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccRoleResourceConfig("test-role-updated", "custom-id-123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_role.test", "id", "custom-id-123"),
					resource.TestCheckResourceAttr("billingbox_role.test", "name", "test-role-updated"),
					resource.TestCheckResourceAttr("billingbox_role.test", "resource_type", "Role"),
					resource.TestCheckResourceAttr("billingbox_role.test", "user.id", "test-user"),
				),
			},
			// Test not-found error handling
			{
				PreConfig: func() {
					client, err := client.NewClient(&client.ClientConfig{
						URL:          os.Getenv("AIDBOX_URL"),
						ClientID:     os.Getenv("AIDBOX_CLIENT_ID"),
						ClientSecret: os.Getenv("AIDBOX_CLIENT_SECRET"),
					})
					if err != nil {
						t.Fatalf("Failed to create client: %v", err)
					}
					err = client.DeleteResource("Role", "custom-id-123")
					if err != nil {
						t.Fatalf("Failed to delete role: %v", err)
					}
				},
				Config: testAccRoleResourceConfig("test-role-updated", ""), // Let Aidbox generate a new ID
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_role.test", "name", "test-role-updated"),
					resource.TestCheckResourceAttr("billingbox_role.test", "resource_type", "Role"),
					resource.TestCheckResourceAttrSet("billingbox_role.test", "id"),
					resource.TestCheckResourceAttr("billingbox_role.test", "user.id", "test-user"),
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

func testAccRoleResourceConfig(name, customID string) string {
	return fmt.Sprintf(`
provider "billingbox" {
  url           = %[1]q
  client_id     = %[2]q
  client_secret = %[3]q
}

resource "billingbox_role" "test" {
  id   = %[4]q
  name = %[5]q
  user = {
    id = "test-user"
  }
}
`, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"), customID, name)
}
