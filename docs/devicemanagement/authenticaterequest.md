# AuthenticateRequest

The authenticate request details.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 10.2, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### BuildVersion

- **Type:** `string`
- **Required:** No

The device’s build version.

### DeviceName

- **Type:** `string`
- **Required:** Yes

The device’s name.

### EnrollmentID

- **Type:** `string`
- **Required:** No

The per-enrollment identifier for the device. The system requires this value if the enrollment type is a user enrollment.

Available in iOS 13 and later, macOS 10.15 and later, and visionOS 2 and later.

### IMEI

- **Type:** `string`
- **Required:** No

The device’s IMEI (International Mobile Equipment Identity).

### MEID

- **Type:** `string`
- **Required:** No

The device’s MEID (Mobile Equipment Identifier).

### MessageType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Authenticate`

The message type, which requires a value of `Authenticate`.

### Model

- **Type:** `string`
- **Required:** Yes

The device’s model.

### ModelName

- **Type:** `string`
- **Required:** Yes

The device’s model name.

### OSVersion

- **Type:** `string`
- **Required:** No

The device’s OS version.

### ProductName

- **Type:** `string`
- **Required:** No

The device’s product name (such as `iPhone17,2`).

### SerialNumber

- **Type:** `string`
- **Required:** No

The device’s serial number.

### Topic

- **Type:** `string`
- **Required:** Yes

The topic that the device subscribes to.

### UDID

- **Type:** `string`
- **Required:** No

The device’s UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.

