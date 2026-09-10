# ExtensibleSSOPlatformSSO_PoliciesObject

Policies for login, unlock, and FileVault behavior.

**Platforms:** macOS 27.0

## Properties

### AuthenticationGracePeriod

- **Type:** `integer`
- **Required:** No

The amount of time (in seconds) after receiving or updating a `Policies.FileVault`, `Policies.Login`, or `Policies.Unlock` that the system can use unregistered local accounts. Required when `AllowAuthenticationGracePeriod` is set.

### FileVault

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `AttemptAuthentication`, `RequireAuthentication`, `AllowOfflineGracePeriod`, `AllowAuthenticationGracePeriod`, `RequireTouchID`, `RequireTouchIDOrWatch`, `AllowOpenIDForTouchIDFallback`

The policy to apply when using Platform SSO at FileVault unlock on a Mac with Apple silicon.

- `AttemptAuthentication`: The device attempts Platform SSO authentication before proceeding. If offline, unlock continues if the local account password matches. If online and the credential is incorrect, then the device requires a successful Platform SSO authentication is required, even if taken offline. Only use when `AuthenticationMethod` is `Password`.
- `RequireAuthentication`: The device requires Platform SSO authentication before proceeding. If the device is offline and `AllowOfflineGracePeriod` is enabled, then the device uses the offline `OfflineGracePeriod` to determine if the user can proceed or not. If online and the credential is incorrect, then the device requires a valid Platform SSO authentication to proceed, regardless of the `OfflineGracePeriod`. If the account isn’t registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the device uses `AuthenticationGracePeriod` to determine if the user can proceed or not. Only use when `AuthenticationMethod` is `Password`.
- `AllowOfflineGracePeriod`: The device allows the use of the `OfflineGracePeriod`. If `AllowOfflineGracePeriod` isn’t set, then the device denies offline access. Only use when `AuthenticationMethod` is `Password` and `RequireAuthentication` is enabled, or `AuthenticationMethod` is `OpenID`.
- `AllowAuthenticationGracePeriod`: The device allows the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication` is enabled. The `AuthenticationGracePeriod` starts when any of the policies are updated. If `AllowAuthenticationGracePeriod` isn’t set, then the device denies unregistered account access. Only use when `AuthenticationMethod` is `Password`.
- `RequireTouchID`: The device requires the use of Touch ID (and not Apple Watch) for FileVault unlock. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
- `RequireTouchIDOrWatch`: The device requires the use of Touch ID or Apple Watch for FileVault unlock. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
- `AllowOpenIDForTouchIDFallback`: The device allows web login as a fallback if Touch ID fails or isn’t available. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.

### Login

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `AttemptAuthentication`, `RequireAuthentication`, `AllowOfflineGracePeriod`, `AllowAuthenticationGracePeriod`, `RequireTouchID`, `RequireTouchIDOrWatch`, `AllowOpenIDForTouchIDFallback`

The policy to apply when using Platform SSO at the Login Window.

- `AttemptAuthentication`: The device attempts Platform SSO authentication before proceeding. If offline, login continues if the local account password matches. If online and the credential is incorrect, then the device requires a successful Platform SSO authentication to proceed, even if taken offline. Only use when `AuthenticationMethod` is `Password`.
- `RequireAuthentication`: The device requires Platform SSO authentication before proceeding. If the device is offline and `AllowOfflineGracePeriod` is enabled, then the device uses the offline `OfflineGracePeriod` to determine if the user can proceed or not. If online and the credential is incorrect, then the device requires a valid Platform SSO authentication to proceed, regardless of the `OfflineGracePeriod`. If the account isn’t registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the device uses the `AuthenticationGracePeriod` to determine if the user can proceed or not. Only use when `AuthenticationMethod` is `Password`.
- `AllowOfflineGracePeriod`: The device allows the use of the `OfflineGracePeriod`. If `AllowOfflineGracePeriod` isn’t set, then the device denies offline access. Only use when `AuthenticationMethod` is `Password` and `RequireAuthentication` is enabled, or `AuthenticationMethod` is `OpenID`.
- `AllowAuthenticationGracePeriod`: The device allows the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication` is enabled. The `AuthenticationGracePeriod` starts when any of the policies have been updated. If `AllowAuthenticationGracePeriod` isn’t set, then the device denies unregistered account access. Only use when `AuthenticationMethod` is `Password`.
- `RequireTouchID`: The device requires the use of Touch ID (and not Apple Watch) for login. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
- `RequireTouchIDOrWatch`: The device requires the use of Touch ID or Apple Watch for login. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
- `AllowOpenIDForTouchIDFallback`: The device allows web login as fallback if Touch ID fails or isn’t available. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.

### NonPlatformSSOAccounts

- **Type:** `[string]`
- **Required:** No

The list of local accounts that aren’t subject to the `Policies.FileVault`, `Policies.Login`, or `Policies.Unlock` policies. The accounts don’t receive a prompt to register for Platform SSO.

### OfflineGracePeriod

- **Type:** `integer`
- **Required:** No

The amount of time (in seconds) after the last successful Platform SSO login for using a local account password offline. Required when setting `AllowOfflineGracePeriod`.

### Unlock

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `AttemptAuthentication`, `RequireAuthentication`, `AllowOfflineGracePeriod`, `AllowAuthenticationGracePeriod`, `AllowTouchIDOrWatchForUnlock`, `RequireTouchID`, `RequireTouchIDOrWatch`, `AllowOpenIDForTouchIDFallback`

The policy to apply when using Platform SSO at screensaver unlock.

- `AttemptAuthentication`: The device attempts Platform SSO authentication before proceeding. If offline, unlock will continue if the local account password matches. If online and the credential is incorrect, then the device requires a successful Platform SSO authentication to proceed, even if taken offline. Only use when `AuthenticationMethod` is `Password`.
- `RequireAuthentication`: The device requires Platform SSO authentication before proceeding. If the device is offline and `AllowOfflineGracePeriod` is enabled, then the offline `OfflineGracePeriod` is used to determine if the user can proceed or not. If online and the credential is incorrect, then the device requires a valid Platform SSO authentication to proceed regardless of the `OfflineGracePeriod`. If the account isn’t registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the device uses `AuthenticationGracePeriod` to determine if the user can proceed or not. Only use when `AuthenticationMethod` is `Password`.
- `AllowOfflineGracePeriod`: The device allows the use of the `OfflineGracePeriod`. If `AllowOfflineGracePeriod` isn’t set, then the device denies offline access. Only use when `AuthenticationMethod` is `Password` and `RequireAuthentication` is enabled, or `AuthenticationMethod` is `OpenID`.
- `AllowAuthenticationGracePeriod`: The device allows the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication` is enabled. The `AuthenticationGracePeriod` starts when any of the policies have been updated. If `AllowAuthenticationGracePeriod` isn’t set, then the device denies the unregistered account access. Only use when `AuthenticationMethod` is `Password`.
- `AllowTouchIDOrWatchForUnlock`: The device allows Touch ID or Apple Watch to unlock the screensaver instead of Platform SSO authentication when `RequireAuthentication` is enabled. Only use when `AuthenticationMethod` is `Password`.
- `RequireTouchID`: The device requires the use of Touch ID (and not Apple Watch) for unlock. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
- `RequireTouchIDOrWatch`: The device requires the use of Touch ID or Apple Watch for unlock. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
- `AllowOpenIDForTouchIDFallback`: The device allows web login as fallback if Touch ID fails or isn’t available. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.

