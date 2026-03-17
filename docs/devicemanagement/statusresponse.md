# StatusResponse

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### eventStatus

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `PENDING`, `COMPLETE`, `FAILED`

### eventType

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `ASSOCIATE`, `DISASSOCIATE`, `REVOKE`, `CREATE`, `UPDATE`, `RETIRE`

### failures

- **Type:** `[ErrorResponse]`
- **Required:** No

### mdmInfo

- **Type:** `MdmInfo`
- **Required:** No

### numCompleted

- **Type:** `int32`
- **Required:** No

### numRequested

- **Type:** `int32`
- **Required:** No

### tokenExpirationDate

- **Type:** `string`
- **Required:** No

### uId

- **Type:** `string`
- **Required:** No

