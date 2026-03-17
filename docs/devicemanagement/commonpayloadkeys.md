# CommonPayloadKeys

The properties common to all payloads.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 9.0, visionOS 1.0, watchOS 1.0

## Properties

### PayloadDescription

- **Type:** `string`
- **Required:** No

The human-readable description of this payload. This description appears on the Detail screen.

### PayloadDisplayName

- **Type:** `string`
- **Required:** No

The human-readable name for the profile payload. The name appears on the Detail screen and doesn’t need to be unique.

### PayloadIdentifier

- **Type:** `string`
- **Required:** Yes

The reverse-DNS-style identifier for the payload. This identifier is usually the same as the [TopLevel](/documentation/devicemanagement/toplevel) value, with an additional appended component. This string must be unique within the profile.

During a profile replacement, the system updates payloads with the same `PayloadIdentifier` and `PayloadUUID` in the old and new profiles.

### PayloadOrganization

- **Type:** `string`
- **Required:** No

The human-readable string containing the name of the organization that provides the profile. This value doesn’t need to match the organization payload value in the enclosing dictionary.

### PayloadType

- **Type:** `string`
- **Required:** Yes

The payload type, which each payload domain’s reference page specifies.

### PayloadUUID

- **Type:** `string`
- **Required:** Yes

The globally unique identifier for the payload. The actual content is unimportant, but must be globally unique. In macOS, use `uuidgen` to generate UUIDs.

During a profile replacement, the system updates payloads with the same `PayloadIdentifier` and `PayloadUUID` in the old and new profiles.

### PayloadVersion

- **Type:** `integer`
- **Required:** Yes
- **Allowed Values:** `1`

The version of this specific payload.

## Discussion

### Profile availability

