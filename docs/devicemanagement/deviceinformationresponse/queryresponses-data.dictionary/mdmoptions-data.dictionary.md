# DeviceInformationResponse.QueryResponses.MDMOptions

The response dictionary that contains MDM options.

**Platforms:** iOS 7.0, iPadOS 7.0, macOS 11.0, tvOS 9.0, visionOS 1.1, watchOS 10.0

## Properties

### ActivationLockAllowedWhileSupervised

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, a supervised device registers itself with Activation Lock when the user enables Find My. Unsupervised devices ignore this value. This value is available in iOS 7 and later, macOS 11 and later, and tvOS 9 and later.

### BootstrapTokenAllowed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the server supports Bootstrap Token commands. This value is available in macOS 11 and later.

### PromptUserToAllowBootstrapTokenForAuthentication

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device can accept a Bootstrap Token from the MDM server instead of prompting for user authentication prior to installation. This only applies when `BootstrapTokenAllowedForAuthentication` is `true` in the [SecurityInfoResponse.SecurityInfo](/documentation/devicemanagement/securityinforesponse/securityinfo-data.dictionary) response. This value is available for a Mac with Apple silicon in macOS 11 and later.

