# StatusResponse

The status of an asynchronous event.

## Properties

### eventStatus

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `PENDING`, `COMPLETE`, `FAILED`

The current status of the asynchronous event.

### eventType

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `ASSOCIATE`, `DISASSOCIATE`, `REVOKE`, `CREATE`, `UPDATE`, `RETIRE`

The type of the asynchronous event.

### failures

- **Type:** `[ErrorResponse]`
- **Required:** No

The set of failures that occurred while the server processed the event.

### mdmInfo

- **Type:** `MdmInfo`
- **Required:** No

The client-specific information that the server stores for your device management service.

### numCompleted

- **Type:** `int32`
- **Required:** No

The number of tasks from the request that the server completed.

### numRequested

- **Type:** `int32`
- **Required:** No

The total number of tasks in the request.

### tokenExpirationDate

- **Type:** `string`
- **Required:** No

The token expiration date in an ISO-8601 format.

Note: The server shows all dates and times in UTC.

### uId

- **Type:** `string`
- **Required:** No

The unique library identifier. When querying records using multiple tokens that may share libraries, use the `uId` field to filter duplicates and avoid double-counting records when different content managers upload duplicate tokens.

## Discussion

Compare `numCompleted` against `numRequested` to track the progress of an event. When an event finishes with failures, `failures` describes each one.

## Topics

### Objects and Data Types

- [ErrorResponse](/documentation/devicemanagement/errorresponse) - The response that contains the error that occurs.
- [MdmInfo](/documentation/devicemanagement/mdminfo) - Information about the MDM client.

