# AIMAccount

The payload that configures an AIM account on the device.

**Platforms:** macOS 10.7

## Properties

### AIMAccountDescription

- **Type:** `string`
- **Required:** No

The description of the account.

### AIMAuthentication

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `AIMAuthPassword`

The authentication method for the account.

### AIMHostName

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `slogin.oscar.aol.com`

The server address.

### AIMPassword

- **Type:** `string`
- **Required:** No

The user’s password.

### AIMPort

- **Type:** `integer`
- **Required:** No
- **Default:** `5190`

The connection port for the server.

### AIMUserName

- **Type:** `string`
- **Required:** No

The user’s login name.

### AIMUseSSL

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables SSL.

## Discussion

Specify `com.apple.AIM.account` as the payload type.

### Profile availability

