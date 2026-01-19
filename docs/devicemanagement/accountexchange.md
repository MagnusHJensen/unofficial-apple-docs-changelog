# AccountExchange

The declaration to configure an Exchange account.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, visionOS 1.1

## Discussion

Specify `com.apple.configuration.account.exchange` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.account.exchange",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "VisibleName": "Work Exchange",
        "HostName": "exchange.example.com",
        "EnabledProtocolTypes": [
            "EAS",
            "EWS"
        ],
        "UserIdentityAssetReference": "CB3E6C7F-2318-437B-8A9E-D50C69376DE4",
        "AuthenticationCredentialsAssetReference": "64BF8F5C-8CFD-40AA-9082-A0B594D4E100",
        "LockMailService": true,
        "NotesServiceActive": false,
        "LockNotesService": true
    }
}
```

## Topics

### Objects

- [AccountExchangeOAuthObject](/documentation/devicemanagement/accountexchangeoauthobject) - The declaration for configuring OAuth authentication of an Exchange account.
- [AccountExchangeSMIMEObject](/documentation/devicemanagement/accountexchangesmimeobject) - Settings for S/MIME.

