# SafariSettings

The declaration to configure Safari settings.

**Platforms:** iOS 26.0, iPadOS 26.0, Mac Catalyst 26.0, macOS 26.0, visionOS 26.0

## Properties

### AcceptCookies

- **Type:** `string`
- **Required:** No
- **Default:** `Always`
- **Allowed Values:** `Never`, `CurrentWebsite`, `VisitedWebsites`, `Always`

The policy Safari uses for managing cookies:

- `Never`: Safari always blocks cookies.
- `CurrentWebsite`: Safari allows cookies only from the current website.
- `VisitedWebsites`: Safari allows cookies only from visited websites.
- `Always`: Safari always allows cookies.

Available: iOS 26+ | iPadOS 26+
Allowed enrollments: supervised

### AllowDisablingFraudWarning

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system forces fraud warnings on in Safari.

Available: iOS 26+ | iPadOS 26+
Allowed enrollments: supervised

### AllowHistoryClearing

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system disables clearing history in Safari.

Allowed enrollments: supervised

### AllowJavaScript

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system disables JavaScript in Safari.

Available: iOS 26+ | iPadOS 26+
Allowed enrollments: supervised

### AllowPopups

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system disables popups in Safari.

Available: iOS 26+ | iPadOS 26+
Allowed enrollments: supervised

### AllowPrivateBrowsing

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system disables private browsing in Safari.

Allowed enrollments: supervised

### AllowSummary

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system disables summarization of content in Safari.

Allowed enrollments: supervised

### NewTabStartPage

- **Type:** `SafariSettingsNewTabStartPageObject`
- **Required:** No

Sets the start page for new tabs in Safari.

### Privacy

- **Type:** `SafariSettingsPrivacyObject`
- **Required:** No

The dictionary of website privacy settings.

Available: iOS 27+ | iPadOS 27+ | macOS 27+
Allowed enrollments: supervised

## Discussion

Specify `com.apple.configuration.safari.settings` as the declaration type.

### Privacy permission defaults

Privacy permission defaults allow an organization to suggest a set of privacy permissions for use on a website. When set, Safari displays a consent prompt listing all the configured defaults. If the user accepts, the system applies those defaults for the website. If the user declines, no defaults are set and Safari prompts the user in the normal way when the website requires permission.

The consent prompt only shows permissions that the user hasn’t previously seen, and won’t appear if the user has seen all permissions. The user can choose from one of two options in the prompt:

- `Allow`: this option sets the website privacy permissions for the specified sub-systems (camera or microphone) to “Allow”. Safari doesn’t prompt the user when the website uses the sub-system.
- `Not Now`: this option ignores the website privacy permission defaults for the specified sub-systems (camera or microphone). Safari prompts the user in the normal way when the website uses the sub-system.

The user can change the website privacy permission settings in Safari settings if they choose.

### Configuration availability

### Configuration examples

#### Settings and restrictions examples

#### Website privacy examples

## Topics

### Objects

- [SafariSettingsNewTabStartPageObject](/documentation/devicemanagement/safarisettingsnewtabstartpageobject) - Sets the start page for new tabs in Safari.
- [SafariSettingsPrivacyObject](/documentation/devicemanagement/safarisettingsprivacyobject) - The dictionary of website privacy settings.

