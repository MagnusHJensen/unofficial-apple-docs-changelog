# RegisterVppUserRequest

The request for registering a user.

**Platforms:** VPP License Management 1.0

## Properties

### clientUserIdStr

- **Type:** `string`
- **Required:** Yes

The identifier supplied by the client when registering a user. The identifier must be unique within the organization.

### email

- **Type:** `string`
- **Required:** No

The user’s email address.

### managedAppleIDStr

- **Type:** `string`
- **Required:** No

The Apple Account associated with the user. This ID’s organization must match the organization of the provided `sToken`.

### sToken

- **Type:** `string`
- **Required:** Yes

The authentication token. For more information, see [Authenticate with the web service](/documentation/devicemanagement/managing-apps-and-books-through-web-services-legacy#Authenticate-with-the-web-service).

## Discussion

### ClientUserIdStr

The `clientUserIdStr` must be unique within the organization and can’t be changed after a user is registered. Don’t use an email address because a future user might reuse an email address. You can use the GUID of the user.

When a user is first registered, the userʼs initial status is `Registered`.

If the user is already registered, as identified by `clientUserIdStr`, the following occurs:

- If the user’s status is `Registered` or `Associated`, that active user account returns.
- If the user’s status is `Retired` and the user was never assigned to an iTunes account, the account’s status changes to `Registered` and the existing user returns.
- If the user’s status is `Retired` and the user was previously assigned to an iTunes account, a new account is created.

It’s possible for more than one user record to exist for the same `clientUserIdStr` value — one for each iTunes account the `clientUserIdStr` value was associated with in the past (in addition to a currently active record or a retired and never-associated record). Each of these users has a unique `userId` value. Over time, with iTunes Store assignment, retirement, and reassignment, it’s possible for the active user’s `userId` value for a given `clientUserIdStr` to change.

Additionally, if two user identifiers exist for a given `clientUserIdStr`, one assigned to an iTunes account and the other unassigned, and a user accepts an invitation to be associated, it’s possible for the user to use the same iTunes account they used previously. If the user does so, the unassigned user record gets marked with the `Retired` status, and the formerly retired user record gets moved to the `Associated` status.

