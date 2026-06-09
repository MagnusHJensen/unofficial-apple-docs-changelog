# GetVppLicensesResponse

The response with the licenses.

**Platforms:** VPP License Management 1.0

## Properties

### batchCount

- **Type:** `int32`
- **Required:** No

The number of licenses returned in the current batch.

See [Handle batched responses](/documentation/devicemanagement/retrieving-a-large-record-set#Handle-batched-responses) for more information.

### batchToken

- **Type:** `string`
- **Required:** No

Use this value in subsequent requests to get the next batch. The server generates the `batchToken` value, and it can be several kilobytes in size. For an initial request that doesn’t include `batchToken`, a `batchToken` value returns if the number of results exceeds a server-controlled limit. Subsequent requests must include `batchToken`. As long as additional batches remain, the server returns a new `batchToken` value in its response.

See [Handle batched responses](/documentation/devicemanagement/retrieving-a-large-record-set#Handle-batched-responses) for more information.

### clientContext

- **Type:** `string`
- **Required:** No

The value currently associated with the provided `sToken`. This field is only included in the response when a value is set with the [Client Configuration](/documentation/devicemanagement/client-configuration) endpoint.

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

### licenses

- **Type:** `[VppLicense]`
- **Required:** No

A list of licenses managed by the provided `sToken`.

### location

- **Type:** `VppLocation`
- **Required:** No

The location associated with the provided `sToken`. This field only returns when using a location token with an Apple School Manager account.

### sinceModifiedToken

- **Type:** `string`
- **Required:** No

A token that marks a place in time. After all licenses return for a request, the server includes a `sinceModifiedToken` value in the response. Use it in subsequent requests to get licenses modified after generating the token.

See [Track changes](/documentation/devicemanagement/retrieving-a-large-record-set#Track-changes) for more information.

### status

- **Type:** `int32`
- **Required:** No

The status code for the response. Possible values are:

`0` = Success.

`-1` = Failure.

### totalBatchCount

- **Type:** `int32`
- **Required:** No

The total number of round trips needed to get the complete result set.

See [Handle batched responses](/documentation/devicemanagement/retrieving-a-large-record-set#Handle-batched-responses) for more information.

### totalCount

- **Type:** `int32`
- **Required:** No

An estimate of the total number of licenses that will return. This value returns only for requests that don’t include `batchToken`, and only for the request that starts the batch process. The actual number of licenses that return can be different by the time the client finishes retrieving all licenses.

### uId

- **Type:** `string`
- **Required:** No

The unique library identifier. When querying records using multiple tokens that may share libraries, use the `uId` field to filter duplicates. In this way, you can avoid double-counting records when duplicate tokens are uploaded by different content managers.

