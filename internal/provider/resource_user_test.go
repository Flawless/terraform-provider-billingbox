// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
			// Create and Read testing
			{
				Config: testAccUserResourceConfig("John", "Doe", "password123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_user.test", "name.given_name", "John"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.family_name", "Doe"),
					resource.TestCheckResourceAttr("billingbox_user.test", "resource_type", "User"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "billingbox_user.test",
				ImportState:       true,
				ImportStateVerify: true,
				// Password is sensitive and won't be imported
				ImportStateVerifyIgnore: []string{"password"},
			},
			// Update and Read testing
			{
				Config: testAccUserResourceConfig("John", "Smith", "newpassword123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_user.test", "name.given_name", "John"),
					resource.TestCheckResourceAttr("billingbox_user.test", "name.family_name", "Smith"),
					resource.TestCheckResourceAttr("billingbox_user.test", "resource_type", "User"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccUserResourceConfig(givenName, familyName, password string) string {
	return fmt.Sprintf(`
provider "billingbox" {
  url           = %[4]q
  client_id     = %[5]q
  client_secret = %[6]q
}

resource "billingbox_user" "test" {
  name = {
    given_name  = %[1]q
    family_name = %[2]q
  }
  password = %[3]q
}
`, givenName, familyName, password, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"))
}
