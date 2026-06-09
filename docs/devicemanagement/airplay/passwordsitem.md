# AirPlay.PasswordsItem

The dictionary that defines passwords for AirPlay destinations.

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, macOS 10.10

## Properties

### DeviceID

- **Type:** `string`
- **Required:** No

The device ID of the AirPlay destination; used in macOS.

Deprecated in macOS 15 and later as tvOS 18 AirPlay destinations don’t support it; use `DeviceName` instead.

Available: macOS 10.10+
Deprecated: macOS 15+

### DeviceName

- **Type:** `string`
- **Required:** No

The name of the AirPlay destination.

Available: iOS 7+ | iPadOS 7+ | macOS 15+

### Password

- **Type:** `string`
- **Required:** Yes

The password for the AirPlay destination.

