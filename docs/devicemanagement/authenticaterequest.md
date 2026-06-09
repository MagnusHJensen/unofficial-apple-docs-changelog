# AuthenticateRequest

The authenticate request details.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Properties

### BuildVersion

- **Type:** `string`
- **Required:** No

The device’s build version.

Available: iOS 9+ | iPadOS 9+ | macOS 10.7+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### DeviceName

- **Type:** `string`
- **Required:** Yes

The device’s name.

Available: macOS 10.7+

### EnrollmentID

- **Type:** `string`
- **Required:** No

The per-enrollment identifier for the device. The system requires this value if the enrollment type is a user enrollment.

Available: iOS 13+ | iPadOS 13+ | macOS 10.15+ | visionOS 1.1+

### IMEI

- **Type:** `string`
- **Required:** No

The device’s IMEI (International Mobile Equipment Identity).

Available: iOS 9+ | iPadOS 9+ | visionOS 1.1+ | watchOS 10+

### MEID

- **Type:** `string`
- **Required:** No

The device’s MEID (Mobile Equipment Identifier).

Available: iOS 9+ | iPadOS 9+ | visionOS 1.1+ | watchOS 10+

### MessageType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Authenticate`

The message type, which requires a value of `Authenticate`.

### Model

- **Type:** `string`
- **Required:** Yes

The device’s model.

Available: macOS 10.7+

### ModelName

- **Type:** `string`
- **Required:** Yes

The device’s model name.

Available: macOS 10.7+

### OSVersion

- **Type:** `string`
- **Required:** No

The device’s OS version.

Available: iOS 9+ | iPadOS 9+ | macOS 10.7+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### ProductName

- **Type:** `string`
- **Required:** No

The device’s product name (such as `iPhone17,2`).

Available: iOS 9+ | iPadOS 9+ | macOS 10.7+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### SerialNumber

- **Type:** `string`
- **Required:** No

The device’s serial number.

Available: iOS 9+ | iPadOS 9+ | macOS 10.7+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### Topic

- **Type:** `string`
- **Required:** Yes

The topic that the device subscribes to.

### UDID

- **Type:** `string`
- **Required:** No

The device’s UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.

