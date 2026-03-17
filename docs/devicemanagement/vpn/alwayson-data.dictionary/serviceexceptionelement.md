# VPN.AlwaysOn.ServiceExceptionElement

The dictionary that defines service exceptions.

**Platforms:** iOS 8.0, iPadOS 8.0, visionOS 1.0

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

`CellularServices` is available in iOS 11.3 and later; it exempts `VoLTE`, `IMS` and `MMS`. WiFiCalling is exempted in iOS 13.4 and later.

`DeviceCommunication` is available in iOS 17.4 and later; it exempts network traffic used for communicating with devices connected via USB or Wi-Fi.

