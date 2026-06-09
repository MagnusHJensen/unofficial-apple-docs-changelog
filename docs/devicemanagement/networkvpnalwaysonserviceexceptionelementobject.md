# NetworkVPNAlwaysOnServiceExceptionElementObject

An array that contains an arbitrary number of service exceptions.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### Action

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Allow`, `Drop`

The action to take with network connections from the named service.

### ServiceName

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `VoiceMail`, `AirPrint`, `CellularServices`, `DeviceCommunication`

The name of a service that’s exempt from Always On VPN.

`CellularServices` exempts `VoLTE`, `IMS`, `MMS`, and Wi-Fi calling.

`DeviceCommunication` exempts network traffic used for communicating with devices connected via USB or Wi-Fi.

