package provider

import (
	// Add necessary imports here, e.g.:
	// "context"
	// "net/http"
	// "github.com/hashicorp/terraform-plugin-framework/resource"
	// "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	// "github.com/hashicorp/terraform-plugin-framework/types"
)

// UserResource implements the Terraform resource for User.
type UserResource struct {
	// Add client or other fields as needed
}

// TODO implement user with following schema taking example resource as refference, add proper test like in example, meta schema should go to the schema_metadata.go file since it would be shared between resources, also mark metadata version_id, created_at and last_updated attributes as computed and suppress diffs for them
// entry:
//   - resource:
//       description: >-
//         NB: this attr is ignored. A Boolean value indicating the User's
//         administrative status.
//       path:
//         - active
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           boolean
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.active
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.active
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.active
//   - resource:
//       description: >-
//         Used to indicate the User's default location for purposes of localizing
//         items such as currency, date time format, or numerical representations.
//       path:
//         - locale
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.locale
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.locale
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.locale
//   - resource:
//       path:
//         - twoFactor
//       module: auth
//       _source: code
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       description: Two factor settings for user
//       resourceType: Attribute
//       id: >-
//         User.twoFactor
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.twoFactor
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.twoFactor
//   - resource:
//       description: Primary phoneNumber
//       path:
//         - phoneNumber
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.phoneNumber
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.phoneNumber
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.phoneNumber
//   - resource:
//       description: >-
//         A label indicating the attribute's function, e.g., 'aim', 'gtalk',
//         'xmpp'.
//       path:
//         - ims
//         - type
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.ims.type
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.ims.type
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.ims.type
//   - resource:
//       description: The value of a role.
//       path:
//         - roles
//         - value
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.roles.value
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.roles.value
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.roles.value
//   - resource:
//       description: >-
//         A Boolean value indicating the 'primary' or preferred attribute value
//         for this attribute.  The primary attribute value 'true' MUST appear no
//         more than once.
//       path:
//         - x509Certificates
//         - primary
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           boolean
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.x509Certificates.primary
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.x509Certificates.primary
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.x509Certificates.primary
//   - resource:
//       description: Identifies the name of a cost center.
//       path:
//         - costCenter
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.costCenter
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.costCenter
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.costCenter
//   - resource:
//       description: >-
//         A Boolean value indicating the 'primary' or preferred attribute value
//         for this attribute, e.g., the preferred messenger or primary messenger.
//         The primary attribute value 'true' MUST appear no more than once.
//       path:
//         - ims
//         - primary
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           boolean
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.ims.primary
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.ims.primary
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.ims.primary
//   - resource:
//       description: A human-readable name, primarily used for display purposes.  READ-ONLY.
//       path:
//         - roles
//         - display
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.roles.display
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.roles.display
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.roles.display
//   - resource:
//       description: The value of an X.509 certificate.
//       path:
//         - x509Certificates
//         - value
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           base64Binary
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.x509Certificates.value
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.x509Certificates.value
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.x509Certificates.value
//   - resource:
//       description: >-
//         The name of the User, suitable for display to end-users.  The name
//         SHOULD be the full name of the User being described, if known.
//       path:
//         - displayName
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.displayName
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.displayName
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.displayName
//   - resource:
//       description: The city or locality component.
//       path:
//         - addresses
//         - locality
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.addresses.locality
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.addresses.locality
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.addresses.locality
//   - resource:
//       description: A Boolean value indicating the User's administrative status.
//       path:
//         - inactive
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           boolean
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.inactive
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.inactive
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.inactive
//   - resource:
//       description: >-
//         The User's time zone in the 'Olson' time zone database format, e.g.,
//         'America/Los_Angeles'.
//       path:
//         - timezone
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.timezone
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.timezone
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.timezone
//   - resource:
//       description: A label indicating the attribute's function.
//       path:
//         - x509Certificates
//         - type
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.x509Certificates.type
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.x509Certificates.type
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.x509Certificates.type
//   - resource:
//       description: A list of entitlements for the User that represent a thing the User has.
//       path:
//         - entitlements
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       isCollection: true
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.entitlements
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.entitlements
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.entitlements
//   - resource:
//       description: A label indicating the attribute's function, e.g., 'work' or 'home'.
//       path:
//         - addresses
//         - type
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.addresses.type
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.addresses.type
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.addresses.type
//   - resource:
//       description: >-
//         A Boolean value indicating the 'primary' or preferred attribute value
//         for this attribute, e.g., the preferred mailing address or primary email
//         address.  The primary attribute value 'true' MUST appear no more than
//         once.
//       path:
//         - emails
//         - primary
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           boolean
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.emails.primary
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.emails.primary
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.emails.primary
//   - resource:
//       description: A label indicating the attribute's function.
//       path:
//         - entitlements
//         - type
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.entitlements.type
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.entitlements.type
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.entitlements.type
//   - resource:
//       description: Identifies the name of a department.
//       path:
//         - department
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.department
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.department
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.department
//   - resource:
//       description: >-
//         A Boolean value indicating the 'primary' or preferred attribute value
//         for this attribute, e.g., the preferred photo or thumbnail.  The primary
//         attribute value 'true' MUST appear no more than once.
//       path:
//         - photos
//         - primary
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           boolean
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.photos.primary
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.photos.primary
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.photos.primary
//   - resource:
//       description: The zip code or postal code component.
//       path:
//         - addresses
//         - postalCode
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.addresses.postalCode
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.addresses.postalCode
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.addresses.postalCode
//   - resource:
//       description: >-
//         The User's manager.  A complex type that optionally allows service
//         providers to represent organizational hierarchy by referencing the 'id'
//         attribute of another User.
//       path:
//         - manager
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       refers:
//         - User
//       type:
//         id: >-
//           Reference
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.manager
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.manager
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.manager
//   - resource:
//       description: >-
//         A Boolean value indicating the 'primary' or preferred attribute value
//         for this attribute.  The primary attribute value 'true' MUST appear no
//         more than once.
//       path:
//         - roles
//         - primary
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           boolean
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.roles.primary
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.roles.primary
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.roles.primary
//   - resource:
//       path:
//         - data
//       isOpen: true
//       module: auth
//       _source: code
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       resourceType: Attribute
//       id: >-
//         User.data
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.data
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.data
//   - resource:
//       description: The country name component.
//       path:
//         - addresses
//         - country
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.addresses.country
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.addresses.country
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.addresses.country
//   - resource:
//       description: >-
//         The full name, including all middle names, titles, and suffixes as
//         appropriate, formatted for display (e.g., 'Ms. Barbara J Jensen, III').
//       path:
//         - name
//         - formatted
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.name.formatted
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.name.formatted
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.name.formatted
//   - resource:
//       description: A list of certificates issued to the User.
//       path:
//         - x509Certificates
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       isCollection: true
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.x509Certificates
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.x509Certificates
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.x509Certificates
//   - resource:
//       description: URL of a photo of the User.
//       path:
//         - photos
//         - value
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           uri
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.photos.value
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.photos.value
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.photos.value
//   - resource:
//       path:
//         - identifier
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       isCollection: true
//       type:
//         id: >-
//           Identifier
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.identifier
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.identifier
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.identifier
//   - resource:
//       description: >-
//         Email addresses for the user.  The value SHOULD be canonicalized by the
//         service provider, e.g., 'bjensen@example.com' instead of
//         'bjensen@EXAMPLE.COM'. Canonical type values of 'work', 'home', and
//         'other'.
//       path:
//         - emails
//         - value
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.emails.value
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.emails.value
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.emails.value
//   - resource:
//       description: >-
//         Used to identify the relationship between the organization and the
//         user.  Typical values used might be 'Contractor', 'Employee', 'Intern',
//         'Temp', 'External', and 'Unknown', but any value may be used.
//       path:
//         - userType
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.userType
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.userType
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.userType
//   - resource:
//       description: >-
//         A physical mailing address for this User. Canonical type values of
//         'work', 'home', and 'other'.  This attribute is a complex type with the
//         following sub-attributes.
//       path:
//         - addresses
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       isCollection: true
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.addresses
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.addresses
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.addresses
//   - resource:
//       path:
//         - link
//         - link
//       type:
//         id: >-
//           Reference
//         resourceType: Entity
//       module: auth
//       _source: code
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       resourceType: Attribute
//       id: >-
//         User.link.link
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.link.link
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.link.link
//   - resource:
//       description: >-
//         The family name of the User, or last name in most Western languages
//         (e.g., 'Jensen' given the full name 'Ms. Barbara J Jensen, III').
//       path:
//         - name
//         - familyName
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.name.familyName
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.name.familyName
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.name.familyName
//   - resource:
//       description: >-
//         A label indicating the attribute's function, e.g., 'work', 'home',
//         'mobile'.
//       path:
//         - phoneNumbers
//         - type
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.phoneNumbers.type
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.phoneNumbers.type
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.phoneNumbers.type
//   - resource:
//       description: Primary email
//       path:
//         - email
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           email
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.email
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.email
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.email
//   - resource:
//       description: >-
//         Phone numbers for the User.  The value SHOULD be canonicalized by the
//         service provider according to the format specified in RFC 3966, e.g.,
//         'tel:+1-201-555-0123'. Canonical type values of 'work', 'home',
//         'mobile', 'fax', 'pager', and 'other'.
//       path:
//         - phoneNumbers
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       isCollection: true
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.phoneNumbers
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.phoneNumbers
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.phoneNumbers
//   - resource:
//       path:
//         - fhirUser
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       refers:
//         - Patient
//         - Practitioner
//         - Person
//       type:
//         id: >-
//           Reference
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.fhirUser
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.fhirUser
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.fhirUser
//   - resource:
//       description: >-
//         Defined whether two-factor auth is enabled in current moment of time or
//         not
//       path:
//         - twoFactor
//         - enabled
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           boolean
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.twoFactor.enabled
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       isRequired: true
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.twoFactor.enabled
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.twoFactor.enabled
//   - resource:
//       description: A label indicating the attribute's function, e.g., 'work' or 'home'.
//       path:
//         - emails
//         - type
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.emails.type
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.emails.type
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.emails.type
//   - resource:
//       description: >-
//         A Boolean value indicating the 'primary' or preferred attribute value
//         for this attribute.  The primary attribute value 'true' MUST appear no
//         more than once.
//       path:
//         - entitlements
//         - primary
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           boolean
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.entitlements.primary
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.entitlements.primary
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.entitlements.primary
//   - resource:
//       description: The user's title, such as "Vice President."
//       path:
//         - title
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.title
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.title
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.title
//   - resource:
//       description: >-
//         A list of roles for the User that collectively represent who the User
//         is, e.g., 'Student', 'Faculty'.
//       path:
//         - roles
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       isCollection: true
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.roles
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.roles
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.roles
//   - resource:
//       description: >-
//         A fully qualified URL pointing to a page representing the User's online
//         profile.
//       path:
//         - profileUrl
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           uri
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.profileUrl
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.profileUrl
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.profileUrl
//   - resource:
//       description: >-
//         Email addresses for the user.  The value SHOULD be canonicalized by the
//         service provider, e.g., 'bjensen@example.com' instead of
//         'bjensen@EXAMPLE.COM'. Canonical type values of 'work', 'home', and
//         'other'.
//       path:
//         - emails
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       isCollection: true
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.emails
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.emails
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.emails
//   - resource:
//       description: Instant messaging address for the User.
//       path:
//         - ims
//         - value
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.ims.value
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.ims.value
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.ims.value
//   - resource:
//       description: The value of an entitlement.
//       path:
//         - entitlements
//         - value
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.entitlements.value
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.entitlements.value
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.entitlements.value
//   - resource:
//       description: A human-readable name, primarily used for display purposes.  READ-ONLY.
//       path:
//         - photos
//         - display
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.photos.display
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.photos.display
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.photos.display
//   - resource:
//       description: >-
//         The honorific suffix(es) of the User, or suffix in most Western
//         languages (e.g., 'III' given the full name 'Ms. Barbara J Jensen, III').
//       path:
//         - name
//         - honorificSuffix
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.name.honorificSuffix
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.name.honorificSuffix
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.name.honorificSuffix
//   - resource:
//       description: A label indicating the attribute's function.
//       path:
//         - roles
//         - type
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.roles.type
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.roles.type
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.roles.type
//   - resource:
//       description: A human-readable name, primarily used for display purposes.  READ-ONLY.
//       path:
//         - x509Certificates
//         - display
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.x509Certificates.display
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.x509Certificates.display
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.x509Certificates.display
//   - resource:
//       description: >-
//         A Boolean value indicating the 'primary' or preferred attribute value
//         for this attribute, e.g., the preferred phone number or primary phone
//         number.  The primary attribute value 'true' MUST appear no more than
//         once.
//       path:
//         - phoneNumbers
//         - primary
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           boolean
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.phoneNumbers.primary
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.phoneNumbers.primary
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.phoneNumbers.primary
//   - resource:
//       description: A human-readable name, primarily used for display purposes.  READ-ONLY.
//       path:
//         - ims
//         - display
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.ims.display
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.ims.display
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.ims.display
//   - resource:
//       description: A human-readable name, primarily used for display purposes.  READ-ONLY.
//       path:
//         - phoneNumbers
//         - display
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.phoneNumbers.display
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.phoneNumbers.display
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.phoneNumbers.display
//   - resource:
//       description: >-
//         A label indicating the attribute's function, i.e., 'photo' or
//         'thumbnail'.
//       path:
//         - photos
//         - type
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.photos.type
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.photos.type
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.photos.type
//   - resource:
//       description: Primary photo for user
//       path:
//         - photo
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           uri
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.photo
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.photo
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.photo
//   - resource:
//       description: Phone number of the User.
//       path:
//         - phoneNumbers
//         - value
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.phoneNumbers.value
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.phoneNumbers.value
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.phoneNumbers.value
//   - resource:
//       description: URLs of photos of the User.
//       path:
//         - photos
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       isCollection: true
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.photos
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.photos
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.photos
//   - resource:
//       path:
//         - link
//         - type
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       module: auth
//       _source: code
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       resourceType: Attribute
//       id: >-
//         User.link.type
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.link.type
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.link.type
//   - resource:
//       description: Code value
//       path:
//         - securityLabel
//         - code
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.securityLabel.code
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.securityLabel.code
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.securityLabel.code
//   - resource:
//       path:
//         - link
//       module: auth
//       _source: code
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       isCollection: true
//       resourceType: Attribute
//       id: >-
//         User.link
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.link
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.link
//   - resource:
//       description: Identifies the name of a division.
//       path:
//         - division
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.division
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.division
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.division
//   - resource:
//       description: Instant messaging addresses for the User.
//       path:
//         - ims
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       isCollection: true
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.ims
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.ims
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.ims
//   - resource:
//       description: >-
//         Numeric or alphanumeric identifier assigned to a person, typically based
//         on order of hire or association with an organization.
//       path:
//         - employeeNumber
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.employeeNumber
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.employeeNumber
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.employeeNumber
//   - resource:
//       description: List of security labes associated to the user
//       path:
//         - securityLabel
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       isCollection: true
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.securityLabel
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.securityLabel
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.securityLabel
//   - resource:
//       description: >-
//         The honorific prefix(es) of the User, or title in most Western languages
//         (e.g., 'Ms.' given the full name 'Ms. Barbara J Jensen, III').
//       path:
//         - name
//         - honorificPrefix
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.name.honorificPrefix
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.name.honorificPrefix
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.name.honorificPrefix
//   - resource:
//       path:
//         - gender
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       module: auth
//       _source: code
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       resourceType: Attribute
//       id: >-
//         User.gender
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.gender
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.gender
//   - resource:
//       description: >-
//         Transport of 2fa confirmation code. The lack of transport means Aidbox
//         do not need to send it over webhook
//       path:
//         - twoFactor
//         - transport
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.twoFactor.transport
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.twoFactor.transport
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.twoFactor.transport
//   - resource:
//       description: TOTP Secret key
//       path:
//         - twoFactor
//         - secretKey
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.twoFactor.secretKey
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       isRequired: true
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.twoFactor.secretKey
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.twoFactor.secretKey
//   - resource:
//       description: >-
//         The User's cleartext password.  This attribute is intended to be used as
//         a means to specify an initial password when creating a new User or to
//         reset an existing User's password.
//       path:
//         - password
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           password
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.password
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.password
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.password
//   - resource:
//       description: >-
//         The full street address component, which may include house number,
//         street name, P.O. box, and multi-line extended street address
//         information.  This attribute MAY contain newlines.
//       path:
//         - addresses
//         - streetAddress
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.addresses.streetAddress
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.addresses.streetAddress
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.addresses.streetAddress
//   - resource:
//       path:
//         - name
//       module: auth
//       _source: code
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       description: >-
//         The components of the user's real name. Providers MAY return just the
//         full name as a single string in the formatted sub-attribute, or they MAY
//         return just the individual component attributes using the other
//         sub-attributes, or they MAY return both.  If both variants are returned,
//         they SHOULD be describing the same name, with the formatted name
//         indicating how the component attributes should be combined.
//       resourceType: Attribute
//       id: >-
//         User.name
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.name
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.name
//   - resource:
//       description: A human-readable name, primarily used for display purposes.  READ-ONLY.
//       path:
//         - emails
//         - display
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.emails.display
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.emails.display
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.emails.display
//   - resource:
//       description: Code system
//       path:
//         - securityLabel
//         - system
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.securityLabel.system
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.securityLabel.system
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.securityLabel.system
//   - resource:
//       description: >-
//         The middle name(s) of the User (e.g., 'Jane' given the full name 'Ms.
//         Barbara J Jensen, III').
//       path:
//         - name
//         - middleName
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.name.middleName
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.name.middleName
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.name.middleName
//   - resource:
//       description: >-
//         Indicates the User's preferred written or spoken language.  Generally
//         used for selecting a localized user interface; e.g., 'en_US' specifies
//         the language English and country US.
//       path:
//         - preferredLanguage
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.preferredLanguage
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.preferredLanguage
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.preferredLanguage
//   - resource:
//       description: >-
//         The full mailing address, formatted for display or use with a mailing
//         label.  This attribute MAY contain newlines.
//       path:
//         - addresses
//         - formatted
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.addresses.formatted
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.addresses.formatted
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.addresses.formatted
//   - resource:
//       description: The state or region component.
//       path:
//         - addresses
//         - region
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.addresses.region
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.addresses.region
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.addresses.region
//   - resource:
//       description: A human-readable name, primarily used for display purposes.  READ-ONLY.
//       path:
//         - entitlements
//         - display
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.entitlements.display
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.entitlements.display
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.entitlements.display
//   - resource:
//       description: >-
//         Unique identifier for the User, typically used by the user to directly
//         authenticate to the service provider. Each User MUST include a non-empty
//         userName value.  This identifier MUST be unique across the service
//         provider's entire set of Users. REQUIRED.
//       path:
//         - userName
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.userName
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.userName
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.userName
//   - resource:
//       description: Identifies the name of an organization.
//       path:
//         - organization
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       refers:
//         - Organization
//       type:
//         id: >-
//           Reference
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.organization
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.organization
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.organization
//   - resource:
//       description: >-
//         The given name of the User, or first name in most Western languages
//         (e.g., 'Barbara' given the full name 'Ms. Barbara J Jensen, III').
//       path:
//         - name
//         - givenName
//       meta:
//         lastUpdated: '2025-03-03T10:25:17.351038Z'
//         createdAt: '2025-03-03T10:25:17.351038Z'
//         versionId: '0'
//       type:
//         id: >-
//           string
//         resourceType: Entity
//       resourceType: Attribute
//       module: auth
//       id: >-
//         User.name.givenName
//       resource:
//         id: >-
//           User
//         resourceType: Entity
//       _source: code
//     search:
//       mode: match
//     fullUrl: http://localhost:8080/fhir/Attribute/User.name.givenName
//     link:
//       - relation: self
//         url: http://localhost:8080/fhir/Attribute/User.name.givenName
