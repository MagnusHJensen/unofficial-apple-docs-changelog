# AccountExchange

The declaration to configure an Exchange account.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, visionOS 1.1, Device Assignment Services , VPP License Management 

## Properties

### AuthenticationCredentialsAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the credentials for this account to authenticate with an Exchange server. Set the corresponding asset type to `CredentialUserNameAndPassword`.

### AuthenticationIdentityAssetReference

- **Type:** `string`
- **Required:** No

The identifier of a credential asset declaration that contains the identity that this account requires to authenticate with the Exchange server.

### CalendarServiceActive

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, activates the calendar service for this account.

### ContactsServiceActive

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, activates the address book service for this account.

### EnabledProtocolTypes

- **Type:** `[string]`
- **Required:** Yes
- **Allowed Values:** `EAS`, `EWS`

The set of protocol types to enable on the Exchange server, in order of preference. This is an array of unique strings with possible values:

- `EAS:` Exchange ActiveSync
- `EWS:` Exchange Web Services

If the device supports one or more of the listed protocol types, it sets up an account for the first supported type.

If the device doesn’t support any of the listed protocol types, it doesn’t set up an account and the system reports an error.

### External Path

- **Type:** `string`
- **Required:** No

The external path of the EWS server. The system uses this only when this declaration has a `ExternalHostName` value.

### ExternalHostName

- **Type:** `string`
- **Required:** No

The external hostname of the EWS server (or IP address).

### ExternalPort

- **Type:** `integer`
- **Required:** No

The external port number of the EWS server. The system uses this only when this declaration has a `ExternalHostName` value.

### HostName

- **Type:** `string`
- **Required:** No

The IP address or fully qualified domain name (FQDN) of the Exchange host.

### LockCalendarService

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system prevents the user from changing the status of the calendar service for this account.

### LockContactsService

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system prevents the user from changing the status of the address book service for this account.

### LockMailService

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system prevents the user from changing the status of the mail service for this account.

### LockNotesService

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system prevents the user from changing the status of the notes service for this account.

### LockRemindersService

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system prevents the user from changing the status of the reminders service for this account.

### MailServiceActive

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the system activates the mail service for this account.

### NotesServiceActive

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the system activates the notes service for this account.

### OAuth

- **Type:** `AccountExchangeOAuthObject`
- **Required:** No

The configuration settings for OAuth for this account.

### Path

- **Type:** `string`
- **Required:** No

The path of the EWS server. The system uses this only when this declaration has a `HostName` value.

### Port

- **Type:** `integer`
- **Required:** No

The port number of the EWS server. The system uses this only when this declaration has a `HostName` value.

### RemindersServiceActive

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the system activates the reminders service for this account.

### SMIME

- **Type:** `AccountExchangeSMIMEObject`
- **Required:** No

Settings for S/MIME.

### UserIdentityAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the user identity for this account. The corresponding asset must be of type `UserIdentity`.

### VisibleName

- **Type:** `string`
- **Required:** No

The name that apps show to the user for this Exchange account. If not present, the system generates a suitable default.

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

