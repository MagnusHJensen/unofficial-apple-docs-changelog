# VppAssignmentRequest

The request for a list of assignments.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### adamIdStr

- **Type:** `string`
- **Required:** No

The unique identifier for a product in the iTunes Store.

### clientUserIdStr

- **Type:** `string`
- **Required:** No

If specified, returns only assignments assigned to the given client user ID.

### pageIndex

- **Type:** `int32`
- **Required:** No

The index of the page to lookup. To page through the assignemnts, use the `nextPageIndex` value returned in the previous [VppAssignmentsResponse](/documentation/devicemanagement/vppassignmentsresponse).

This must be used in combination with a `requestID`, also from the previous response.

### requestId

- **Type:** `string`
- **Required:** No

A unique ID that is used when making paginated requests.

### serialNumber

- **Type:** `string`
- **Required:** No

If specified, returns only assignments assigned to the given serial number.

### sToken

- **Type:** `string`
- **Required:** Yes

The authentication token.

