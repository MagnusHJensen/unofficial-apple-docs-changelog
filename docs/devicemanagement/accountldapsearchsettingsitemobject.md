# AccountLDAPSearchSettingsItemObject

The settings for configuring the search behavior with an LDAP server.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, visionOS 1.1

## Properties

### Scope

- **Type:** `string`
- **Required:** No
- **Default:** `Subtree`
- **Allowed Values:** `Base`, `OneLevel`, `Subtree`

The type of recursion to use in the search:

- `Base`: The search uses only the `SearchBase` node.
- `OneLevel`: The search uses the `SearchBase` node and its immediate children.
- `Subtree`: The search uses the `SearchBase` node and all its children, regardless of depth.

### SearchBase

- **Type:** `string`
- **Required:** Yes

The path to the node where a search starts. For example, `ou=people,o=example corp`.

### VisibleDescription

- **Type:** `string`
- **Required:** No

The description of this search setting in the Contacts and Settings apps. If not present, the apps display no name.

