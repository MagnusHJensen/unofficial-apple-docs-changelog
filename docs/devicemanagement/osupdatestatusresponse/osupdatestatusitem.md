# OSUpdateStatusResponse.OSUpdateStatusItem

A dictionary that describes the status of a software update.

**Platforms:** iOS 9.0, iPadOS 9.0, Mac Catalyst 9.0, macOS 10.11.5, tvOS 12.0, Device Assignment Services , VPP License Management 

## Properties

### DeferralsRemaining

- **Type:** `integer`
- **Required:** No

The number of remaining user deferrals for this OS update.

Available in macOS 12.3 and later.

### DownloadPercentComplete

- **Type:** `number`
- **Required:** Yes

A floating-point number between `0.0` and `1.0` that indicates the download progress as a percentage.

### IsDownloaded

- **Type:** `boolean`
- **Required:** Yes

If `true`, the update has finished downloading.

### MaxDeferrals

- **Type:** `integer`
- **Required:** No

The number of times a user can defer this OS update.

Available in macOS 12.3 and later.

### NextScheduledInstall

- **Type:** `date`
- **Required:** No

The date of the next attempt at installing this OS update.

Available in macOS 12.3 and later.

### PastNotifications

- **Type:** `[date]`
- **Required:** No

The dates/times when the OS notified the user about installing this OS update.

Available in macOS 12.3 and later.

### ProductKey

- **Type:** `string`
- **Required:** Yes

The product key that represents the update.

### Status

- **Type:** `string`
- **Required:** Yes

The status of the update, which is one of the following values:

- `Idle`: The update is idle.
- `Downloading`: The software update is downloading and subsequently preparing.
- `Installing`: The software update is installing.

