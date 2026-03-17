# EditVppUserRequest

The request to edit a user.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### clientUserIdStr

- **Type:** `string`
- **Required:** No

The identifier supplied by the client when registering a user. Either `clientUserIdStr` or `userId` is required. If both `clientUserIdStr` and `userId` are supplied, `userId` takes precedence.

### email

- **Type:** `string`
- **Required:** No

The user’s email address. The `email` field is updated only if the value is provided in the request.

### itsIdHash

- **Type:** `string`
- **Required:** No

The hash of the user’s iTunes Store ID.

### managedAppleIDStr

- **Type:** `string`
- **Required:** No

The Apple Account associated with the user. This ID’s organization must match that of the provided sToken.

See [Associating an Apple Account with a Volume Purchase Program (VPP) User](/documentation/devicemanagement/associating-an-apple-id-with-a-volume-purchase-program-vpp-user) for more information.

### sToken

- **Type:** `string`
- **Required:** Yes

The authentication token. For more information, see [Authentication](/documentation/devicemanagement/managing-apps-and-books-through-web-services-legacy#Authentication).

### userId

- **Type:** `int64`
- **Required:** No

The unique identifier assigned by the VPP when registering the user. Either `clientUserIdStr` or `userId` is required. If both `clientUserIdStr` and `userId` are supplied, `userId` takes precedence.

