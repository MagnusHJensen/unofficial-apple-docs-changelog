# UserListResponse.UsersItem

A dictionary that contains information about an active account on a device.

**Platforms:** iOS 9.3, iPadOS 9.3, Mac Catalyst 9.3, macOS 10.13

## Properties

### DataQuota

- **Type:** `integer`
- **Required:** Yes

If present, the user’s data quota in bytes. This isn’t present if the account doesn’t enforce a quota.

Available: iOS 9.3+ | iPadOS 9.3+

### DataUsed

- **Type:** `integer`
- **Required:** Yes

The amount of data, in bytes, that the user has used.

Available: iOS 9.3+ | iPadOS 9.3+

### FullName

- **Type:** `string`
- **Required:** Yes

The user’s full name.

Available: macOS 10.13+

### HasDataToSync

- **Type:** `boolean`
- **Required:** Yes

If `true`, the user has data to sync to the cloud.

Available: iOS 9.3+ | iPadOS 9.3+

### HasSecureToken

- **Type:** `boolean`
- **Required:** Yes

If `true`, the user currently has a secure token set.

Available: macOS 11+

### IsLoggedIn

- **Type:** `boolean`
- **Required:** Yes

If `true`, the user is currently logged in on the device.

### MobileAccount

- **Type:** `boolean`
- **Required:** Yes

If `true`, the account is a mobile account.

Available: macOS 10.13+

### UID

- **Type:** `integer`
- **Required:** Yes

The user’s unique identifier.

Available: macOS 10.13+

### UserGUID

- **Type:** `string`
- **Required:** Yes

The user’s `GeneratedUID`.

Available: macOS 10.13+

### UserName

- **Type:** `string`
- **Required:** Yes

The user name for the account. In macOS, this is the short name of the user account.

