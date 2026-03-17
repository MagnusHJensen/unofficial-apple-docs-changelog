# ClientConfigRequest

The request for the client configuration.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### notificationTypes

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `ASSET_MANAGEMENT`, `USER_MANAGEMENT`, `USER_ASSOCIATED`, `ASSET_COUNT`

The complete set of notification types to which MDM subscribes.

### notificationUrl

- **Type:** `string`
- **Required:** No

The URL to which subscribed notifications POST. This URL should only include a host and path.

### mdmInfo

- **Type:** `MdmInfo`
- **Required:** No

This value is returned by the server on all subsequent responses, and MDM uses it to ensure that no other MDM manages the same organization.

### notificationAuthToken

- **Type:** `string`
- **Required:** No

The bearer token that the server provides in the Authorization header of notifications. This is a shared secret between you and the server to verify that incoming notifications are from Apple.

## Topics

### Objects and Data Types

- [MdmInfo](/documentation/devicemanagement/mdminfo) - Information about the MDM client.

