# GetAssignmentsResponse

The paginated response that contains requested assignments.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### assignments

- **Type:** `[Assignment]`
- **Required:** No

The list of  requested assignments.

### currentPageIndex

- **Type:** `int32`
- **Required:** No

The current page index of the paginated response.

### nextPageIndex

- **Type:** `int32`
- **Required:** No

The next page index in the paginated response.

The response only includes this field when  there is a next page.

### size

- **Type:** `int32`
- **Required:** No

The number of assignments on the current page.

### totalPages

- **Type:** `int32`
- **Required:** No

The total number of pages in the paginated response.

### versionId

- **Type:** `string`
- **Required:** No

The current version identifier. When traversing the paginated response, use `versionId` to identify when changes occur to underlying data.

When any writes occur to the underlying data in a fetch, `versionId` updates.

### mdmInfo

- **Type:** `MdmInfo`
- **Required:** No

The current information for the provided token.

The response only includes this field when MDM sets a value using the [Client Config](/documentation/devicemanagement/client-config-4szk1) endpoint.

### tokenExpirationDate

- **Type:** `string`
- **Required:** No

The token expiration date in an ISO-8601 format.

Note: The server shows all dates and times in UTC.

### uId

- **Type:** `string`
- **Required:** No

The unique library identifier. When querying records using multiple tokens that may share libraries, use the `uId` field to filter duplicates and avoid double-counting records when different content managers upload duplicate tokens.

## Topics

### Objects and Data Types

- [MdmInfo](/documentation/devicemanagement/mdminfo) - Information about the MDM client.

