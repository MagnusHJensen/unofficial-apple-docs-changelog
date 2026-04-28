# ExtensibleSingleSignOn.PlatformSSO

The dictionary to configure Platform SSO. Requires `Type` to be set to `Redirect`.

**Platforms:** macOS 14.0, Device Assignment Services , VPP License Management 

## Properties

### AccessKeyReaderGroupIdentifier

- **Type:** `data`
- **Required:** No

The reader group identifier for use with the `AccessKey`. The value needs to match the configured access key. Required if `NewUserAuthenticationMethods` contains `AccessKey`.

### AccessKeyReaderIssuerCertificateUUID

- **Type:** `string`
- **Required:** No

The `PayloadUUID` of a certificate payload for the issuer certificate of the `Terminal` identity of the access key. Other specifications refer to the key as the “Reader CA Public Key”. The key must be an elliptic curve key. Required if `NewUserAuthenticationMethods` includes `AccessKey`. The issuer of the Terminal identity of the access key needs to match this certificate, otherwise the device fails the authentication.

### AccessKeyTerminalIdentityUUID

- **Type:** `string`
- **Required:** No

The `PayloadUUID` of an identity payload to use as the `Terminal` identity of the access key. The identity needs to be trusted by the access key. Required if `NewUserAuthenticationMethods` includes `AccessKey`. Allowed identity payload types:

- `com.apple.security.pkcs12`
- `com.apple.security.acme`
- `com.apple.security.scep`

### AccountDisplayName

- **Type:** `string`
- **Required:** No

The display name for the account in notifications and authentication requests.

### AdditionalGroups

- **Type:** `[string]`
- **Required:** No

The list of created groups that don’t have administrator access.

### AdministratorGroups

- **Type:** `[string]`
- **Required:** No

The list of groups to use for administrator access. The system requests membership during authentication.

### AllowAccessKeyExpressMode

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system uses the access key in express mode, and doesn’t require authentication before use.

### AllowDeviceIdentifiersInAttestation

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system includes the device UDID and serial number in Platform SSO attestations.

### AuthenticationGracePeriod

- **Type:** `integer`
- **Required:** No

The amount of time after receiving or updating a `FileVaultPolicy`, `LoginPolicy`, or `UnlockPolicy` that the system can use unregistered local accounts. Required when `AllowAuthenticationGracePeriod` is set. Available in macOS 15 and later.

### AuthenticationMethod

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Password`, `UserSecureEnclaveKey`, `SmartCard`

The Platform SSO authentication method to use with the extension. Requires that the SSO Extension also support the method.

### AuthorizationGroups

- **Type:** `ExtensibleSingleSignOn.PlatformSSO.AuthorizationGroups`
- **Required:** No

The pairing of Authorization Rights to group names. When using this, the system updates the Authorization Right to use the group.

### EnableAuthorization

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

Enables using identity provider accounts at authorization prompts. Requires that `UseSharedDeviceKeys` is `true`. The system assigns groups using `AdministratorGroups`, `AdditionalGroups`, or `AuthorizationGroups`.

### EnableCreateFirstUserDuringSetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the device uses Platform SSO to create the first user account on the Mac during `Setup Assistant`.

### EnableCreateUserAtLogin

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

Enables creating users at the Login Window with an `AuthenticationMethod` of either `Password` or `SmartCard`. Requires that `UseSharedDeviceKeys` is `true`.

### EnableRegistrationDuringSetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the PlatformSSO registration process during Setup Assistant on devices running macOS 26 and later. Set this key to `true` when configuring PlatformSSO before enrollment using the `com.apple.psso.required` error response.

### FileVaultPolicy

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `AttemptAuthentication`, `RequireAuthentication`, `AllowOfflineGracePeriod`, `AllowAuthenticationGracePeriod`

The policy to apply when using Platform SSO at FileVault unlock on a Mac with Apple silicon. Applies when `AuthenticationMethod` is `Password`. Available in macOS 15 and later.

### LoginFrequency

- **Type:** `integer`
- **Required:** No
- **Default:** `64800`

The duration, in seconds, until the system requires a full login instead of a refresh. The default value is 64,800 (18 hours). The minimum value is 3600 (1 hour).

### LoginPolicy

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `AttemptAuthentication`, `RequireAuthentication`, `AllowOfflineGracePeriod`, `AllowAuthenticationGracePeriod`

The policy to apply when using Platform SSO at the Login Window. Applies when `AuthenticationMethod` is `Password`. Available in macOS 15 and later.

### NewUserAuthenticationMethods

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `Password`, `SmartCard`, `AccessKey`

The set of authentication methods to use for newly created accounts at login or during `Setup Assistant`. The system uses `Password` and `SmartCard` if this key isn’t present.

### NewUserAuthorizationMode

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Standard`, `Admin`, `Groups`, `Temporary`

The permission to apply to newly created accounts at login. Allowed values:

- `Standard`: The account is a standard user.
- `Admin`: The system adds the account to the local administrators group.
- `Groups`: The system assigns groups to the account using `AdministratorGroups`, `AdditionalGroups`, or `AuthorizationGroups`.
- `Temporary`: The system uses a temporary session configuration for newly created accounts at login.

### NonPlatformSSOAccounts

- **Type:** `[string]`
- **Required:** No

The list of local accounts that aren’t subject to the `FileVaultPolicy`, `LoginPolicy`, or `UnlockPolicy`. The accounts don’t receive a prompt to register for Platform SSO. Available in macOS 15 and later.

### OfflineGracePeriod

- **Type:** `integer`
- **Required:** No

The amount of time after the last successful Platform SSO login for using a local account password offline. Required when setting `AllowOfflineGracePeriod`. Available in macOS 15 and later.

### SynchronizeProfilePicture

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system requests the user’s profile picture from the SSO extension.

### TemporarySessionQuickLogin

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system uses a quicker Authenticated Guest Mode login to Mac behavior. The system erases user data from only select locations in the user home directory after each session completes. Once every eight hours the system erases the full user home directory after a session completes. Turn this on for shared environments that have a high frequency of short sessions.

### TokenToUserMapping

- **Type:** `ExtensibleSingleSignOn.PlatformSSO.TokenToUserMapping`
- **Required:** No

The attribute mapping to use when creating users, or for authorization.

### UnlockPolicy

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `AttemptAuthentication`, `RequireAuthentication`, `AllowOfflineGracePeriod`, `AllowAuthenticationGracePeriod`, `AllowTouchIDOrWatchForUnlock`

The policy to apply when using Platform SSO at screensaver unlock. Applies when `AuthenticationMethod` is `Password`. Available in macOS 15 and later.

### UserAuthorizationMode

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Standard`, `Admin`, `Groups`

The permission to apply to an account each time the user authenticates. Allowed values:

### UseSharedDeviceKeys

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system uses the same signing and encryption keys for all users. Only supported on the device channel.

## Topics

### Objects

- [ExtensibleSingleSignOn.PlatformSSO.AuthorizationGroups](/documentation/devicemanagement/extensiblesinglesignon/platformsso-data.dictionary/authorizationgroups-data.dictionary) - The pairing of Authorization Rights to group names.
- [ExtensibleSingleSignOn.PlatformSSO.TokenToUserMapping](/documentation/devicemanagement/extensiblesinglesignon/platformsso-data.dictionary/tokentousermapping-data.dictionary) - The attribute mapping to use when creating users, or for authorization.

