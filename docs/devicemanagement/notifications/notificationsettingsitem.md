# Notifications.NotificationSettingsItem

The notification settings dictionary.

**Platforms:** iOS 9.3, iPadOS 9.3, Mac Catalyst 9.3, macOS 10.15, Device Assignment Services , VPP License Management 

## Properties

### AlertType

- **Type:** `integer`
- **Required:** No
- **Default:** `1`
- **Allowed Values:** `0`, `1`, `2`

The type of alert for notifications for this app:

- `0`: None
- `1`: Temporary Banner
- `2`: Persistent Banner

Available in iOS 9.3 and later and macOS 10.15 and later.

### BadgesEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables badges for this app.

Available in iOS 9.3 and later and macOS 10.15 and later.

### BundleIdentifier

- **Type:** `string`
- **Required:** Yes

The bundle identifier of the app to which to apply these notification settings.

Available in iOS 9.3 and later and macOS 10.15 and later.

### CriticalAlertEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables critical alerts that can ignore Do Not Disturb and ringer settings for this app.

Available in iOS 12 and later and macOS 10.15 and later.

### GroupingType

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`, `2`

The type of grouping for notifications for this app:

- `0`: Automatic: Group notifications into app-specified groups.
- `1`: By app: Group notifications into one group.
- `2`: Off: Don’t group notifications.

Available in iOS 12 and later.

### NotificationsEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables notifications for this app.

Available in iOS 9.3 and later and macOS 10.15 and later.

### PreviewType

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `0`, `1`, `2`

The type previews for notifications. This key overrides the value at Settings>Notifications>Show Previews.

- `0` - Always: Previews will be shown when the device is locked and unlocked
- `1` - When Unlocked: Previews will only be shown when the device is unlocked
- `2` - Never: Previews will never be shown

Available in iOS 14 and later.

### ShowInCarPlay

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables notifications in CarPlay for this app.

Available in iOS 12 and later.

### ShowInLockScreen

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables notifications on the Lock Screen for this app.

Available in iOS 9.3 and later and macOS 10.15 and later.

### ShowInNotificationCenter

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables notifications in the notification center for this app.

Available in iOS 9.3 and later and macOS 10.15 and later.

### SoundsEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables sounds for this app.

