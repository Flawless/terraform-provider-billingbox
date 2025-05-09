package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAccessPolicyResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing with auto-generated ID
			{
				Config: testAccAccessPolicyResourceConfig("test-role", "read", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "role_name", "test-role"),
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "engine", "matcho"),
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "resource_type", "AccessPolicy"),
					resource.TestCheckResourceAttrSet("billingbox_access_policy.test", "id"),
				),
			},
			// Create and Read testing with custom ID
			{
				Config: testAccAccessPolicyResourceConfig("test-role-custom", "read", "custom-id-123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "id", "custom-id-123"),
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "role_name", "test-role-custom"),
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "engine", "matcho"),
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "resource_type", "AccessPolicy"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "billingbox_access_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccAccessPolicyResourceConfig("test-role-custom", "write", "custom-id-123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "id", "custom-id-123"),
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "role_name", "test-role-custom"),
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "engine", "matcho"),
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "resource_type", "AccessPolicy"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccAccessPolicyResourceConfig(roleName, action, customID string) string {
	idConfig := ""
	if customID != "" {
		idConfig = fmt.Sprintf(`  id        = %q`, customID)
	}

	return fmt.Sprintf(`
provider "billingbox" {
  url           = %[3]q
  client_id     = %[4]q
  client_secret = %[5]q
}

resource "billingbox_access_policy" "test" {
%[6]s
  role_name = %[1]q
  engine    = "matcho"
  matcho = {
    request-method = %[2]q == "read" ? "get" : "post"
    resource      = "Patient"
    action        = %[2]q
  }
}
`, roleName, action, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"), idConfig)
}
