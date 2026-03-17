# GetVppAssetResponse

The response with the asset.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### assets

- **Type:** `[VppAsset]`
- **Required:** No

The list of assets managed by the provided `sToken`.

### clientContext

- **Type:** `string`
- **Required:** No

The value currently associated with the provided `sToken`. This field is only included in the response when a value has been set via the [Client Configuration](/documentation/devicemanagement/client-configuration) endpoint.

See [Protecting Your VPP Account](/documentation/devicemanagement/protecting-your-vpp-account) for more information.

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

The UNIX epoch timestamp, in milliseconds, when the account’s `sToken` or password expires (whichever is earlier).

### location

- **Type:** `VppLocation`
- **Required:** No

The location associated with the provided `sToken`. This field is only returned when a location token is used with an Apple School Manager account.

### status

- **Type:** `int32`
- **Required:** No

The status code for the response. Possible values are: `0` = Success. `-1` = Failure.

### totalCount

- **Type:** `int32`
- **Required:** No

The total number of assets that will be returned.

### uId

- **Type:** `string`
- **Required:** No

The unique library identifier. When querying records using multiple tokens that may share libraries, use the `uId` field to filter duplicates. In this way, you can avoid double counting records when duplicate tokens are uploaded by different content managers.

