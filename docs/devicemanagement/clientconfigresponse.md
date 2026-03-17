# ClientConfigResponse

The response that contains the client configuration.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### countryISO2ACode

- **Type:** `string`
- **Required:** No

The ISO alpha-2 country code that designates the organization’s location.

### defaultPlatform

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `volumestore`, `enterprisestore`

The value that the MDM client passes for the platform parameter in the `contentMetadataLookup` request. For more information about how the MDM client queries metadata by using `contentMetadataLookup`, see [Getting App and Book Information (Legacy)](/documentation/devicemanagement/getting-app-and-book-information-legacy).

### notificationUrl

- **Type:** `string`
- **Required:** No

The current URL to post notifications to.

### subscribedNotificationTypes

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `ASSET_MANAGEMENT`, `USER_MANAGEMENT`, `USER_ASSOCIATED`, `ASSET_COUNT`

The set of currently subscribed notification types.

### websiteURL

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `https://school.apple.com`, `https://business.apple.com`

The current website URL for the specified platform.

### mdmInfo

- **Type:** `MdmInfo`
- **Required:** No

The current information for the provided token.

The response only includes this field when the MDM client sets a value using the [Client Config](/documentation/devicemanagement/client-config-4szk1) endpoint.

### notificationAuthToken

- **Type:** `string`
- **Required:** No

The current shared secret that the server returns in the Authorization header of notifications.

### tokenExpirationDate

- **Type:** `string`
- **Required:** No

The token’s expiration date in an ISO-8601 format.

Note: The server shows all dates and times in UTC.

### uId

- **Type:** `string`
- **Required:** No

The unique library identifier. When querying records using multiple tokens that may share libraries, use the `uId` field to filter duplicates and avoid double-counting records when different content managers upload duplicate tokens.

### locationName

- **Type:** `string`
- **Required:** No

The current name of the library.

## Discussion

Client config should be checked regularly in order to verify expected values for the various fields.

## Topics

### Content Metadata

- [Apps and Books for Organizations](/documentation/devicemanagement/apps-and-books-for-organizations) - Get details about apps and books to show to your users.

