# VPN.AlwaysOn.ApplicationExceptionElement

The dictionary that defines which applications can have traffic outside the VPN tunnel.

**Platforms:** iOS 13.6, iPadOS 13.6, Mac Catalyst 13.6, visionOS 1.0, Device Assignment Services , VPP License Management 

## Properties

### BundleIdentifier

- **Type:** `string`
- **Required:** Yes

The app’s bundle identifier.

### LimitToProtocols

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `UDP`

Limit the exception to only the specified list of protocols, with support for `UDP` only.

