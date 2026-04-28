# StatusSoftwareUpdatePendingVersionDictionaryObject

A dictionary that contains details about a pending software update.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, macOS 14.0, tvOS 18.4, visionOS 26.0, Device Assignment Services , VPP License Management 

## Properties

### build-version

- **Type:** `string`
- **Required:** Yes

The build version of the pending software update, including any Background Security Improvement version. This string is empty if no update is pending.

### os-version

- **Type:** `string`
- **Required:** Yes

The OS version of the pending software update, including any Background Security Improvement version. This string is empty if no update is pending.

### target-local-date-time

- **Type:** `string`
- **Required:** No

The local date time value that indicates when the pending software update will be installed. This key is only present when the pending software update is being enforced.

