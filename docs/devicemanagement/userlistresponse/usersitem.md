# UserListResponse.UsersItem

A dictionary that contains information about an active account on a device.

**Platforms:** iOS 9.3, iPadOS 9.3, macOS 10.13

## Properties

### DataQuota

- **Type:** `integer`
- **Required:** Yes

If present, the user’s data quota in bytes. This isn’t present if the account doesn’t enforce a quota. This value is available in iOS 9.3 and later.

### DataUsed

- **Type:** `integer`
- **Required:** Yes

The amount of data, in bytes, that the user has used. This value is available in iOS 9.3 and later.

### FullName

- **Type:** `string`
- **Required:** Yes

The user’s full name. This value is available in macOS 10.13 and later.

### HasDataToSync

- **Type:** `boolean`
- **Required:** Yes

If `true`, the user has data to sync to the cloud. This value is available in iOS 9.3 and later.

### HasSecureToken

- **Type:** `boolean`
- **Required:** Yes

If `true`, the user currently has a secure token set. This value is available in macOS 11 and later.

### IsLoggedIn

- **Type:** `boolean`
- **Required:** Yes

If `true`, the user is currently logged in on the device. This value is available in iOS 9.3 and later, and macOS 10.13 and later.

### MobileAccount

- **Type:** `boolean`
- **Required:** Yes

If `true`, the account is a mobile account. This value is available in macOS 10.13 and later.

### UID

- **Type:** `integer`
- **Required:** Yes

The user’s unique identifier. This value is available in macOS 10.13 and later.

### UserGUID

- **Type:** `string`
- **Required:** Yes

The user’s `GeneratedUID`. This value is available in macOS 10.13 and later.

### UserName

- **Type:** `string`
- **Required:** Yes

The user name for the account. In macOS, this is the short name of the user account. This value is available in iOS 9.3 and later, and macOS 10.13 and later.

