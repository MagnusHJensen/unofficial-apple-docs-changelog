# LegacyInteractiveProfile

The declaration to configure an interactive legacy profile.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1

## Discussion

Specify `com.apple.configuration.legacy.interactive` as the declaration type.

This declaration specifies an MDMv1 profile to present to the user, who may choose to download and install the profile.

The profile may contain any payload type other than the following:

- `com.apple.mdm`
- `com.apple.declarations`

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.legacy.interactive",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "ProfileURL": "https://www.example.com/profiles/passcode.mobileconfig",
        "VisibleName": "Passcode Policy"
    }
}
```

