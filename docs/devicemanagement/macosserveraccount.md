# MacOSServerAccount

The payload that configures a macOS Server account.

**Platforms:** iOS 9.0, iPadOS 9.0

## Properties

### AccountDescription

- **Type:** `string`
- **Required:** No

The description of the account.

### ConfiguredAccounts

- **Type:** `[MacOSServerAccount.ConfiguredAccountsItem]`
- **Required:** Yes

An array of dictionaries containing configured account types and relevant settings

### HostName

- **Type:** `string`
- **Required:** Yes

The server’s address.

### Password

- **Type:** `string`
- **Required:** No

The user’s password.

### UserName

- **Type:** `string`
- **Required:** Yes

The user’s user name.

## Discussion

Specify `com.apple.osxserver.account` as the payload type.

### Profile availability

## Topics

### Objects

- [MacOSServerAccount.ConfiguredAccountsItem](/documentation/devicemanagement/macosserveraccount/configuredaccountsitem) - An array of dictionaries containing configured account types and relevant settings

