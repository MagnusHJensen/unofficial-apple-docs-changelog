# LDAP.LDAPSearchSettingsItem

An array of search settings dictionaries.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, visionOS 1.1, Device Assignment Services , VPP License Management 

## Properties

### LDAPSearchSettingDescription

- **Type:** `string`
- **Required:** No

The description of this search setting.

### LDAPSearchSettingScope

- **Type:** `string`
- **Required:** No
- **Default:** `LDAPSearchSettingScopeSubtree`
- **Allowed Values:** `LDAPSearchSettingScopeBase`, `LDAPSearchSettingScopeOneLevel`, `LDAPSearchSettingScopeSubtree`

The type of recursion to use in the search:

- `LDAPSearchSettingScopeBase`: The search uses only the immediate node that the search base points to.
- `LDAPSearchSettingScopeOneLevel`: The search uses the node plus its immediate children.
- `LDAPSearchSettingScopeSubtree`: The search uses the node plus all children, regardless of depth.

### LDAPSearchSettingSearchBase

- **Type:** `string`
- **Required:** Yes

The path to the node where a search should start.

