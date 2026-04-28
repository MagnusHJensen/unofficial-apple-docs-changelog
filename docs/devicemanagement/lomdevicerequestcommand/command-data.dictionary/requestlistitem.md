# LOMDeviceRequestCommand.Command.RequestListItem

A dictionary that contains a requested action to perform on a device using lights-out management (LOM).

**Platforms:** macOS 11.0, Device Assignment Services , VPP License Management 

## Properties

### DeviceDNSName

- **Type:** `string`
- **Required:** Yes

The DNS name of the device. This should match the `dNSName` in [SCEP.PayloadContent.SubjectAltName](/documentation/devicemanagement/scep/payloadcontent-data.dictionary/subjectaltname-data.dictionary) or an equivalent in a PKCS12 identity.

### DeviceRequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `PowerON`, `PowerOFF`, `Reset`

The requested action to perform on the device.

### DeviceRequestUUID

- **Type:** `string`
- **Required:** Yes

The unique identifier of the request.

### LOMProtocolVersion

- **Type:** `integer`
- **Required:** Yes

The LOM protocol version that the device supports. Provide the same value that `LOMProtocolVersion` receives in the [LOMSetupRequestResponse](/documentation/devicemanagement/lomsetuprequestresponse).

### PrimaryIPv6AddressList

- **Type:** `[string]`
- **Required:** Yes

An array that contains the IPv6 addresses for primary LOM-compatible Ethernet interfaces for the device.

### SecondaryIPv6AddressList

- **Type:** `[string]`
- **Required:** Yes

An array that contains the IPv6 addresses for secondary LOM-compatible Ethernet interfaces for the device.

