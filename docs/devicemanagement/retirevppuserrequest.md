# RetireVppUserRequest

The request to retire a user.

**Platforms:** VPP License Management 1.0

## Properties

### clientUserIdStr

- **Type:** `string`
- **Required:** No

The identifier supplied by the client when registering a user. Either `clientUserIdStr` or `userId` is required. If both `clientUserIdStr` and `userId` are supplied, `userId` takes precedence.

### sToken

- **Type:** `string`
- **Required:** Yes

The authentication token. For more information, see [Authenticate with the web service](/documentation/devicemanagement/managing-apps-and-books-through-web-services-legacy#Authenticate-with-the-web-service).

### userId

- **Type:** `int64`
- **Required:** No

The unique identifier assigned by the VPP when registering the user. Either `clientUserIdStr` or `userId` is required. If both `clientUserIdStr` and `userId` are supplied, `userId` takes precedence.

