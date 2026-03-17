# AccountMail

The declaration to configure a Mail account.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, visionOS 1.1

## Properties

### IncomingServer

- **Type:** `AccountMailIncomingServerObject`
- **Required:** Yes

The settings for the incoming mail server for this account.

### OutgoingServer

- **Type:** `AccountMailOutgoingServerObject`
- **Required:** Yes

The settings for the outgoing mail server for this account.

### SMIME

- **Type:** `AccountMailSMIMEObject`
- **Required:** No

Settings for S/MIME.

### UserIdentityAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the user identity for this account. Set the corresponding asset type to `UserIdentity`.

### VisibleName

- **Type:** `string`
- **Required:** No

The name that apps show to the user for this mail account. If not present, the system generates a suitable default.

## Discussion

Specify `com.apple.configuration.account.mail` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.account.mail",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "VisibleName": "Work Mail",
        "UserIdentityAssetReference": "CB3E6C7F-2318-437B-8A9E-D50C69376DE4",
        "IncomingServer": {
            "ServerType": "IMAP",
            "HostName": "imap.example.com",
            "AuthenticationMethod": "Password",
            "AuthenticationCredentialsAssetReference": "64BF8F5C-8CFD-40AA-9082-A0B594D4E100"
        },
        "OutgoingServer": {
            "HostName": "smtp.example.com",
            "AuthenticationMethod": "Password",
            "AuthenticationCredentialsAssetReference": "64BF8F5C-8CFD-40AA-9082-A0B594D4E100"
        }
    }
}
```

## Topics

### Objects

- [AccountMailIncomingServerObject](/documentation/devicemanagement/accountmailincomingserverobject) - The settings for configuring an incoming mail server.
- [AccountMailOutgoingServerObject](/documentation/devicemanagement/accountmailoutgoingserverobject) - The settings for configuring an outgoing mail server.
- [AccountMailSMIMEObject](/documentation/devicemanagement/accountmailsmimeobject) - Settings for S/MIME.

