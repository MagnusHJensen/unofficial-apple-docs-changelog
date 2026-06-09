# AppLock.App

The only app available for use on the iOS device.

**Platforms:** iOS 6.0, iPadOS 6.0, Mac Catalyst 6.0, tvOS 10.2

## Properties

### Identifier

- **Type:** `string`
- **Required:** Yes

The app’s bundle identifier.

### Options

- **Type:** `AppLock.App.Options`
- **Required:** No

A dictionary of options that the user can’t change.

Available: iOS 7+ | iPadOS 7+ | tvOS 10.2+

### UserEnabledOptions

- **Type:** `AppLock.App.UserEnabledOptions`
- **Required:** No

A dictionary of user-editable options.

Available: iOS 7+ | iPadOS 7+ | tvOS 10.2+

## Topics

### Objects

- [AppLock.App.Options](/documentation/devicemanagement/applock/app-data.dictionary/options-data.dictionary) - The dictionary of options to set for the app.
- [AppLock.App.UserEnabledOptions](/documentation/devicemanagement/applock/app-data.dictionary/userenabledoptions-data.dictionary) - The dictionary of user-editable options to set for the app.

