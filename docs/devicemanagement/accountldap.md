# AccountLDAP

The declaration to configure a Lightweight Directory Access Protocol (LDAP) account.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, visionOS 1.1, Device Assignment Services , VPP License Management 

## Properties

### AuthenticationCredentialsAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the credentials for this account. Set the corresponding asset type to `CredentialUserNameAndPassword`.

### HostName

- **Type:** `string`
- **Required:** Yes

The hostname or IP address of the LDAP server.

### Port

- **Type:** `integer`
- **Required:** No

The port number or IP address of the LDAP server.

### SearchSettings

- **Type:** `[AccountLDAPSearchSettingsItemObject]`
- **Required:** No

The array of nodes to start LDAP searches from. There must be at least one node for this account to be useful. macOS only searches one node and ignores other items in the array.

### VisibleName

- **Type:** `string`
- **Required:** No

The name that apps show to the user for this LDAP account. If not present, the system generates a suitable default.

## Discussion

Specify `com.apple.configuration.account.ldap` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.account.ldap",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "VisibleName": "Work Directory",
        "HostName": "ldap.example.com",
        "SearchSettings": [
            {
                "VisibleName": "Search Work",
                "SearchBase": "dc=example,dc=com",
                "Scope": "Subtree"
            }
        ]
    }
}
```

## Topics

### Objects

- [AccountLDAPSearchSettingsItemObject](/documentation/devicemanagement/accountldapsearchsettingsitemobject) - The settings for configuring the search behavior with an LDAP server.

