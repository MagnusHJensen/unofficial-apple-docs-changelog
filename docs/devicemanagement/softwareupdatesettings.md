# SoftwareUpdateSettings

The declaration to configure software updates.

**Platforms:** iOS 18.0, iPadOS 18.0, macOS 15.0, tvOS 18.4, visionOS 26.0

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

