# JabberAccount

The payload that configures a Jabber account.

**Platforms:** macOS 10.7, Device Assignment Services , VPP License Management 

## Properties

### JabberAccountDescription

- **Type:** `string`
- **Required:** No

The description of the account.

### JabberAuthentication

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `JabberAuthPassword`

The authentication method for the account.

### JabberHostName

- **Type:** `string`
- **Required:** Yes

The server’s address.

### JabberPassword

- **Type:** `string`
- **Required:** No

The user’s password.

### JabberPort

- **Type:** `integer`
- **Required:** No
- **Default:** `5222`

The server’s port.

### JabberUserName

- **Type:** `string`
- **Required:** No

The user’s user name.

### JabberUseSSL

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables SSL.

## Discussion

Specify `com.apple.jabber.account` as the payload type.

### Profile availability

