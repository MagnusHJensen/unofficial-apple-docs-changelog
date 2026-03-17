# VppAssignmentsResponse

The response that contains a list of assignments.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### assignments

- **Type:** `[VppAssignment]`
- **Required:** No

An array of dictionaries representing the current assignments.

### assignmentsInCurrentPage

- **Type:** `int32`
- **Required:** No

The total number of assignments in the current page.

### clientContext

- **Type:** `string`
- **Required:** No

The value currently associated with the provided `sToken`. This field is only included in the response when a value has been set via the [Client Configuration](/documentation/devicemanagement/client-configuration) endpoint.

See [Protecting Your VPP Account](/documentation/devicemanagement/protecting-your-vpp-account) for more information.

### currentPageIndex

- **Type:** `int32`
- **Required:** No

The index of the page being returned.

### expirationMillis

- **Type:** `int64`
- **Required:** No

The UNIX epoch timestamp, in milliseconds, when the account’s `sToken` or password expires (whichever is earlier).

### location

- **Type:** `VppLocation`
- **Required:** No

The location associated with the provided `sToken`. This field is only returned when a location token is used with an Apple School Manager account.

### nextPageIndex

- **Type:** `int32`
- **Required:** No

The index of the next assignments page. This field will only be returned if there are additional assignments pages to read.

### requestId

- **Type:** `string`
- **Required:** No

The ID to be used for subsequent assignments page lookups. This field will only be returned if there are greater than 300 assignments.

### status

- **Type:** `int32`
- **Required:** No

The status code for the response. Possible values are:

`0` = Success. `-1` = Failure.

### totalAssignments

- **Type:** `int32`
- **Required:** No

The total number of assignments for the provided criteria.

### totalPages

- **Type:** `int32`
- **Required:** No

The total number of pages of assignments. There will be 300 assignments per page.

### getuId

- **Type:** `string`
- **Required:** No



