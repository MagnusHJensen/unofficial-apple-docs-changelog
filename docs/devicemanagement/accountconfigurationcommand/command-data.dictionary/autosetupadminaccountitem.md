# AccountConfigurationCommand.Command.AutoSetupAdminAccountItem

A dictionary that describes the administrator account to create with Setup Assistant, which uses the first element and ignores additional elements.

**Platforms:** macOS 10.11

## Properties

### fullName

- **Type:** `string`
- **Required:** No

The full name of the user, which defaults to `shortName` if not specified.

### hidden

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, this sets the account attribute to make the account hidden in the Login Window and Users & Groups.

### passwordHash

- **Type:** `data`
- **Required:** No

Data that contains the pre-created salted PBKDF2 SHA512 password hash for the account.

### shortName

- **Type:** `string`
- **Required:** Yes

The short name of the user.

