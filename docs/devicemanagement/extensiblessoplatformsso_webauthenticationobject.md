# ExtensibleSSOPlatformSSO_WebAuthenticationObject

Settings for web authentication behavior.

**Platforms:** macOS 27.0 (Beta)

## Properties

### AllowPasswordSync

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system detects the password during web authentication and synchronizes it to the local account password for the user.

### URLAllowList

- **Type:** `[string]`
- **Required:** No

The set of allowed hosts that the system can load in the PSSO web view. Required if `AuthenticationMethod` is set to `OpenID`, or `UserCreation.AuthenticationMethods` contains `OpenID`.

