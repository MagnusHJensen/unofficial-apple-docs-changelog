# StatusManagementDeclarationsDeclarationObject

A processed declaration for the client.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### active

- **Type:** `boolean`
- **Required:** Yes

If `true`, the declaration is active on the device.

### identifier

- **Type:** `string`
- **Required:** Yes

The `identifier` of the declaration this status report refers to.

### reasons

- **Type:** `[StatusManagementDeclarationsStatusReasonObject]`
- **Required:** No

The details of any client errors.

### server-token

- **Type:** `string`
- **Required:** Yes

The `ServerToken` of the declaration this status report refers to.

### valid

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `unknown`, `invalid`, `valid`

This string defines the validity of the declaration. If it’s `invalid`, the `reasons` property contains more details.

## Topics

### Objects

- [StatusManagementDeclarationsStatusReasonObject](/documentation/devicemanagement/statusmanagementdeclarationsstatusreasonobject) - The details of an error in a status report.

