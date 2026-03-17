# SecurityInfoResponse.SecurityInfo.SecureBoot

The response object for the secure boot settings.

**Platforms:** macOS 10.15

## Properties

### ExternalBootLevel

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `allowed`, `disallowed`, `not supported`

The device’s external boot level, which indicates whether it allows booting from an external device, disallows it, or doesn’t support it.

### ReducedSecurity

- **Type:** `[SecurityInfoResponse.SecurityInfo.SecureBoot.ReducedSecurity]`
- **Required:** No

Reports which security features the user disables in `recoveryOS`. This property is only present for a Mac with Apple silicon when `SecureBootLevel` is `medium`.

Available in iOS 11 and later.

### SecureBootLevel

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `off`, `medium`, `full`, `not supported`

The security level for the bootable operating system versions.

## Topics

### Objects

- [SecurityInfoResponse.SecurityInfo.SecureBoot.ReducedSecurity](/documentation/devicemanagement/securityinforesponse/securityinfo-data.dictionary/secureboot-data.dictionary/reducedsecurity-data.dictionary) - Reports which security features the user disables in `recoveryOS`. This property is only present for a Mac with Apple silicon when `SecureBootLevel` is `medium`.

