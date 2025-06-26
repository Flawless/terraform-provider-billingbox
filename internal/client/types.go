package client

type Resource struct {
	ResourceType string `json:"resourceType,omitempty"`
	ID           string `json:"id,omitempty"`
	Meta         *Meta  `json:"meta,omitempty"`
}

type Meta struct {
	VersionId   string `json:"versionId,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	LastUpdated string `json:"lastUpdated,omitempty"`
}

// UserName represents the name components of a user.
type UserName struct {
	GivenName       string `json:"givenName,omitempty"`
	MiddleName      string `json:"middleName,omitempty"`
	FamilyName      string `json:"familyName,omitempty"`
	HonorificPrefix string `json:"honorificPrefix,omitempty"`
}

// User represents a user resource.
type User struct {
	Resource
	Password string    `json:"password,omitempty"`
	Name     *UserName `json:"name,omitempty"`
	Email    string    `json:"email,omitempty"`
}

// RoleUser represents a user reference in a role.
type RoleUser struct {
	ResourceType string `json:"resourceType"`
	ID           string `json:"id"`
}

// Role represents a role resource.
type Role struct {
	Resource
	User *RoleUser `json:"user,omitempty"`
	Name string    `json:"name,omitempty"`
}

// AccessPolicy represents an access policy resource.
type AccessPolicy struct {
	Resource
	RoleName    string             `json:"roleName,omitempty"`
	Engine      string             `json:"engine,omitempty"`
	Matcho      interface{}        `json:"matcho,omitempty"`
	Description string             `json:"description,omitempty"`
	SQL         *SQLConfig         `json:"sql,omitempty"`
	Schema      interface{}        `json:"schema,omitempty"`
	And         []interface{}      `json:"and,omitempty"`
	Or          []interface{}      `json:"or,omitempty"`
	RPC         interface{}        `json:"rpc,omitempty"`
	Link        []AccessPolicyLink `json:"link,omitempty"`
}

// SQLConfig represents the SQL engine configuration.
type SQLConfig struct {
	Query string `json:"query"`
}

// AccessPolicyLink represents a link to User, Client, or Operation resources.
type AccessPolicyLink struct {
	ResourceType string `json:"resourceType"`
	ID           string `json:"id"`
}
