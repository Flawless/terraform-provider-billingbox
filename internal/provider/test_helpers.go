package provider

import (
	"fmt"
	"os"
)

// testAccProviderConfig returns a provider configuration block for testing.
func testAccProviderConfig() string {
	return fmt.Sprintf(`
provider "billingbox" {
  url           = %q
  client_id     = %q
  client_secret = %q
}
`, os.Getenv("AIDBOX_URL"), os.Getenv("AIDBOX_CLIENT_ID"), os.Getenv("AIDBOX_CLIENT_SECRET"))
}
