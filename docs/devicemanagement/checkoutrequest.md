# CheckOutRequest

The check out request details.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Properties

### EnrollmentID

- **Type:** `string`
- **Required:** Yes

The per-enrollment identifier for the device. The system requires this value if the enrollment type is a user enrollment.

Available in iOS 13 and later, macOS 10.15 and later, and visionOS 2 and later.

### MessageType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `CheckOut`

The message type, which requires a value of `CheckOut`.

### Topic

- **Type:** `string`
- **Required:** Yes

The topic the device subscribes to.

### UDID

- **Type:** `string`
- **Required:** Yes

The device’s UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.

