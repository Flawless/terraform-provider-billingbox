// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRoleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccRoleResourceConfig("test-user", "test-role"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_role.test", "user.id", "test-user"),
					resource.TestCheckResourceAttr("billingbox_role.test", "user.resource_type", "User"),
					resource.TestCheckResourceAttr("billingbox_role.test", "name", "test-role"),
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
				Config: testAccRoleResourceConfig("test-user", "updated-role"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_role.test", "user.id", "test-user"),
					resource.TestCheckResourceAttr("billingbox_role.test", "user.resource_type", "User"),
					resource.TestCheckResourceAttr("billingbox_role.test", "name", "updated-role"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccRoleResourceConfig(userID, roleName string) string {
	return fmt.Sprintf(`
provider "billingbox" {
  url           = %[3]q
  client_id     = %[4]q
  client_secret = %[5]q
}

resource "billingbox_role" "test" {
  name = %[1]q
  user = {
    id = %[2]q
  }
}
`, roleName, userID, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"))
}
