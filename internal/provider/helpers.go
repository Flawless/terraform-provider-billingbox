package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// getStringField safely extracts a string from a map and records diagnostics.
func getStringField(m map[string]interface{}, key string, diags *diag.Diagnostics, ctxMsg string) (string, bool) {
	raw, ok := m[key]
	if !ok {
		diags.AddError("Missing field", fmt.Sprintf("%s: %q not found in response", ctxMsg, key))
		return "", false
	}
	str, ok := raw.(string)
	if !ok {
		diags.AddError("Type error", fmt.Sprintf("%s: %q is not a string", ctxMsg, key))
		return "", false
	}
	return str, true
}
