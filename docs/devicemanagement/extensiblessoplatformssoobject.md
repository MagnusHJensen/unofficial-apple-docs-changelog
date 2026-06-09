# ExtensibleSSOPlatformSSOObject

The dictionary to configure Platform SSO.

**Platforms:** macOS 27.0 (Beta)

## Properties

### AccessKey

- **Type:** `ExtensibleSSOPlatformSSO_AccessKeyObject`
- **Required:** No

Settings for Access Key authentication.

### Account

- **Type:** `ExtensibleSSOPlatformSSO_AccountObject`
- **Required:** No

Account display and profile settings.

### AllowDeviceIdentifiersInAttestation

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system includes the device UDID and serial number in Platform SSO attestations.

### AuthenticationMethod

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Password`, `UserSecureEnclaveKey`, `SmartCard`, `OpenID`

The Platform SSO authentication method to use with the extension. Requires that the SSO Extension also support the method.

### Authorization

- **Type:** `ExtensibleSSOPlatformSSO_AuthorizationObject`
- **Required:** No

Settings for authorization prompts and group management.

### LoginFrequency

- **Type:** `integer`
- **Required:** No
- **Default:** `64800`

The duration, in seconds, until the system requires a full login instead of a refresh. The default value is 64,800 (18 hours). The minimum value is 3600 (1 hour).

### Policies

- **Type:** `ExtensibleSSOPlatformSSO_PoliciesObject`
- **Required:** No

Policies for login, unlock, and FileVault behavior.

### RegistrationToken

- **Type:** `string`
- **Required:** No

The token this device uses for registration with Platform SSO. Use it for silent registration with the Identity Provider. Requires that `AuthenticationMethod` in `PlatformSSO` isn’t empty.

### UserCreation

- **Type:** `ExtensibleSSOPlatformSSO_UserCreationObject`
- **Required:** No

Settings for creating new users via Platform SSO.

### UseSharedDeviceKeys

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system uses the same signing and encryption keys for all users.

Allowed scopes: system

### WebAuthentication

- **Type:** `ExtensibleSSOPlatformSSO_WebAuthenticationObject`
- **Required:** No

Settings for web authentication behavior.

## Topics

### Objects

- [ExtensibleSSOPlatformSSO_AccessKeyObject](/documentation/devicemanagement/extensiblessoplatformsso_accesskeyobject) - Settings for Access Key authentication.
- [ExtensibleSSOPlatformSSO_AccountObject](/documentation/devicemanagement/extensiblessoplatformsso_accountobject) - Account display and profile settings.
- [ExtensibleSSOPlatformSSO_AuthorizationObject](/documentation/devicemanagement/extensiblessoplatformsso_authorizationobject) - Settings for authorization prompts and group management.
- [ExtensibleSSOPlatformSSO_PoliciesObject](/documentation/devicemanagement/extensiblessoplatformsso_policiesobject) - Policies for login, unlock, and FileVault behavior.
- [ExtensibleSSOPlatformSSO_UserCreationObject](/documentation/devicemanagement/extensiblessoplatformsso_usercreationobject) - Settings for creating new users via Platform SSO.
- [ExtensibleSSOPlatformSSO_WebAuthenticationObject](/documentation/devicemanagement/extensiblessoplatformsso_webauthenticationobject) - Settings for web authentication behavior.

