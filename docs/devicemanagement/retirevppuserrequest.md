# RetireVppUserRequest

The request to retire a user.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### clientUserIdStr

- **Type:** `string`
- **Required:** No

The identifier supplied by the client when registering a user. Either `clientUserIdStr` or `userId` is required. If both `clientUserIdStr` and `userId` are supplied, `userId` takes precedence.

### sToken

- **Type:** `string`
- **Required:** Yes

The authentication token. For more information, see [Authentication](/documentation/devicemanagement/managing-apps-and-books-through-web-services-legacy#Authentication).

### userId

- **Type:** `int64`
- **Required:** No

The unique identifier assigned by the VPP when registering the user. Either `clientUserIdStr` or `userId` is required. If both `clientUserIdStr` and `userId` are supplied, `userId` takes precedence.

