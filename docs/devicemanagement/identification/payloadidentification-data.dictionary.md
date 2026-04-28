# Identification.PayloadIdentification

The dictionary containing details about the user.

**Platforms:** macOS 10.7, Device Assignment Services , VPP License Management 

## Properties

### AuthMethod

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Password`, `UserEnteredPassword`

The authorization method. Either the profile contains the password or the user provides it.

### EmailAddress

- **Type:** `string`
- **Required:** Yes

The address for the account.

### FullName

- **Type:** `string`
- **Required:** Yes

The full name of the account.

### Password

- **Type:** `string`
- **Required:** Yes

The password for the account. Required when the `AuthMethod` is `Password`.

### Prompt

- **Type:** `string`
- **Required:** No

The custom instructions for the user, if needed.

### PromptMessage

- **Type:** `string`
- **Required:** No

The additional descriptive text for the user prompt.

### UserName

- **Type:** `string`
- **Required:** Yes

The UNIX user name for the accounts.

