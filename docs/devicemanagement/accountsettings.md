# AccountSettings

The declaration to configure account settings.

**Platforms:** macOS 27.2

## Properties

### Exchange

- **Type:** `AccountSettingsExchangeObject`
- **Required:** No

If present, configures settings for Exchange accounts.

## Discussion

Specify `com.apple.configuration.account.settings` as the declaration type.

### Configuration availability

### Configuration example

This configuration prevents alerts about the retirement of Exchange Web Services.

```json
{
    "Type": "com.apple.configuration.account.settings",
    "Identifier": "D699D1DB-28B3-4899-8554-1630D201EF55",
    "ServerToken": "62EA7624-7603-4C15-B747-7334476C1BDE",
    "Payload": {
        "Exchange": {
            "AllowEWSRetirementAlerts": false
        }
    }
}
```

## Topics

### Objects

- [AccountSettingsExchangeObject](/documentation/devicemanagement/accountsettingsexchangeobject) - If present, configures settings for Exchange accounts.

