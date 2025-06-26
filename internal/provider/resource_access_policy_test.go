package provider

import (
	"fmt"
	"os"
	"terraform-provider-billingbox/internal/client"
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
					err = client.DeleteResource("AccessPolicy", "custom-id-123")
					if err != nil {
						t.Fatalf("Failed to delete access policy: %v", err)
					}
				},
				Config: testAccAccessPolicyResourceConfig("test-role-custom", "write", "custom-id-123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "id", "custom-id-123"),
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "role_name", "test-role-custom"),
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "engine", "matcho"),
					resource.TestCheckResourceAttr("billingbox_access_policy.test", "resource_type", "AccessPolicy"),
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

func TestAccAccessPolicyResource_SQL(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessPolicyResourceSQLConfig("sql-test-role"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_access_policy.sql_test", "role_name", "sql-test-role"),
					resource.TestCheckResourceAttr("billingbox_access_policy.sql_test", "engine", "sql"),
					resource.TestCheckResourceAttr("billingbox_access_policy.sql_test", "sql.query", "SELECT true FROM patient WHERE id = {{jwt.patient_id}} LIMIT 1;"),
				),
			},
		},
	})
}

func TestAccAccessPolicyResource_JSONSchema(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessPolicyResourceJSONSchemaConfig("json-schema-test-role"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_access_policy.json_schema_test", "role_name", "json-schema-test-role"),
					resource.TestCheckResourceAttr("billingbox_access_policy.json_schema_test", "engine", "json-schema"),
				),
			},
		},
	})
}

func TestAccAccessPolicyResource_Complex(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessPolicyResourceComplexConfig("complex-test-role"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_access_policy.complex_test", "role_name", "complex-test-role"),
					resource.TestCheckResourceAttr("billingbox_access_policy.complex_test", "engine", "complex"),
				),
			},
		},
	})
}

func TestAccAccessPolicyResource_Allow(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessPolicyResourceAllowConfig("allow-test-role"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_access_policy.allow_test", "role_name", "allow-test-role"),
					resource.TestCheckResourceAttr("billingbox_access_policy.allow_test", "engine", "allow"),
				),
			},
		},
	})
}

func TestAccAccessPolicyResource_RPC(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessPolicyResourceRPCConfig("rpc-test-role"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("billingbox_access_policy.rpc_test", "role_name", "rpc-test-role"),
					resource.TestCheckResourceAttr("billingbox_access_policy.rpc_test", "engine", "matcho-rpc"),
				),
			},
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

func testAccAccessPolicyResourceSQLConfig(roleName string) string {
	return fmt.Sprintf(`
provider "billingbox" {
  url           = %[2]q
  client_id     = %[3]q
  client_secret = %[4]q
}

resource "billingbox_access_policy" "sql_test" {
  role_name = %[1]q
  engine    = "sql"
  sql = {
    query = "SELECT true FROM patient WHERE id = {{jwt.patient_id}} LIMIT 1;"
  }
  description = "SQL engine test policy"
}
`, roleName, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"))
}

func testAccAccessPolicyResourceJSONSchemaConfig(roleName string) string {
	return fmt.Sprintf(`
provider "billingbox" {
  url           = %[2]q
  client_id     = %[3]q
  client_secret = %[4]q
}

resource "billingbox_access_policy" "json_schema_test" {
  role_name = %[1]q
  engine    = "json-schema"
  schema = {
    type = "object"
    required = ["user"]
    properties = {
      user = {
        type = "object"
        required = ["data"]
        properties = {
          data = {
            type = "object"
            required = ["practitioner_id"]
          }
        }
      }
    }
  }
  description = "JSON Schema engine test policy"
}
`, roleName, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"))
}

func testAccAccessPolicyResourceComplexConfig(roleName string) string {
	return fmt.Sprintf(`
provider "billingbox" {
  url           = %[2]q
  client_id     = %[3]q
  client_secret = %[4]q
}

resource "billingbox_access_policy" "complex_test" {
  role_name = %[1]q
  engine    = "complex"
  and = [
    {
      engine = "sql"
      sql = {
        query = "SELECT true"
      }
    },
    {
      engine = "complex"
      or = [
        {
          engine = "sql"
          sql = {
            query = "SELECT false"
          }
        },
        {
          engine = "sql"
          sql = {
            query = "SELECT true"
          }
        }
      ]
    }
  ]
  description = "Complex engine test policy with AND/OR logic"
}
`, roleName, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"))
}

func testAccAccessPolicyResourceAllowConfig(roleName string) string {
	return fmt.Sprintf(`
provider "billingbox" {
  url           = %[2]q
  client_id     = %[3]q
  client_secret = %[4]q
}

resource "billingbox_access_policy" "allow_test" {
  role_name = %[1]q
  engine    = "allow"
  description = "Allow engine test policy - grants unrestricted access"
}
`, roleName, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"))
}

func testAccAccessPolicyResourceRPCConfig(roleName string) string {
	return fmt.Sprintf(`
provider "billingbox" {
  url           = %[2]q
  client_id     = %[3]q
  client_secret = %[4]q
}

resource "billingbox_access_policy" "rpc_test" {
  role_name = %[1]q
  engine    = "matcho-rpc"
  description = "RPC engine test policy"
}
`, roleName, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"))
}
