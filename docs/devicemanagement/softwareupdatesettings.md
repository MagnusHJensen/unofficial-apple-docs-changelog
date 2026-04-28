# SoftwareUpdateSettings

The declaration to configure software updates.

**Platforms:** iOS 18.0, iPadOS 18.0, Mac Catalyst 18.0, macOS 15.0, tvOS 18.4, visionOS 26.0, Device Assignment Services , VPP License Management 

## Properties

### AllowStandardUserOSUpdates

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If set to `true`, a standard user can perform Major and Minor Software Updates.

If set to `false`, only administrators can perform Major and Minor Software Updates.

### AutomaticActions

- **Type:** `SoftwareUpdateSettingsAutomaticActionsObject`
- **Required:** No

This object configures various automatic Software Update functionality.

### Beta

- **Type:** `SoftwareUpdateSettingsBetaObject`
- **Required:** No

This object configures the beta program settings for a device.

### Deferrals

- **Type:** `SoftwareUpdateSettingsDeferralsObject`
- **Required:** No

This object configures the deferral of software updates. Background Security Improvements aren’t considered in `Major`, `Minor`, or `System` deferral mechanism.

### Notifications

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If set to `true`, the device shows all software update enforcement notifications.

If set to `false`, the device only shows notifications triggered one hour before the enforcement deadline, and the restart countdown notification.

### RapidSecurityResponse

- **Type:** `SoftwareUpdateSettingsRapidSecurityResponseObject`
- **Required:** No

These configurations set user access to interacting with Background Security Improvement.

### RecommendedCadence

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `All`, `Oldest`, `Newest`

This string specifies how the device shows software updates to the user. When more than one update is available update, the device behaves as follows:

- `All` - Shows all software update versions.
- `Oldest` - Shows only the oldest (lower numbered) software update version.
- `Newest` - Shows only the newest (highest numbered) software update version.

## Discussion

Specify `com.apple.configuration.softwareupdate.settings` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.softwareupdate.settings",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "Notifications": false,
        "Deferrals": {
            "MajorPeriodInDays": 30
        },
        "RecommendedCadence": "All",
        "AutomaticActions": {
            "Download": "AlwaysOn",
            "InstallOSUpdates": "AlwaysOn",
            "InstallSecurityUpdate": "AlwaysOn"
        },
        "RapidSecurityResponse": {
            "Enable": false
        },
        "AllowStandardUserOSUpdates": false,
        "Beta": {
            "ProgramEnrollment": "AlwaysOn"
        }
    }
}
```

## Topics

### Objects

- [SoftwareUpdateSettingsAutomaticActionsObject](/documentation/devicemanagement/softwareupdatesettingsautomaticactionsobject) - The object that configures various automatic Software Update functionality.
- [SoftwareUpdateSettingsBetaObject](/documentation/devicemanagement/softwareupdatesettingsbetaobject) - The object that configures overall beta program settings.
- [SoftwareUpdateSettingsDeferralsObject](/documentation/devicemanagement/softwareupdatesettingsdeferralsobject) - The object that configures update deferrals.
- [SoftwareUpdateSettingsRapidSecurityResponseObject](/documentation/devicemanagement/softwareupdatesettingsrapidsecurityresponseobject) - The object that configures Background Security Improvement settings.

