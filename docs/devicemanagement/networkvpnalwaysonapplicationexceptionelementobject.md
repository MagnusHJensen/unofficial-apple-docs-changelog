# NetworkVPNAlwaysOnApplicationExceptionElementObject

An array that contains an arbitrary number of apps whose connections occur outside the VPN.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), visionOS 27.0 (Beta)

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

