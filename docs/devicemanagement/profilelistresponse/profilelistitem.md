# ProfileListResponse.ProfileListItem

A dictionary that describes a profile list item.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 9.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### HasRemovalPasscode

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the profile has a passcode for removal.

### IsEncrypted

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, it’s an encrypted profile.

### IsManaged

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the current MDM service installed the profile. MDM doesn’t return this value for supervised devices, and can remove or replace all profiles on supervised devices.

### PayloadContent

- **Type:** `[ProfileListResponse.ProfileListItem.PayloadContentItem]`
- **Required:** No

An array of payload content items. This value isn’t present if `IsEncrypted` is `true`.

### PayloadDescription

- **Type:** `string`
- **Required:** No

The description of the profile.

### PayloadDisplayName

- **Type:** `string`
- **Required:** No

The human-readable name of the profile.

### PayloadIdentifier

- **Type:** `string`
- **Required:** Yes

The reverse-DNS-style identifier of the profile; for example, `com.example.myprofile`.

### PayloadOrganization

- **Type:** `string`
- **Required:** No

The human-readable name of the organization that provided the profile.

### PayloadRemovalDisallowed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the user can’t delete the profile unless it has a removal password and the user provides it. The framework ignores this field on unsupervised devices.

### PayloadUUID

- **Type:** `string`
- **Required:** Yes

The unique identifier for the profile.

### PayloadVersion

- **Type:** `integer`
- **Required:** No

The version of the configuration profile as a whole, not of the individual profiles within it. The value should be `1`.

### SignerCertificates

- **Type:** `[data]`
- **Required:** No

An array that contains the certificate for signing the profile, followed by any intermediate certificates, in DER-encoded X.509 format.

### Source

- **Type:** `string`
- **Required:** No

A string set to `Declarative Device Management` when the profile is managed by Declarative Device Management.

## Topics

### Objects

- [ProfileListResponse.ProfileListItem.PayloadContentItem](/documentation/devicemanagement/profilelistresponse/profilelistitem/payloadcontentitem) - A dictionary that describes a profile payload content item.

