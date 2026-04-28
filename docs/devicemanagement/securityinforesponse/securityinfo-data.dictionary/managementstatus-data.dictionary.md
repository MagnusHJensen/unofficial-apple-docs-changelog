# SecurityInfoResponse.SecurityInfo.ManagementStatus

A dictionary that contains the status of the device’s MDM enrollment.

**Platforms:** iOS 13.0, iPadOS 13.0, Mac Catalyst 13.0, macOS 10.13.2, tvOS 13.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### EnrolledViaDEP

- **Type:** `boolean`
- **Required:** No

If `true`, the device enrolled in MDM through the Device Enrollment Program (DEP). This value is available in macOS 10.13.2 and later.

### IsActivationLockManageable

- **Type:** `boolean`
- **Required:** No

If `true`, the type of enrollment allows the MDM to manage Activation Lock for this device. This value is available in macOS 10.15 and later.

### IsUserEnrollment

- **Type:** `boolean`
- **Required:** No

If `true`, the device is user-enrolled. This value is available in iOS 13 and later, and macOS 10.15 and later.

### UserApprovedEnrollment

- **Type:** `boolean`
- **Required:** No

If `true`, the enrollment was user-approved. If `false`, the device may reject certain security-sensitive payloads or commands. This value is available in macOS 10.13.2 and later.

