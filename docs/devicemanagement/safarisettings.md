# SafariSettings

The declaration to configure Safari settings.

**Platforms:** iOS 26.0, iPadOS 26.0, Mac Catalyst 26.0, macOS 26.0, visionOS 26.0, Device Assignment Services , VPP License Management 

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

### AllowDisablingFraudWarning

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system forces fraud warnings on in Safari.

### AllowHistoryClearing

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system disables clearing history in Safari.

### AllowJavaScript

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system disables JavaScript in Safari.

### AllowPopups

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system disables popups in Safari.

### AllowPrivateBrowsing

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system disables private browsing in Safari.

### AllowSummary

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system disables summarization of content in Safari.

### NewTabStartPage

- **Type:** `SafariSettingsNewTabStartPageObject`
- **Required:** No

Sets the start page for new tabs in Safari.

## Discussion

Specify `com.apple.configuration.safari.settings` as the declaration type.

### Configuration availability

### Configuration examples

## Topics

### Objects

- [SafariSettingsNewTabStartPageObject](/documentation/devicemanagement/safarisettingsnewtabstartpageobject) - Sets the start page for new tabs in Safari.

