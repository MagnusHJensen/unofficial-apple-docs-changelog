# PasscodeSettings

The declaration to configure passcode policy settings.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, visionOS 2.0, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### ChangeAtNextAuth

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system forces a password reset the next time the user tries to authenticate. If you set this key in a configuration in the system scope (device channel), the setting takes effect for all users, and admin authentication may fail until the admin user password is also reset.

### CustomRegex

- **Type:** `PasscodeSettingsCustomRegexObject`
- **Required:** No

Specifies a regular expression, and its description, to enforce password compliance. Use the simpler passcode settings whenever possible, and rely on regular expression matching only when necessary. Mistakes in regular expressions can lead to frustrating user experiences, such as unsatisfiable passcode policies, or policy descriptions that don’t match the enforced policy.

### FailedAttemptsResetInMinutes

- **Type:** `integer`
- **Required:** No

The number of minutes before the login is reset after the maximum number of failed attempts. Also set the `MaximumFailedAttempts` key for this to take effect.

### MaximumFailedAttempts

- **Type:** `integer`
- **Required:** No
- **Default:** `11`

The number of failed passcode attempts that the system allows the user before it erases or locks the device. After six failed attempts, the device imposes a time delay before the user can enter a passcode again. The time delay increases with each failed attempt. On macOS, set `FailedAttemptsResetInMinutes` to define the time delay. The time delay begins after the sixth attempt, so if `MaximumFailedAttempts` is six or lower, the system has no time delay and triggers the erase or lock as soon as the user exceeds the limit.

After the final failed attempt, the system locks a macOS device, or securely erases all data and settings from an iOS, visionOS, or watchOS device.

### MaximumGracePeriodInMinutes

- **Type:** `integer`
- **Required:** No

The maximum period that a user can select, during which the user can unlock the device without a passcode. A value of `0` means no grace period, and the device requires a passcode immediately. In the absence of this key, the user can select any period. In macOS, the system translates this to screensaver settings.

### MaximumInactivityInMinutes

- **Type:** `integer`
- **Required:** No

The maximum period that a user can select, during which the device can be idle before the system automatically locks it. When the device reaches this limit, the device locks and the user must enter the passcode to unlock it. In the absence of this key, the user can select any period. In macOS, the system translates this to screensaver settings.

### MaximumPasscodeAgeInDays

- **Type:** `integer`
- **Required:** No

Specifies the maximum number of days that the passcode can remain unchanged. After this number of days, the system forces the user to change the passcode before it unlocks the device.

### MinimumComplexCharacters

- **Type:** `integer`
- **Required:** No
- **Default:** `0`

Specifies the minimum number of complex characters in the password. A complex character is a character other than a number or a letter, such as `&`, `%`, `$`, and `#`.

### MinimumLength

- **Type:** `integer`
- **Required:** No
- **Default:** `0`

The minimum number of characters a passcode can contain.

### PasscodeReuseLimit

- **Type:** `integer`
- **Required:** No

The number of historical passcode entries the system checks when validating a new passcode. The device refuses a new passcode if it matches a previously used passcode within the specified passcode history range. In the absence of this key, the system performs no historical check.

### RequireAlphanumericPasscode

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the passcode needs to consist of at least one alphabetic character and at least one number.

### RequireComplexPasscode

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system requires a complex passcode. A complex passcode is one that doesn’t contain repeated characters or increasing or decreasing characters (such as 123 or CBA).

### RequirePasscode

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system requires the user to set a passcode without any requirements about the length or quality of the passcode. The presence of any other keys implicitly requires a passcode, and overrides this key’s value.

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

