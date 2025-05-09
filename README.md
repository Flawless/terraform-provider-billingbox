# Terraform Provider for BillingBox

This Terraform provider enables you to manage BillingBox resources using Terraform. It provides resources for managing users, roles, and access policies in your BillingBox instance.

## Features

- User Management: Create and manage users with customizable attributes
- Role Management: Define and assign roles to users
- Access Policy Management: Configure access policies with fine-grained permissions

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.23
- BillingBox instance with API access

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

To use the provider, you'll need to configure it with your BillingBox instance details. Here's an example configuration:

```hcl
provider "billingbox" {
  url           = "https://your-billingbox-instance.com"
  client_id     = "your-client-id"
  client_secret = "your-client-secret"
  username      = "optional-username"  # Optional: for password-grant authentication
  password      = "optional-password"  # Optional: for password-grant authentication
}
```

### Example: Creating a User with a Role

```hcl
resource "billingbox_user" "example" {
  name = {
    given_name  = "John"
    family_name = "Doe"
  }
  password = "secure-password"
}

resource "billingbox_role" "example" {
  name = "admin-role"
  user = {
    id = billingbox_user.example.id
  }
}

resource "billingbox_access_policy" "example" {
  role_name = "admin-role"
  engine    = "matcho"
  matcho = {
    request-method = {"$enum": ["get", "post", "put", "delete", "patch"]}
    user = {
      data = {
        roles = {"$contains": "Administrator"}
      }
    }
  }
}
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `make generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```

## Environment Variables for Testing

The following environment variables must be set for running acceptance tests:

- `AIDBOX_URL`: The URL of your BillingBox instance
- `AIDBOX_CLIENT_ID`: Your client ID
- `AIDBOX_CLIENT_SECRET`: Your client secret
