# MacOSServerAccount.ConfiguredAccountsItem

An array of dictionaries containing configured account types and relevant settings

**Platforms:** iOS 9.0, iPadOS 9.0

## Properties

### Port

- **Type:** `integer`
- **Required:** No

Designates the port number to use when contacting the server. If no port number is specified, the default port is used.

### Type

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `com.apple.osxserver.documents`

com.apple.osxserver.documents (the Documents account type).

