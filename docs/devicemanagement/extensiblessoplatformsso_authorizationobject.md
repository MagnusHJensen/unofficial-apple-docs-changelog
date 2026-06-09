# ExtensibleSSOPlatformSSO_AuthorizationObject

Settings for authorization prompts and group management.

**Platforms:** macOS 27.0 (Beta)

## Properties

### AdditionalGroups

- **Type:** `[string]`
- **Required:** No

The list of created groups that don’t have administrator access.

### AdministratorGroups

- **Type:** `[string]`
- **Required:** No

The list of groups to use for administrator access. The system requests membership during authentication.

### AuthorizationGroups

- **Type:** `ExtensibleSSOPlatformSSO_Authorization_AuthorizationGroupsObject`
- **Required:** No

The pairing of Authorization Rights to group names. When using this, the system updates the Authorization Right to use the group.

### EnableIdentityProviderAccounts

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

Enables using identity provider accounts at authorization prompts. Requires that `UseSharedDeviceKeys` is `true`. The system assigns groups using `AdministratorGroups`, `AdditionalGroups`, or `AuthorizationGroups`.

### UserAuthorizationMode

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Standard`, `Admin`, `Groups`

The permission to apply to an account each time the user authenticates. Allowed values:

- `Standard`: The account is a standard user.
- `Admin`: The system adds the account to the local administrators group.
- `Groups`: The system assigns group to the account using `AdministratorGroups`, `AdditionalGroups`, or `AuthorizationGroups`.

## Topics

### Objects

- [ExtensibleSSOPlatformSSO_Authorization_AuthorizationGroupsObject](/documentation/devicemanagement/extensiblessoplatformsso_authorization_authorizationgroupsobject) - The pairing of Authorization Rights to group names. When using this, the system updates the Authorization Right to use the group.

