# DeviceInformationResponse.QueryResponses.OrganizationInfo

The response dictionary that contains organization information.

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, macOS 10.11, tvOS 9.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### OrganizationAddress

- **Type:** `string`
- **Required:** No

The organization’s address. Use the LF character (`&#10`) to insert line breaks. This value is available in iOS 7 and later, macOS 10.11 and later, and tvOS 9 and later.

### OrganizationEmail

- **Type:** `string`
- **Required:** No

The organization’s support email address. This value is available in iOS 7 and later, macOS 10.11 and later, and tvOS 9 and later.

### OrganizationMagic

- **Type:** `string`
- **Required:** No

A unique identifier for the various services a single organization manages. This value is available in iOS 7 and later, macOS 10.11 and later, and tvOS 9 and later.

### OrganizationName

- **Type:** `string`
- **Required:** Yes

A string that describes the organization operating the MDM server. This value is available in iOS 7 and later, macOS 10.11 and later, and tvOS 9 and later.

### OrganizationPhone

- **Type:** `string`
- **Required:** No

The organization’s phone number. This value is available in iOS 7 and later, macOS 10.11 and later, and tvOS 9 and later.

