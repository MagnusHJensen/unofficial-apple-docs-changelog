# Identification.PayloadIdentification

The dictionary containing details about the user.

**Platforms:** macOS 10.7

## Properties

### AuthMethod

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Password`, `UserEnteredPassword`

The authorization method. Either the profile contains the password or the user provides it.

Deprecated: macOS 15.4+

### EmailAddress

- **Type:** `string`
- **Required:** Yes

The address for the account.

Deprecated: macOS 15.4+

### FullName

- **Type:** `string`
- **Required:** Yes

The full name of the account.

Deprecated: macOS 15.4+

### Password

- **Type:** `string`
- **Required:** Yes

The password for the account. Required when the `AuthMethod` is `Password`.

Deprecated: macOS 15.4+

### Prompt

- **Type:** `string`
- **Required:** No

The custom instructions for the user, if needed.

Deprecated: macOS 15.4+

### PromptMessage

- **Type:** `string`
- **Required:** No

The additional descriptive text for the user prompt.

Deprecated: macOS 15.4+

### UserName

- **Type:** `string`
- **Required:** Yes

The UNIX user name for the accounts.

Deprecated: macOS 15.4+

