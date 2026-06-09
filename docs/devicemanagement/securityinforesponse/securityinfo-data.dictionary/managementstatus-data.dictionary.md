# SecurityInfoResponse.SecurityInfo.ManagementStatus

A dictionary that contains the status of the device’s MDM enrollment.

**Platforms:** iOS 13.0, iPadOS 13.0, Mac Catalyst 13.0, macOS 10.13.2, tvOS 13.0, visionOS 1.1, watchOS 10.0

## Properties

### EnrolledViaDEP

- **Type:** `boolean`
- **Required:** No

If `true`, the device enrolled in MDM through Automated Device Enrollment (ADE).

Available: macOS 10.13.2+

### IsActivationLockManageable

- **Type:** `boolean`
- **Required:** No

If `true`, the type of enrollment allows the MDM to manage Activation Lock for this device.

Available: macOS 10.15+

### IsUserEnrollment

- **Type:** `boolean`
- **Required:** No

If `true`, the device is user-enrolled.

Available: iOS 13+ | iPadOS 13+ | macOS 10.15+ | tvOS 13+ | visionOS 1.1+ | watchOS 10+

### UserApprovedEnrollment

- **Type:** `boolean`
- **Required:** No

If `true`, the enrollment was user-approved. If `false`, the device may reject certain security-sensitive payloads or commands.

Available: macOS 10.13.2+

