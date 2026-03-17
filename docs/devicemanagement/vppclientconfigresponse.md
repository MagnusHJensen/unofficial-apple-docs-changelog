# VppClientConfigResponse

The response that contains the client configuration.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### apnToken

- **Type:** `string`
- **Required:** No

The Apple Push Notification token to use for notifications.

### appleId

- **Type:** `string`
- **Required:** No

The AppleID associated with the provided sToken.

### clientContext

- **Type:** `string`
- **Required:** No

The value currently associated with the provided sToken. This field is only included in the response when a value has been set via the [Client Configuration](/documentation/devicemanagement/client-configuration) endpoint.

See [Protecting Your VPP Account](/documentation/devicemanagement/protecting-your-vpp-account) for more information.

### countryCode

- **Type:** `string`
- **Required:** No

The two-letter ISO 3166-1 code that designates the country where the VPP account is located. For example, `US` stands for United States, `CA` for Canada, `JP` for Japan, and so on.

### defaultPlatform

- **Type:** `string`
- **Required:** No

The value to be passed for the platform parameter in the `contentMetadataLookupUrl` request. Possible values are:

- `volumestore`: For apps in the educational store.
- `enterprisestore`: For apps in the enterprise store.

See [Getting App and Book Information (Legacy)](/documentation/devicemanagement/getting-app-and-book-information-legacy) for more information.

### email

- **Type:** `string`
- **Required:** No

The email address associated with the provided sToken.

### errorMessage

- **Type:** `string`
- **Required:** No

The human-readable explanation of the error.

### errorNumber

- **Type:** `int32`
- **Required:** No

The numeric code of the error.

### expirationMillis

- **Type:** `int64`
- **Required:** No

The UNIX epoch timestamp, in milliseconds, when the account’s sToken or password expires (whichever is earlier).

### location

- **Type:** `VppLocation`
- **Required:** No

The location associated with the provided sToken. This field is only returned when a location token is used with an Apple School Manager account.

### organizationId

- **Type:** `string`
- **Required:** No

The unique identifier assigned to an organization by the VPP.

### organizationIdHash

- **Type:** `string`
- **Required:** No

The hash of `organizationId`.

### status

- **Type:** `int32`
- **Required:** No

The status code for the response. Possible values are:

- `0` = Success.
- `-1` = Failure.

### uId

- **Type:** `string`
- **Required:** No

The unique library identifier. When querying records using multiple tokens that may share libraries, use the `uId` field to filter duplicates. In this way, you can avoid double-counting records when duplicate tokens are uploaded by different content managers.

