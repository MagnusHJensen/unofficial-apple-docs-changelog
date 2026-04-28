# ReturnToServiceRequest

The return-to-service request details.

**Platforms:** iOS 26.0, iPadOS 26.0, Mac Catalyst 26.0, visionOS 26.0, Device Assignment Services , VPP License Management 

## Properties

### MessageType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ReturnToService`

The message type, which requires a value of `ReturnToService`.

### UDID

- **Type:** `string`
- **Required:** Yes

The device’s UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.

