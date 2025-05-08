package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a new provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"billingbox": providerserver.NewProtocol6WithError(New("test")()),
}

func TestAccUserRoleAndAccessPolicy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create a user and role in a single step
			{
				Config: testAccUserAndRoleConfig("user1", "password1", "admin-role"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_user.admin", "name.given_name", "Camila"),
					resource.TestCheckResourceAttr("billingbox_user.admin", "name.family_name", "Harrington"),
					resource.TestCheckResourceAttr("billingbox_role.admin", "name", "admin-role"),
					resource.TestCheckResourceAttr("billingbox_role.admin", "user.id", "user1"),
				),
			},
			// Create access policies for different roles
			{
				Config: testAccAccessPoliciesConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check administrator policy
					resource.TestCheckResourceAttr("billingbox_access_policy.admin", "engine", "matcho"),
					resource.TestCheckResourceAttr("billingbox_access_policy.admin", "role_name", "admin-role"),
					// Check patient policy
					resource.TestCheckResourceAttr("billingbox_access_policy.patient", "engine", "matcho"),
					resource.TestCheckResourceAttr("billingbox_access_policy.patient", "role_name", "patient-role"),
				),
			},
		},
	})
}

func TestAccUserWithRole(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUserWithRoleConfig("user1", "password123", "practitioner"),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check user attributes
					resource.TestCheckResourceAttr("billingbox_user.test", "name.given_name", "Amy"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.family_name", "Shaw"),
					// Check role attributes
					resource.TestCheckResourceAttr("billingbox_role.test", "name", "practitioner"),
					resource.TestCheckResourceAttr("billingbox_role.test", "user.id", "user1"),
				),
			},
		},
	})
}

func testAccUserAndRoleConfig(userId, password, roleName string) string {
	return fmt.Sprintf(`
provider "billingbox" {
  url           = %[4]q
  client_id     = %[5]q
  client_secret = %[6]q
}

resource "billingbox_user" "admin" {
  id       = %[1]q
  password = %[2]q
  name = {
    given_name  = "Camila"
    family_name = "Harrington"
  }
}

resource "billingbox_role" "admin" {
  name = %[3]q
  user = {
    id = billingbox_user.admin.id
  }
}
`, userId, password, roleName, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"))
}

func testAccAccessPoliciesConfig() string {
	return fmt.Sprintf(`
provider "billingbox" {
  url           = %[1]q
  client_id     = %[2]q
  client_secret = %[3]q
}

# Administrator policy with full access
resource "billingbox_access_policy" "admin" {
  role_name = "admin-role"
  engine    = "matcho"
  matcho    = {
    request-method = {"$enum": ["get", "post", "put", "delete", "patch"]}
    user = {
      data = {
        roles = {"$contains": "Administrator"}
      }
    }
    client = {
      id = "postman"
    }
  }
}

# Patient policy with read-only access
resource "billingbox_access_policy" "patient" {
  role_name = "patient-role"
  engine    = "matcho"
  matcho    = {
    request-method = "get"
    user = {
      data = {
        roles = {"$contains": "Patient"}
      }
    }
    client = {
      id = "postman"
    }
  }
}
`, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"))
}

func testAccUserWithRoleConfig(userId, password, roleName string) string {
	return fmt.Sprintf(`
provider "billingbox" {
  url           = %[4]q
  client_id     = %[5]q
  client_secret = %[6]q
}

resource "billingbox_user" "test" {
  id       = %[1]q
  password = %[2]q
  name = {
    given_name  = "Amy"
    family_name = "Shaw"
  }
}

resource "billingbox_role" "test" {
  name = %[3]q
  user = {
    id = billingbox_user.test.id
  }
}
`, userId, password, roleName, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"))
}
