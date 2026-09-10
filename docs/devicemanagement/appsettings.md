# AppSettings

The declaration to configure app settings.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, tvOS 27.0, visionOS 27.0

## Properties

### Allowed

- **Type:** `AppSettingsAllowedObject`
- **Required:** No

The dictionary of allowed app settings.

### Privacy

- **Type:** `AppSettingsPrivacyObject`
- **Required:** No

The dictionary of app settings.

Available: iOS 27+ | iPadOS 27+ | macOS 27+
Allowed scopes: iOS: system | macOS: user

## Discussion

Specify `com.apple.configuration.app.settings` as the declaration type.

### Binary identifier rules

The following combinations of binary identifiers are supported for each key:

- `AllowedBinaries`:
- `DeniedBinaries`:

### Privacy permission defaults

Privacy permission defaults allow an organization to suggest a set of privacy permissions for use with an app. When set, the app displays a consent prompt listing all the configured defaults. If the user accepts, the device applies those defaults for the app. If the user declines, no defaults are set and the device prompts the user in the normal way when the app requires permission.

The consent prompt only shows permissions that the user hasn’t previously seen, and won’t appear if the user has seen all permissions. The user can choose from one of two options in the prompt:

- `Allow`: this option sets the app privacy permissions for the specified sub-systems (camera, microphone, and so on) to “Allow”. The device doesn’t prompt the user when the app uses the sub-system.
- `Not Now`: this option ignores the app privacy permission defaults for the specified sub-systems (camera, microphone, and so on). The device prompts the user in the normal way when the app uses the sub-system.

The user can change the app permission privacy settings in Settings.app if they choose.

Only AppKit-based apps on macOS support this feature.

### Configuration availability

### Configuration examples

#### App privacy examples

#### App settings allowlist examples

#### App settings denylist examples

#### App settings allowlist and denylist examples

This configuration allows one set of apps to run, and prevents ome other apps that would otherwise be allowed from running.

```json
{
    "Type": "com.apple.configuration.app.settings",
    "Identifier": "AF389B6F-5784-4DB6-BEFF-EA6D689BD4B3",
    "ServerToken": "A5CA3371-559E-44B4-B9ED-A0A7DFEC193D",
    "Payload": {
        "Allowed": {
            "AllowedBinaries": [
                {
                    "TeamID": "XXXXXXXXXX",
                    "CDHash": "90bc96cd95be55c12e7d9b1611cbc677610bb70c",
                    "SigningID": "com.example.app",
                    "PathPrefix": "/Applications/Example.app",
                    "SigningState": "All"
                },
                {
                    "TeamID": "*APPLE*",
                    "SigningID": "com.apple.iCal"
                }
            ],
            "DeniedBinaries": [
                {
                    "SigningID": "com.apple.iCal"
                },
                {
                    "CDHash": "03552d8140254d0c190af06f1e470dbc5ded53ba"
                }
            ]
        }
    }
}
```

#### App settings managed apps examples

This configuration always allows all managed apps and one other app to run on macOS.

```json
{
    "Type": "com.apple.configuration.app.settings",
    "Identifier": "AF389B6F-5784-4DB6-BEFF-EA6D689BD4B4",
    "ServerToken": "A5CA3371-559E-44B4-B9ED-A0A7DFEC193E",
    "Payload": {
        "Allowed": {
            "AlwaysAllowManagedApps": true,
            "AllowedBinaries": [
                {
                    "TeamID": "XXXXXXXXXX",
                    "SigningID": "com.example.app"
                }
            ]
        }
    }
}
```

#### App settings WebClips examples

## Topics

### Objects

- [AppSettingsAllowedObject](/documentation/devicemanagement/appsettingsallowedobject) - The dictionary of allowed app settings.
- [AppSettingsPrivacyObject](/documentation/devicemanagement/appsettingsprivacyobject) - The dictionary of app settings.

