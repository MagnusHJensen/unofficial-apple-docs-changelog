# ResponseUser

The user in the organization.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### clientUserId

- **Type:** `string`
- **Required:** No

The unique identifier for a user in your organization.

### email

- **Type:** `string`
- **Required:** No

The user’s email address.

### idHash

- **Type:** `string`
- **Required:** No

The hash of the user’s unique store identifier. The `idHash` field is only present when the user has an associated Apple Account.

### inviteCode

- **Type:** `string`
- **Required:** No

The invitation code that associates an Apple Account to a user.

### status

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Registered`, `Associated`, `Retired`, `Deleted`

The current status of the user in the organization.

