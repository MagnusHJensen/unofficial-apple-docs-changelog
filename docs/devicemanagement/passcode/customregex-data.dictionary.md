# Passcode.CustomRegex

The regex defining the passcode policy.

**Platforms:** macOS 14.0

## Properties

### passwordContentDescription

- **Type:** `Passcode.CustomRegex.PasswordContentDescription`
- **Required:** No

Contains a dictionary of keys for supported OS language IDs (for example, “en-US”), and whose values represent a localized description of the policy enforced by the regular expression. Use the special `default` key can for languages that aren’t contained in the dictionary.

Deprecated: iOS 27+ | iPadOS 27+ | macOS 27+ | visionOS 27+ | watchOS 27+

### passwordContentRegex

- **Type:** `string`
- **Required:** Yes

A regular expression string that the system matches against the password to determine whether it complies with a policy. The regular expression uses the ICU syntax ([https://unicode-org.github.io/icu/userguide/strings/regexp.html](https://unicode-org.github.io/icu/userguide/strings/regexp.html)). The string must not exceed 2048 characters in length.

Deprecated: iOS 27+ | iPadOS 27+ | macOS 27+ | visionOS 27+ | watchOS 27+

## Topics

### Objects

- [Passcode.CustomRegex.PasswordContentDescription](/documentation/devicemanagement/passcode/customregex-data.dictionary/passwordcontentdescription-data.dictionary) - Descriptions of the policy, localized to supported locales.

