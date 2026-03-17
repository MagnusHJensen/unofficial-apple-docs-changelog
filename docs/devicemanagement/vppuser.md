# VppUser

A user within the purchase program.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### clientUserIdStr

- **Type:** `string`
- **Required:** No

The identifier supplied by the client when registering a user. An identifier must be unique within the organization.

### email

- **Type:** `string`
- **Required:** No

The user’s email address.

### itsIdHash

- **Type:** `string`
- **Required:** No

The hash of the user’s iTunes Store ID. The `itsIdHash` field is included only when the user is associated with an iTunes Store account.

### licenses

- **Type:** `[VppLicense]`
- **Required:** No

The list of licenses which have been assigned to this user.

### status

- **Type:** `string`
- **Required:** No

The status of the user’s account in VPP. Possible values are:

- `Registered`: Indicates the user has been created, but is not yet associated with an iTunes account.
- `Associated`: Indicates that the user has been associated with an iTunes account. When a user is associated with an iTunes account, an `itsIdHash` is generated for the user.
- `Retired`: Indicates the user has been retired via  [Retire a User](/documentation/devicemanagement/retire-a-user).
- `Deleted`: Indicates that a VPP user is retired and its associated iTunes user has since been invited and associated with a new VPP user that shares the same `clientUserIdStr`. Because there are two VPP users with distinct `userId` values but the same `clientUserIdStr` value, the Deleted status is used to ensure database consistency.

A user with a `Deleted` status will never change status; its sole purpose is to ensure that your software can recognize that the userId is no longer associated with the `clientUserIdStr` record, and update any internal references appropriately.

### userId

- **Type:** `int64`
- **Required:** No

The unique identifier assigned by the VPP when registering the user.

