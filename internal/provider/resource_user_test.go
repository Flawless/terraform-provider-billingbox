package provider

import (
	"fmt"
	"os"
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
				Config: testAccUserResourceConfig("John", "Doe", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_user.test", "name.given_name", "John"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.family_name", "Doe"),
					resource.TestCheckResourceAttr("billingbox_user.test", "resource_type", "User"),
					resource.TestCheckResourceAttrSet("billingbox_user.test", "id"),
				),
			},
			// Create and Read testing with custom ID
			{
				Config: testAccUserResourceConfig("Jane", "Smith", "custom-id-123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_user.test", "id", "custom-id-123"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.given_name", "Jane"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.family_name", "Smith"),
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
				Config: testAccUserResourceConfig("Jane", "Johnson", "custom-id-123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_user.test", "id", "custom-id-123"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.given_name", "Jane"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.family_name", "Johnson"),
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

func testAccUserResourceConfig(givenName, familyName, customID string) string {
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
  password = "test-password"
}
`, givenName, familyName, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"), idConfig)
}
