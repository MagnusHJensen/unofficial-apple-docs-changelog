# ProfileListResponse.ProfileListItem.PayloadContentItem

A dictionary that describes a profile payload content item.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 9.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### PayloadDescription

- **Type:** `string`
- **Required:** No

A description of the payload.

### PayloadDisplayName

- **Type:** `string`
- **Required:** No

The human-readable name of the payload.

### PayloadIdentifier

- **Type:** `string`
- **Required:** Yes

The reverse-DNS-style identifier of the payload, such as `com.example.mypayload`.

### PayloadOrganization

- **Type:** `string`
- **Required:** No

The human-readable name of the organization that provided the payload.

### PayloadType

- **Type:** `string`
- **Required:** Yes

The type of payload, such as `com.apple.wifi.managed`.

### PayloadUUID

- **Type:** `string`
- **Required:** Yes

The unique identifier of the payload.

### PayloadVersion

- **Type:** `integer`
- **Required:** Yes
- **Allowed Values:** `1`

The version of the payload. The value is `1`.

