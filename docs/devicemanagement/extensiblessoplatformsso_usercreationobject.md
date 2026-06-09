# ExtensibleSSOPlatformSSO_UserCreationObject

Settings for creating new users via Platform SSO.

**Platforms:** macOS 27.0 (Beta)

## Properties

### EnableAtLogin

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

Enables creating users at the Login Window with an `AuthenticationMethod` of either `Password` or `SmartCard`. Requires that `UseSharedDeviceKeys` is `true`.

### EnableFirstUserDuringSetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the device uses Platform SSO to create the first user account on the Mac during `Setup Assistant`.

### EnableRegistrationDuringSetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the PlatformSSO registration process during Setup Assistant on devices running macOS 26 and later. Set this key to `true` when configuring PlatformSSO before enrollment using the `com.apple.psso.required` error response.

### NewUserAuthenticationMethods

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `Password`, `SmartCard`, `AccessKey`, `OpenID`

The set of authentication methods to use for newly created accounts at login or during `Setup Assistant`. The system uses `Password` and `SmartCard` if this key isn’t present.

### NewUserAuthorizationMode

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Standard`, `Admin`, `Groups`, `Temporary`

The permission to apply to newly created accounts at login. Allowed values:

- `Standard`: The account is a standard user.
- `Admin`: The system adds the account to the local administrators group.
- `Groups`: The system assigns groups to the account using `Authorization.AdministratorGroups`, `Authorization.AdditionalGroups`, or `Authorization.AuthorizationGroups`.
- `Temporary`: The system uses a temporary session configuration for newly created accounts at login.

### TemporarySessionQuickLogin

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system uses a quicker Authenticated Guest Mode login to Mac behavior. The system erases user data from only select locations in the user home directory after each session completes. Once every eight hours the system erases the full user home directory after a session completes. Turn this on for shared environments that have a high frequency of short sessions.

### TokenToUserMapping

- **Type:** `ExtensibleSSOPlatformSSO_UserCreation_TokenToUserMappingObject`
- **Required:** No

The attribute mapping to use when creating users, or for authorization.

## Topics

### Objects

- [ExtensibleSSOPlatformSSO_UserCreation_TokenToUserMappingObject](/documentation/devicemanagement/extensiblessoplatformsso_usercreation_tokentousermappingobject) - The attribute mapping to use when creating users, or for authorization.

