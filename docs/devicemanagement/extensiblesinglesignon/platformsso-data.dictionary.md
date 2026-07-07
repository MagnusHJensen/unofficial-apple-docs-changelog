# ExtensibleSingleSignOn.PlatformSSO

The dictionary to configure Platform SSO. Requires setting `Type` to `Redirect`.

**Platforms:** macOS 14.0

## Properties

### AccessKeyReaderGroupIdentifier

- **Type:** `data`
- **Required:** No

The reader group identifier for use with the `AccessKey`. The value needs to match the configured access key. Required if `NewUserAuthenticationMethods` contains `AccessKey`.

Available: macOS 26+

### AccessKeyReaderIssuerCertificateUUID

- **Type:** `string`
- **Required:** No

The `PayloadUUID` of a certificate payload for the issuer certificate of the `Terminal` identity of the access key. Other specifications refer to the key as the “Reader CA Public Key”. The key must be an elliptic curve key. Required if `NewUserAuthenticationMethods` includes `AccessKey`. The issuer of the Terminal identity of the access key needs to match this certificate, otherwise the device fails the authentication.

Available: macOS 26.2+

### AccessKeyTerminalIdentityUUID

- **Type:** `string`
- **Required:** No

The `PayloadUUID` of an identity payload to use as the `Terminal` identity of the access key. The identity needs to be trusted by the access key. Required if `NewUserAuthenticationMethods` includes `AccessKey`. Allowed identity payload types:

- `com.apple.security.pkcs12`
- `com.apple.security.acme`
- `com.apple.security.scep`

Available: macOS 26+

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

Available: macOS 26+

### AllowDeviceIdentifiersInAttestation

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system includes the device UDID and serial number in Platform SSO attestations.

Available: macOS 15.4+

### AllowWebLoginPasswordSync

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system detects the password during web login and synchronizes it to the local account password for the user.

Available: macOS 27+

### AuthenticationGracePeriod

- **Type:** `integer`
- **Required:** No

The amount of time after receiving or updating a `FileVaultPolicy`, `LoginPolicy`, or `UnlockPolicy` that the system can use unregistered local accounts. Required when `AllowAuthenticationGracePeriod` is set.

Available: macOS 15+

### AuthenticationMethod

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Password`, `UserSecureEnclaveKey`, `SmartCard`, `OpenID`

The Platform SSO authentication method to use with the extension. Requires that the SSO Extension also support the method. `OpenID` is available in macOS 27 and later.

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

Available: macOS 26+

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

Available: macOS 26+

### FileVaultPolicy

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `AttemptAuthentication`, `RequireAuthentication`, `AllowOfflineGracePeriod`, `AllowAuthenticationGracePeriod`, `RequireTouchID`, `RequireTouchIDOrWatch`, `AllowOpenIDForTouchIDFallback`

The policy to apply when using Platform SSO at FileVault unlock on a Mac with Apple silicon. Applies when `AuthenticationMethod` is `Password`.

- `AttemptAuthentication`: The device attempts Platform SSO authentication before proceeding. If offline, unlock continues if the local account password matches. If online and the credential is incorrect, then the device requires a successful Platform SSO authentication, even if taken offline.
- `RequireAuthentication`: The device requires Platform SSO authentication before proceeding.  If the device is offline and `AllowOfflineGracePeriod` is enabled, then the device uses the offline `OfflineGracePeriod` to determine if the user can proceed or not. If online and the credential is incorrect, then the device requires a valid Platform SSO authentication to proceed, regardless of the `OfflineGracePeriod`. If the account isn’t registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the device uses `AuthenticationGracePeriod` to determine if the user can proceed or not.
- `AllowOfflineGracePeriod`: The device allows the use of the `OfflineGracePeriod` when `RequireAuthentication` is enabled.  If `AllowOfflineGracePeriod` isn’t set, then the device denies offline access.
- `AllowAuthenticationGracePeriod`: The device allows the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication` is enabled. The `AuthenticationGracePeriod` starts when any of the policies are updated. If `AllowAuthenticationGracePeriod` isn’t set, then the device denies unregistered account access.
- `RequireTouchID`: The device requires the use of Touch ID (and not Apple Watch) for File Vault unlock. Available in macOS 27 and later.
- `RequireTouchIDOrWatch`: The device requires the use of Touch ID or Apple Watch for File Vault unlock. Available in macOS 27 and later.
- `AllowOpenIDForTouchIDFallback`: The device allows web login as a fallback if touchID fails or isn’t available. Available in macOS 27 and later.

Available: macOS 15+

### LoginFrequency

- **Type:** `integer`
- **Required:** No
- **Default:** `64800`

The duration, in seconds, until the system requires a full login instead of a refresh. The default value is 64,800 (18 hours). The minimum value is 3600 (1 hour).

### LoginPolicy

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `AttemptAuthentication`, `RequireAuthentication`, `AllowOfflineGracePeriod`, `AllowAuthenticationGracePeriod`, `RequireTouchID`, `RequireTouchIDOrWatch`, `AllowOpenIDForTouchIDFallback`

The policy to apply when using Platform SSO at the Login Window. Applies when `AuthenticationMethod` is `Password`.

- `AttemptAuthentication`: The device attempts Platform SSO authentication before proceeding. If offline, login continues if the local account password matches. If online and the credential is incorrect, then the device requires a successful Platform SSO authentication to proceed, even if taken offline.
- `RequireAuthentication`: The device requires Platform SSO authentication before proceeding.  If the device is offline and `AllowOfflineGracePeriod` is enabled, then the device uses the offline `OfflineGracePeriod` to determine if the user can proceed or not. If online and the credential is incorrect, then the device requires a valid Platform SSO authentication to proceed, regardless of the `OfflineGracePeriod`. If the account isn’t registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the device uses the `AuthenticationGracePeriod` to determine if the user can proceed or not.
- `AllowOfflineGracePeriod`: The device allows the use of the `OfflineGracePeriod` when `RequireAuthentication` is enabled.  If `AllowOfflineGracePeriod` isn’t set, then the device denies offline access. Applies to web login and all offline passwords.
- `AllowAuthenticationGracePeriod`: The device allows the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication` is enabled. The `AuthenticationGracePeriod` starts when any of the policies have been updated. If `AllowAuthenticationGracePeriod` isn’t set, then the device denies unregistered account access.
- `RequireTouchID`: The device requires the use of Touch ID (and not Apple Watch) for login. Available in macOS 27 and later.
- `RequireTouchIDOrWatch`: The device requires the use of Touch ID or Apple Watch for login. Available in macOS 27 and later.
- `AllowOpenIDForTouchIDFallback`: The device allows web login as fallback if touchID fails or isn’t available. Available in macOS 27 and later.

Available: macOS 15+

### NewUserAuthenticationMethods

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `Password`, `SmartCard`, `AccessKey`, `OpenID`

The set of authentication methods to use for newly created accounts at login or during `Setup Assistant`. The system uses `Password` and `SmartCard` if this key isn’t present.

Available: macOS 26+

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

The list of local accounts that aren’t subject to the `FileVaultPolicy`, `LoginPolicy`, or `UnlockPolicy`. The accounts don’t receive a prompt to register for Platform SSO.

Available: macOS 15+

### OfflineGracePeriod

- **Type:** `integer`
- **Required:** No

The amount of time after the last successful Platform SSO login for using a local account password offline. Required when setting `AllowOfflineGracePeriod`.

Available: macOS 15+

### SynchronizeProfilePicture

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system requests the user’s profile picture from the SSO extension.

Available: macOS 26+

### TemporarySessionQuickLogin

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system uses a quicker Authenticated Guest Mode login to Mac behavior. The system erases user data from only select locations in the user home directory after each session completes. Once every eight hours the system erases the full user home directory after a session completes. Turn this on for shared environments that have a high frequency of short sessions.

Available: macOS 26+

### TokenToUserMapping

- **Type:** `ExtensibleSingleSignOn.PlatformSSO.TokenToUserMapping`
- **Required:** No

The attribute mapping to use when creating users, or for authorization.

### UnlockPolicy

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `AttemptAuthentication`, `RequireAuthentication`, `AllowOfflineGracePeriod`, `AllowAuthenticationGracePeriod`, `AllowTouchIDOrWatchForUnlock`, `RequireTouchID`, `RequireTouchIDOrWatch`, `AllowOpenIDForTouchIDFallback`

The policy to apply when using Platform SSO at screensaver unlock. Applies when `AuthenticationMethod` is `Password`.

- `AttemptAuthentication`: The device attempts Platform SSO authentication before proceeding. If offline, unlock will continue if the local account password matches. If online and the credential is incorrect, then the device requires a successful Platform SSO authentication to proceed, even if taken offline.
- `RequireAuthentication`: The device requires Platform SSO authentication before proceeding.  If the device is offline and `AllowOfflineGracePeriod` is enabled, then the device uses the offline `OfflineGracePeriod` to determine if the user can proceed or not. If online and the credential is incorrect, then the device requires a valid Platform SSO authentication to proceed regardless of the `OfflineGracePeriod`. If the account isn’t registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the device uses `AuthenticationGracePeriod` to determine if the user can proceed or not.
- `AllowOfflineGracePeriod`: The device allows the use of the `OfflineGracePeriod` when `RequireAuthentication` is enabled.  If `AllowOfflineGracePeriod` isn’t set, then the device denies offline access.
- `AllowAuthenticationGracePeriod`: The device allows the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication` is enabled. The `AuthenticationGracePeriod` starts when any of the policies have been updated. If `AllowAuthenticationGracePeriod` isn’t set, then the device denies the unregistered account access.
- `AllowTouchIDOrWatchForUnlock`: The device allows TouchID or Watch to unlock the screensaver instead of Platform SSO authentication when `RequireAuthentication` is enabled.
- `RequireTouchID`: The device requires the use of Touch ID (and not Apple Watch) for unlock. Available in macOS 27 and later.
- `RequireTouchIDOrWatch`: RThe device requires the use of Touch ID or Apple Watch for unlock. Available in macOS 27 and later.
- `AllowOpenIDForTouchIDFallback`: The device allows web login as fallback if touchID fails or isn’t available. Available in macOS 27 and later.

Available: macOS 15+

### UserAuthorizationMode

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Standard`, `Admin`, `Groups`

The permission to apply to an account each time the user authenticates. Allowed values:

- `Standard`: The account is a standard user.
- `Admin`: The system adds the account to the local administrators group.
- `Groups`: The system assigns group to the account using `AdministratorGroups`, `AdditionalGroups`, or `AuthorizationGroups`.

### UseSharedDeviceKeys

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system uses the same signing and encryption keys for all users. Only supported on the device channel.

### WebLoginURLAllowList

- **Type:** `[string]`
- **Required:** No

The set of allowed hosts that the system can load in the PSSO web view. Required if `AuthenticationMethod` is `OpenID`, or `NewUserAuthenticationMethods` contains `OpenID`.

Available: macOS 27+

## Topics

### Objects

- [ExtensibleSingleSignOn.PlatformSSO.AuthorizationGroups](/documentation/devicemanagement/extensiblesinglesignon/platformsso-data.dictionary/authorizationgroups-data.dictionary) - The pairing of Authorization Rights to group names.
- [ExtensibleSingleSignOn.PlatformSSO.TokenToUserMapping](/documentation/devicemanagement/extensiblesinglesignon/platformsso-data.dictionary/tokentousermapping-data.dictionary) - The attribute mapping to use when creating users, or for authorization.

