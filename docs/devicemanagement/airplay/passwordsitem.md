# AirPlay.PasswordsItem

The dictionary that defines passwords for AirPlay destinations.

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, macOS 10.10, Device Assignment Services , VPP License Management 

## Properties

### DeviceID

- **Type:** `string`
- **Required:** No

The device ID of the AirPlay destination; used in macOS.

Deprecated in macOS 15 and later as tvOS 18 AirPlay destinations don’t support it; use `DeviceName` instead.

### DeviceName

- **Type:** `string`
- **Required:** No

The name of the AirPlay destination; used in iOS, and available in macOS 15 and later.

### Password

- **Type:** `string`
- **Required:** Yes

The password for the AirPlay destination.

