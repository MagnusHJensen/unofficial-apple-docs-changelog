# PrivacyPreferencesPolicyControl.Services.Identity

A dictionary listing apps and the privacy policy to apply to them.

**Platforms:** macOS 10.14

## Properties

### AEReceiverCodeRequirement

- **Type:** `string`
- **Required:** No

The code requirement for the receiving binary. This code requirement is required for AppleEvents service; not valid for other services.

Deprecated: macOS 27+

### AEReceiverIdentifier

- **Type:** `string`
- **Required:** No

The identifier of the process receiving an AppleEvent sent by the Identifier process. This identifier is required for AppleEvents service; not valid for other services.

Deprecated: macOS 27+

### AEReceiverIdentifierType

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `bundleID`, `path`

The type of AEReceiverIdentifier value, either `bundleID` or `path`. This setting is required for AppleEvents service; not valid for other services.

Deprecated: macOS 27+

### Allowed

- **Type:** `boolean`
- **Required:** No

If `true`, access is granted; otherwise, the process doesn’t have access. The user isn’t prompted and can’t change this value.

> 

Deprecated: macOS 27+

### Authorization

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Allow`, `Deny`, `AllowStandardUserToSetSystemService`

The `Authorization` key is an optional replacement for the `Allowed` key, which has one of the following possible values:

- `Allow`: Equivalent to a `true` value for the `Allowed` key
- `Deny`: Equivalent to a `false` value for the `Allowed` key
- `AllowStandardUserToSetSystemService`: Allows a standard (non-admin) user to configure the permissions for the specified app in the Privacy preferences for services that otherwise require admin authorization; only valid for the `ListenEvent` and `ScreenCapture` services

> 

Available: macOS 11+
Deprecated: macOS 27+

### CodeRequirement

- **Type:** `string`
- **Required:** Yes

Obtained via the command `codesign -display -r -`.

Deprecated: macOS 27+

### Comment

- **Type:** `string`
- **Required:** No

Not used.

Deprecated: macOS 27+

### Identifier

- **Type:** `string`
- **Required:** Yes

The bundle ID or installation path of the binary.

> 

Deprecated: macOS 27+

### IdentifierType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `bundleID`, `path`

The type of identifier value. Application bundles must be identified by bundle ID. Nonbundled binaries must be identified by installation path. Helper tools embedded within an application bundle automatically inherit the permissions of their enclosing app bundle.

Deprecated: macOS 27+

### StaticCode

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, statically validate the code requirement. Used only if the process invalidates its dynamic code signature.

Deprecated: macOS 27+

