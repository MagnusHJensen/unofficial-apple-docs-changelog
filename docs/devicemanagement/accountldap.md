# AccountLDAP

The declaration to configure a Lightweight Directory Access Protocol (LDAP) account.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, visionOS 1.1

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

