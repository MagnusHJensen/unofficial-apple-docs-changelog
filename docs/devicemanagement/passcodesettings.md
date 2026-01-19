# PasscodeSettings

The declaration to configure passcode policy settings.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, visionOS 2.0, watchOS 10.0

## Discussion

Specify `com.apple.configuration.passcode.settings` as the declaration type.

The presence of this configuration type causes the device to present the user with a passcode entry mechanism. The configuration controls the complexity of the passcode.

For user enrollments, the system allows this configuration type, but ignores most of the keys. Instead, the presence of the configuration forces only these settings:

- `RequirePasscode`: always set to `true`
- `RequireComplexPasscode`: always set to `true`
- `MinimumLength`: always set to `6`
- `MaximumInactivityInMinutes`: if this key is present its value is ignored, but the `never` option is removed in the Settings UI.

### Configuration availability

### Configuration examples

## Topics

### Objects

- [PasscodeSettingsCustomRegexObject](/documentation/devicemanagement/passcodesettingscustomregexobject) - A regular expression and its description to enforce password compliance.

