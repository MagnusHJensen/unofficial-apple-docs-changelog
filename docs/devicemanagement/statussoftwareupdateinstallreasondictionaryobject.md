# StatusSoftwareUpdateInstallReasonDictionaryObject

A status report that contains details about the reason for a pending software update.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, tvOS 18.4, visionOS 26.0

## Properties

### declaration-id

- **Type:** `string`
- **Required:** No

The identifier of the declaration that caused the software update to occur. This key is present only if the `reason` array contains the `declaration` value.

### reason

- **Type:** `[string]`
- **Required:** Yes
- **Allowed Values:** `system-settings`, `install-tonight`, `auto-update`, `notification`, `setup-assistant`, `command-line`, `mdm`, `declaration`

A list of reasons for the pending software update. An empty list indicates that no software update is pending.

