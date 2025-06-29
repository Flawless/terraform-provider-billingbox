---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "billingbox_access_policy Resource - billingbox"
subcategory: ""
description: |-
  Access Policy resource
---

# billingbox_access_policy (Resource)

Access Policy resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `engine` (String) Engine of the Access Policy. Must be one of: json-schema, allow, sql, complex, matcho, clj, matcho-rpc, allow-rpc, signed-rpc
- `role_name` (String) Name of the Role

### Optional

- `and` (Dynamic) Array of engine rules that must all be satisfied (AND logic). Only used when engine is set to 'complex'. Cannot be used together with 'or'.
- `description` (String) Description of the Access Policy
- `id` (String) Unique identifier for the Access Policy. If not set, a random ID will be generated.
- `link` (Attributes List) List of links to Users, Clients, or Operations. If empty, the policy is considered global. (see [below for nested schema](#nestedatt--link))
- `matcho` (Dynamic) Match object of the Access Policy. Can contain nested maps with string, number, or boolean values. Only used when engine is set to 'matcho'.
- `or` (Dynamic) Array of engine rules where at least one must be satisfied (OR logic). Only used when engine is set to 'complex'. Cannot be used together with 'and'.
- `resource_type` (String) Resource type of the Access Policy. Always set to 'AccessPolicy'.
- `rpc` (Dynamic) RPC configuration object. Only used when engine is set to 'matcho-rpc' or 'allow-rpc'.
- `schema` (Dynamic) JSON Schema object for validation. Only used when engine is set to 'json-schema'.
- `sql` (Attributes) SQL configuration for the Access Policy. Only used when engine is set to 'sql'. (see [below for nested schema](#nestedatt--sql))

### Read-Only

- `meta` (Attributes) Server-side resource metadata (see [below for nested schema](#nestedatt--meta))

<a id="nestedatt--link"></a>
### Nested Schema for `link`

Required:

- `id` (String) ID of the resource to link to
- `resource_type` (String) Type of resource to link to. Must be one of: User, Client, Operation


<a id="nestedatt--sql"></a>
### Nested Schema for `sql`

Required:

- `query` (String) SQL query to execute. Should return a single row with one column containing a boolean result.


<a id="nestedatt--meta"></a>
### Nested Schema for `meta`

Read-Only:

- `created_at` (String) Time when the resource was created
- `last_updated` (String) Time when the resource was last updated
- `version_id` (String) Version identifier for the resource
