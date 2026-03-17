# ScheduleOSUpdateCommand.Command.UpdatesItem

A dictionary that describes the available operating-system updates item.

**Platforms:** iOS 9.0, iPadOS 9.0, macOS 10.11, tvOS 12.0

## Properties

### InstallAction

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Default`, `DownloadOnly`, `InstallASAP`, `NotifyOnly`, `InstallLater`, `InstallForceRestart`

The install action, which is one of the following values:

> 

### MaxUserDeferrals

- **Type:** `integer`
- **Required:** No

The maximum number of times the system allows the user to postpone an update before it’s installed. The system prompts the user once a day.

This key is only supported when `InstallAction` is `InstallLater` and only supported for minor OS updates (for example, macOS 12.x to 12.y).

### Priority

- **Type:** `string`
- **Required:** No
- **Default:** `Low`
- **Allowed Values:** `Low`, `High`

The scheduling priority for downloading and preparing the requested update. This is only supported for minor OS updates (macOS 12.x to 12.y).

Available in macOS 12.3 and later. Prior versions of macOS used a priority of `Low`.

### ProductKey

- **Type:** `string`
- **Required:** No

The product key that represents the update.

### ProductVersion

- **Type:** `string`
- **Required:** No

The version of the update, which the system requires if `ProductKey` isn’t present. This value is available in iOS 11.3 and later, macOS 12 and later, and tvOS 12.2 and later.

> 

